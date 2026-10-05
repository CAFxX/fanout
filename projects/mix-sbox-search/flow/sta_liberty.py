#!/usr/bin/env python3
"""Simplified liberty STA for combinational mapped netlists.

Purpose: give the GH S-box screen a REAL cell-delay critical path (ps,
typical corner) to complement the unit-delay ranking — without a full P&R.

Documented simplifications (all deterministic; wire effects are EXCLUDED
by design — measuring those is the P&R stage's job, see README calibration
plan):
  * output load = sum of driven standard-cell input-pin capacitances
    (no wire capacitance);
  * input slew propagated via the liberty rise/fall_transition tables;
    primary inputs use SREF_NS reference slew;
  * rise/fall tracked separately; arc timing_sense honored;
  * delay tables bilinearly interpolated with clamping at the rails;
  * combinational loops are a hard error.

Usage: sta_liberty.crit_ps(mapped_verilog, liberty_path) -> dict
  {crit_ps, n_cells, area_um2}
"""
import functools
import re

SREF_NS = 0.06  # reference slew for primary inputs, ns


# ---------------- liberty parser ----------------
def _tokenize(text):
    text = re.sub(r"/\*.*?\*/", " ", text, flags=re.S)
    text = re.sub(r"//[^\n]*", " ", text)
    text = text.replace("\\\n", " ")
    toks = []
    i, n = 0, len(text)
    while i < n:
        c = text[i]
        if c in " \t\r\n":
            i += 1
        elif c in "{}();:,":
            toks.append(c)
            i += 1
        elif c == '"':
            j = text.index('"', i + 1)
            toks.append(("STR", text[i + 1:j]))
            i = j + 1
        else:
            j = i
            while j < n and text[j] not in " \t\r\n{}();:,":
                j += 1
            toks.append(text[i:j])
            i = j
    return toks


def _parse_group(toks, pos):
    """group_type [ (args...) ] [{ ... } | ;] -> (node, pos).
    node: {type, args (list, unjoined), attrs (name -> [parts]),
           children}. Bare statements (define;) have no children."""
    gtype = toks[pos]
    pos += 1
    args = []
    if pos < len(toks) and toks[pos] == "(":
        depth = 0
        while True:
            t = toks[pos]
            if t == "(":
                depth += 1
            elif t == ")":
                depth -= 1
                if depth == 0:
                    pos += 1
                    break
            else:
                if not (isinstance(t, str) and t == ","):
                    args.append(t[1] if isinstance(t, tuple) else t)
            pos += 1
    if pos < len(toks) and toks[pos] == ";":
        return {"type": gtype, "args": args, "attrs": {},
                "children": []}, pos + 1
    assert toks[pos] == "{", f"expected {{ in {gtype}"
    pos += 1
    attrs, children = {}, []
    while toks[pos] != "}":
        t = toks[pos]
        if (isinstance(t, str) and re.fullmatch(r"[A-Za-z_][\w.]*", t)
                and pos + 1 < len(toks) and toks[pos + 1] == ":"):
            aname = t
            pos += 2
            parts = []
            while toks[pos] != ";":
                vt = toks[pos]
                parts.append(vt[1] if isinstance(vt, tuple) else vt)
                pos += 1
                if toks[pos] == ",":
                    pos += 1
            pos += 1
            attrs[aname] = parts
        else:
            child, pos = _parse_group(toks, pos)
            children.append(child)
    return {"type": gtype, "args": args, "attrs": attrs,
            "children": children}, pos + 1


def _numlist(s):
    return [float(x) for x in s.replace(",", " ").split()]


def _scalar(attrs, key, default=0.0):
    v = attrs.get(key)
    if not v:
        return default
    try:
        return float(v[0].replace(",", " ").split()[0])
    except (ValueError, IndexError):
        return default


def _table(tc):
    """timing-table group node -> (index_1, index_2, rows) or None."""
    parts = {}
    for ch in tc["children"]:
        if ch["type"] in ("index_1", "index_2", "values"):
            parts[ch["type"]] = ch["args"]
    if not all(k in parts for k in ("index_1", "index_2", "values")):
        return None
    return (_numlist(parts["index_1"][0]), _numlist(parts["index_2"][0]),
            [_numlist(r) for r in parts["values"]])


@functools.lru_cache(maxsize=2)
def load_liberty(path):
    """path -> {cell: {area, pins: {p: {dir, cap}},
    arcs: [(related_pin, out_pin, sense, tables)]}},
    tables: {cell_rise, cell_fall, rise_transition, fall_transition}
            -> (i1, i2, rows) | None."""
    toks = _tokenize(open(path).read())
    lib, _ = _parse_group(toks, 0)
    cells = {}
    for ch in lib["children"]:
        if ch["type"] != "cell" or not ch["args"]:
            continue
        cname = ch["args"][0]
        pins, arcs = {}, []
        area = _scalar(ch["attrs"], "area")
        for p in ch["children"]:
            if p["type"] != "pin" or not p["args"]:
                continue
            pname = p["args"][0]
            pins[pname] = {
                "dir": (p["attrs"].get("direction") or ["input"])[0],
                "cap": _scalar(p["attrs"], "capacitance"),
            }
            for t in p["children"]:
                if t["type"] != "timing":
                    continue
                rel = (t["attrs"].get("related_pin") or [None])[0]
                sense = (t["attrs"].get("timing_sense") or
                         ["non_unate"])[0]
                tables = {}
                for tc in t["children"]:
                    if tc["type"] in ("cell_rise", "cell_fall",
                                      "rise_transition", "fall_transition"):
                        tab = _table(tc)
                        if tab:
                            tables[tc["type"]] = tab
                arcs.append((rel, pname, sense, tables))
        cells[cname] = {"area": area, "pins": pins, "arcs": arcs}
    return cells


# ---------------- bilinear interpolation ----------------
def _bilerp(i1, i2, rows, x, y):
    def brack(v, xs):
        if v <= xs[0]:
            return 0, 0, 0.0
        if v >= xs[-1]:
            return len(xs) - 1, len(xs) - 1, 0.0
        for k in range(len(xs) - 1):
            if xs[k] <= v <= xs[k + 1]:
                f = 0.0 if xs[k + 1] == xs[k] else (v - xs[k]) / (xs[k + 1] - xs[k])
                return k, k + 1, f
        return len(xs) - 1, len(xs) - 1, 0.0
    x0, x1, fx = brack(x, i1)
    y0, y1, fy = brack(y, i2)
    v00, v01 = rows[x0][y0], rows[x0][y1]
    v10, v11 = rows[x1][y0], rows[x1][y1]
    return (v00 * (1 - fx) + v10 * fx) * (1 - fy) + \
           (v01 * (1 - fx) + v11 * fx) * fy


def _lookup(tables, key, slew, load):
    t = tables.get(key)
    if t is None:
        return 0.0
    i1, i2, rows = t
    if not rows or not rows[0]:
        return 0.0
    return _bilerp(i1, i2, rows, slew, load)


# ---------------- netlist ----------------
def _parse_netlist(text):
    def expand(decl):
        m = re.match(r"\[(.*?):(.*?)\]\s*(\w+)", decl)
        if m:
            hi, lo, nm = int(m.group(1)), int(m.group(2)), m.group(3)
            return [f"{nm}[{i}]" for i in range(lo, hi + 1)]
        return [decl.strip()]

    def norm(n):
        return n.replace("\\", "").replace(" ", "")

    top_inputs, top_outputs = [], []
    for m in re.finditer(r"^\s*input\s+([^;]+);", text, re.M):
        for d in m.group(1).split(","):
            top_inputs.extend(expand(d))
    for m in re.finditer(r"^\s*output\s+([^;]+);", text, re.M):
        for d in m.group(1).split(","):
            top_outputs.extend(expand(d))
    cells = []
    for m in re.finditer(r"^\s*(\w+)\s+_\d+_\s*\((.*?)\);", text, re.M | re.S):
        gtype, conn = m.group(1), m.group(2)
        pins = {p: norm(n) for p, n in
                re.findall(r"\.(\w+)\(\s*([^),]+?)\s*\)", conn)}
        cells.append((gtype, pins))
    # trivial aliases (constants): resolve to literal
    alias = {}
    for m in re.finditer(r"^\s*assign\s+([^;]+);", text, re.M):
        stmt = m.group(1).strip()
        if "=" in stmt:
            lhs, rhs = (norm(x) for x in stmt.split("=", 1))
            if re.fullmatch(r"1'b[01]", rhs):
                alias[lhs] = rhs

    def resolve(net):
        seen = set()
        while net in alias and net not in seen:
            seen.add(net)
            net = alias[net]
        return net

    return top_inputs, top_outputs, cells, resolve


# ---------------- STA ----------------
def crit_ps(mapped_verilog, liberty_path):
    """-> {crit_ps, n_cells, area_um2} (crit_ps float, ns->ps)."""
    lib = load_liberty(liberty_path)
    text = open(mapped_verilog).read()
    top_inputs, top_outputs, cells, resolve = _parse_netlist(text)

    # driver map + fanout pin caps
    driver = {}   # net -> (cell_idx, out_pin)
    loads = {}    # net -> total sink pin cap (pF)
    for idx, (gtype, pins) in enumerate(cells):
        spec = lib.get(gtype)
        if spec is None:
            raise RuntimeError(f"cell {gtype} not in liberty file")
        for p, net in pins.items():
            net = resolve(net)
            pdir = spec["pins"].get(p, {}).get("dir", "input")
            if pdir == "output":
                driver[net] = (idx, p)
            else:
                loads[net] = loads.get(net, 0.0) + \
                    spec["pins"].get(p, {}).get("cap", 0.0)

    # topological order (Kahn) over cells
    indeg = [0] * len(cells)
    adj = [[] for _ in range(len(cells))]
    net_src = {}
    for net, (idx, _) in driver.items():
        net_src[net] = idx
    cell_inputs = []
    for idx, (gtype, pins) in enumerate(cells):
        spec = lib[gtype]
        ins = []
        for p, net in pins.items():
            net = resolve(net)
            if spec["pins"].get(p, {}).get("dir", "input") != "output":
                if net in net_src and net_src[net] != idx:
                    ins.append(net_src[net])
        cell_inputs.append(ins)
        for s in set(ins):
            adj[s].append(idx)
        indeg[idx] = len(set(ins))
    order, queue = [], [i for i in range(len(cells)) if indeg[i] == 0]
    while queue:
        u = queue.pop()
        order.append(u)
        for w in adj[u]:
            indeg[w] -= 1
            if indeg[w] == 0:
                queue.append(w)
    if len(order) != len(cells):
        raise RuntimeError("combinational loop in mapped netlist")

    arr_r = {n: 0.0 for n in top_inputs}
    arr_f = {n: 0.0 for n in top_inputs}
    slew_r = {n: SREF_NS for n in top_inputs}
    slew_f = {n: SREF_NS for n in top_inputs}

    def get(d, net, default=0.0):
        return d.get(resolve(net), default)

    for idx in order:
        gtype, pins = cells[idx]
        spec = lib[gtype]
        rpins = {p: resolve(n) for p, n in pins.items()}
        for rel, opin, sense, tables in spec["arcs"]:
            if rel not in rpins or opin not in rpins:
                continue
            onet = rpins[opin]
            load = loads.get(onet, 0.0)
            pairs = []
            if sense == "positive_unate":
                pairs = [("r", "r", "cell_rise", "rise_transition"),
                         ("f", "f", "cell_fall", "fall_transition")]
            elif sense == "negative_unate":
                pairs = [("r", "f", "cell_fall", "fall_transition"),
                         ("f", "r", "cell_rise", "rise_transition")]
            else:
                pairs = [("r", "r", "cell_rise", "rise_transition"),
                         ("r", "f", "cell_fall", "fall_transition"),
                         ("f", "r", "cell_rise", "rise_transition"),
                         ("f", "f", "cell_fall", "fall_transition")]
            for si, so, dk, tk in pairs:
                a = get(arr_r if si == "r" else arr_f, rpins[rel])
                s = get(slew_r if si == "r" else slew_f, rpins[rel],
                        SREF_NS)
                dly = _lookup(tables, dk, s, load)
                dsl = _lookup(tables, tk, s, load)
                A = arr_r if so == "r" else arr_f
                S = slew_r if so == "r" else slew_f
                if a + dly > A.get(onet, 0.0):
                    A[onet] = a + dly
                    S[onet] = dsl

    crit_ns = 0.0
    for o in top_outputs:
        o = resolve(o)
        crit_ns = max(crit_ns, arr_r.get(o, 0.0), arr_f.get(o, 0.0))
    area = sum(lib[g]["area"] for g, _ in cells)
    return {"crit_ps": crit_ns * 1000.0, "n_cells": len(cells),
            "area_um2": area}


if __name__ == "__main__":
    import json
    import sys
    r = crit_ps(sys.argv[1], sys.argv[2])
    print(json.dumps(r, indent=1))
