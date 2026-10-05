#!/usr/bin/env python3
"""Agent-side helper for the fanout repo — push/commit loop (no gh CLI).

The authenticated `github` CLI (Muse GitHub App) exposes NO Actions APIs
(no workflow dispatch, no run polling, no artifact download), and `gh` is
not logged in. So the loop is push-triggered with results committed back:

  submit:    write jobs/<project>/<job-id>.json via
             `github call-tool --name push_files`
             -> dispatch.yml triggers on push, runs the shard matrix, and
                its collect job commits results/<project>/<job-id>/
                (shard-<i>/ + STATUS.json + FILES.json) via GITHUB_TOKEN
  poll:      read results/<project>/<job-id>/STATUS.json via
             `github call-read-tool --name get_file_contents` until terminal
             (done|partial|failed)
  fetch:     download result files via get_file_contents into dest/
  aggregate: fetch + run the project's aggregate script locally
  budget:    ledger estimates vs the hard cap (metered actuals live on
             github.com billing — no billing API on this tool surface)
  queue:     pending submits ordered by priority

Budget: 2000 Linux-min/month shared across ALL projects, hard cap 1800
(default, configurable). Every submit prints the estimate first and
REFUSES without --confirm, and REFUSES when over the hard cap.
"""
import argparse
import base64
import datetime
import json
import os
import subprocess
import sys
import time

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import budget as budget_mod

DEFAULT_REPO = "CAFxX/fanout"
DEFAULT_BRANCH = "main"
TERMINAL = ("done", "partial", "failed")


def sh(*args):
    r = subprocess.run(args, capture_output=True, text=True)
    if r.returncode != 0:
        sys.exit(f"command failed: {' '.join(args)}\n{r.stderr[-2000:]}")
    return r.stdout.strip()


def gh_read(owner, repo, path):
    """get_file_contents; returns (found, parsed-or-text)."""
    out = sh("github", "call-read-tool", "--name", "get_file_contents",
             "--arguments-json", json.dumps(
                 {"owner": owner, "repo": repo, "path": path}))
    try:
        d = json.loads(out)
    except json.JSONDecodeError:
        return False, out[-500:]
    content = d.get("result", {}).get("content", [])
    if not content:
        return False, "empty result"
    text = content[0].get("text", "")
    if content[0].get("isError"):
        return False, text[:500]
    try:
        payload = json.loads(text)
    except json.JSONDecodeError:
        return False, text[:500]
    # file content arrives base64-encoded
    if isinstance(payload, dict) and payload.get("encoding") == "base64" \
            and "content" in payload:
        raw = base64.b64decode(payload["content"]).decode()
        try:
            return True, json.loads(raw)
        except json.JSONDecodeError:
            return True, raw
    return True, payload


def gh_push(owner, repo, branch, path, content_str, message):
    out = sh("github", "call-tool", "--name", "push_files",
             "--arguments-json", json.dumps({
                 "owner": owner, "repo": repo, "branch": branch,
                 "message": message,
                 "files": [{"path": path, "content": content_str}]}))
    return out


def split_repo(repo):
    owner, name = repo.split("/", 1)
    return owner, name


def sanitize(s):
    return "".join(c if (c.isalnum() or c in "-_") else "-" for c in s)


def build_job_id(project, purpose, job_id):
    ts = datetime.datetime.now(datetime.timezone.utc).strftime("%Y%m%d-%H%M%S")
    job_id = job_id or f"{ts}-{sanitize(purpose)}"
    return sanitize(job_id)


def submit_spec(owner, repo, branch, spec, priority, est_min, cap,
                confirm):
    """Budget gate + push a job spec. Returns (project, job_id)."""
    try:
        b = budget_mod.check_submit(est_min, priority, cap)
    except budget_mod.BudgetRefused as e:
        sys.exit(str(e))
    print(f"ESTIMATE: {spec['shards']} shards x "
          f"{spec.get('per_shard_min', '?')} min/shard = {est_min} minutes "
          f"[priority: {priority}]")
    print(f"BUDGET: {b['spent']} min projected this month; hard cap "
          f"{b['cap']}; {b['remaining_after']} min would remain")
    print("NOTE: estimates only — metered actuals are on github.com billing")
    if not confirm:
        sys.exit("REFUSED: re-run with --confirm to submit "
                 "(estimates are never submitted silently)")
    path = f"jobs/{spec['project']}/{spec['job_id']}.json"
    gh_push(owner, repo, branch, path,
            json.dumps(spec, indent=1) + "\n",
            f"fanout submit: {spec['project']}/{spec['job_id']} "
            f"({priority})")
    budget_mod.record_submit(spec["job_id"], spec["project"], est_min,
                             priority, spec["shards"],
                             spec.get("per_shard_min", 0))
    print(f"SUBMITTED {spec['project']}/{spec['job_id']}")
    print(f"poll with: fanout.py poll --project {spec['project']} "
          f"--job-id {spec['job_id']} --repo {owner}/{repo}")
    print(f"this submit = 1 approval covering 1 job / "
          f"~{est_min:.0f} estimated minutes")
    return spec["project"], spec["job_id"]


def cmd_submit_batch(a):
    """Batch MULTIPLE job specs into ONE push_files call (one approval).

    Each --spec-file is a JSON job spec (project, job_id, image, command,
    shards, timeout_minutes, env, cache_key, cache_paths, priority,
    est_min, per_shard_min). The dispatch workflow fans out to every spec
    in the push.
    """
    owner, repo = split_repo(a.repo)
    specs = []
    total_est = 0
    for fp in a.spec_files:
        spec = json.load(open(fp))
        for f in ("project", "job_id", "image", "command"):
            if not spec.get(f):
                sys.exit(f"spec {fp}: missing {f}")
        spec["project"] = sanitize(spec["project"])
        spec["job_id"] = sanitize(spec["job_id"])
        spec["shards"] = max(1, min(64, int(spec.get("shards", 8))))
        spec["timeout_minutes"] = max(1, min(300,
                                     int(spec.get("timeout_minutes", 60))))
        est = float(spec.get("est_min",
                             spec["shards"] *
                             float(spec.get("per_shard_min", 0))))
        spec["est_min"] = est
        spec.setdefault("priority", a.priority)
        spec["submitted_at"] = datetime.datetime.now(
            datetime.timezone.utc).isoformat()
        # budget gate per spec (hard cap refuses)
        try:
            budget_mod.check_submit(est, spec["priority"], a.cap)
        except budget_mod.BudgetRefused as e:
            sys.exit(f"spec {spec['job_id']}: {e}")
        specs.append(spec)
        total_est += est
    if not a.confirm:
        sys.exit("REFUSED: re-run with --confirm to submit "
                 "(estimates are never submitted silently)")
    files = [{"path": f"jobs/{s['project']}/{s['job_id']}.json",
              "content": json.dumps(s, indent=1) + "\n"} for s in specs]
    out = sh("github", "call-tool", "--name", "push_files",
             "--arguments-json", json.dumps({
                 "owner": owner, "repo": repo, "branch": a.branch,
                 "message": f"fanout batch submit: {len(specs)} jobs",
                 "files": files}))
    for s in specs:
        budget_mod.record_submit(s["job_id"], s["project"],
                                 s["est_min"], s["priority"],
                                 s["shards"],
                                 s.get("per_shard_min", 0))
        print(f"SUBMITTED {s['project']}/{s['job_id']}")
    print(f"this submit = 1 approval covering {len(specs)} jobs / "
          f"~{total_est:.0f} estimated minutes")
    print("NOTE: estimates only — metered actuals are on github.com billing")


def poll_until_terminal(owner, repo, project, job_id, timeout_min,
                        quiet=False):
    """Poll STATUS.json until done|partial|failed. Returns the payload."""
    path = status_path(project, job_id)
    deadline = time.time() + timeout_min * 60
    last = None
    while True:
        found, data = gh_read(owner, repo, path)
        if found and isinstance(data, dict):
            st = data.get("status")
            if st != last and not quiet:
                print(f"STATUS: {st} "
                      f"({data.get('n_ok', '?')}/{data.get('shards', '?')} "
                      f"shards ok)", flush=True)
                last = st
            if st in TERMINAL:
                budget_mod.mark_status(job_id, st)
                return data
        elif not quiet:
            print(f"waiting for {path} ...", flush=True)
        if time.time() > deadline:
            sys.exit("poll timeout: STATUS.json never reached a terminal "
                     "state — the push may not have triggered dispatch; "
                     "check the repo Actions tab")
        time.sleep(30)


def fetch_results(owner, repo, project, job_id, dest):
    """Download every file in FILES.json into dest/."""
    base = f"results/{project}/{job_id}"
    found, manifest = gh_read(owner, repo, base + "/FILES.json")
    if not found:
        sys.exit(f"no FILES.json at {base}/FILES.json — "
                 f"has the collect job finished?")
    dest = os.path.abspath(dest)
    for rel in manifest.get("files", []):
        found, data = gh_read(owner, repo, base + "/" + rel)
        if not found:
            print(f"WARNING: missing {rel}")
            continue
        p = os.path.join(dest, rel)
        os.makedirs(os.path.dirname(p), exist_ok=True)
        if isinstance(data, (dict, list)):
            open(p, "w").write(json.dumps(data, indent=1))
        else:
            open(p, "w").write(data)
    print(f"fetched {len(manifest.get('files', []))} files -> {dest}")
    return dest


def cmd_submit(a):
    owner, repo = split_repo(a.repo)
    project = sanitize(a.project)
    purpose = sanitize(a.purpose or "job")
    job_id = build_job_id(project, purpose, a.job_id)
    shards = int(a.shards)
    est = shards * a.per_shard_min

    spec = {
        "project": project,
        "job_id": job_id,
        "image": a.image,
        "command": a.command,
        "shards": shards,
        "timeout_minutes": int(a.timeout_minutes),
        "env": a.env or "",
        "cache_key": a.cache_key or "",
        "cache_paths": a.cache_paths or "",
        "registry_user": a.registry_user or "",
        "registry_password": a.registry_password or "",
        "priority": a.priority,
        "est_min": est,
        "per_shard_min": a.per_shard_min,
        "submitted_at": datetime.datetime.now(
            datetime.timezone.utc).isoformat(),
    }
    submit_spec(owner, repo, a.branch, spec, a.priority, est, a.cap,
                a.confirm)


def status_path(project, job_id):
    return f"results/{project}/{job_id}/STATUS.json"


def cmd_poll(a):
    owner, repo = split_repo(a.repo)
    data = poll_until_terminal(owner, repo, a.project, a.job_id, a.timeout)
    print(f"terminal: {data.get('status')}")
    if data.get("status") != "done":
        for s in data.get("shard_status", []):
            if s.get("driver") != "done":
                print(f"  shard {s['idx']}: {s.get('detail')}")


def cmd_fetch(a):
    owner, repo = split_repo(a.repo)
    fetch_results(owner, repo, a.project, a.job_id, a.dest)


def cmd_aggregate(a):
    cmd_fetch(a)
    dest = os.path.abspath(a.dest)
    agg = subprocess.run(
        [sys.executable, a.script, "--shards-dir", dest] + a.script_args,
        capture_output=True, text=True)
    print(agg.stdout)
    if agg.returncode != 0:
        sys.exit(f"aggregate script failed:\n{agg.stderr[-3000:]}")
    print("aggregate OK")


def cmd_budget(a):
    month = budget_mod.month_key()
    entries = [e for e in budget_mod.load_ledger() if e["month"] == month]
    spent = sum(e["est_min"] for e in entries)
    cap = budget_mod.cap() if a.cap is None else a.cap
    print(f"month {month}: {spent}/{cap} min projected "
          f"({cap - spent} remaining of hard cap; plan gives 2000)")
    by_proj = budget_mod.spent_by_project(entries, month)
    for proj, mins in sorted(by_proj.items(), key=lambda t: -t[1]):
        print(f"  {proj}: {mins} min")
    for e in entries:
        print(f"  {e['ts'][:16]} {e['project']}/{e['job_id']} "
              f"[{e.get('priority', '?')}] {e['shards']}x{e['per_shard_min']} "
              f"= {e['est_min']} min ({e.get('last_status', 'submitted')})")
    print("metered actuals: github.com repo Settings > Billing")


def cmd_queue(_a):
    pend = budget_mod.pending_by_priority()
    if not pend:
        print("queue empty")
        return
    for e in pend:
        print(f"  [{e.get('priority', '?')}] {e['project']}/{e['job_id']} "
              f"{e['est_min']} min ({e.get('last_status', 'submitted')})")


def main():
    ap = argparse.ArgumentParser(prog="fanout.py")
    sub = ap.add_subparsers(dest="cmd", required=True)

    p = sub.add_parser("submit")
    p.add_argument("--project", required=True)
    p.add_argument("--purpose", default="job")
    p.add_argument("--job-id", default="")
    p.add_argument("--image", required=True)
    p.add_argument("--command", required=True)
    p.add_argument("--shards", default="8")
    p.add_argument("--timeout-minutes", default="60")
    p.add_argument("--per-shard-min", type=float, default=12.0)
    p.add_argument("--env", default="")
    p.add_argument("--cache-key", default="")
    p.add_argument("--cache-paths", default="")
    p.add_argument("--registry-user", default="")
    p.add_argument("--registry-password", default="")
    p.add_argument("--priority", default="new-leads",
                   choices=budget_mod.PRIORITIES)
    p.add_argument("--cap", type=int, default=None,
                   help="hard-cap override (default 1800 / $FANOUT_CAP)")
    p.add_argument("--repo", default=DEFAULT_REPO)
    p.add_argument("--branch", default=DEFAULT_BRANCH)
    p.add_argument("--confirm", action="store_true")
    p.set_defaults(fn=cmd_submit)

    p = sub.add_parser("submit-batch")
    p.add_argument("--spec-files", nargs="+", required=True,
                   help="job spec JSON files; all pushed in ONE push_files "
                        "call (one approval)")
    p.add_argument("--priority", default="new-leads",
                   choices=budget_mod.PRIORITIES,
                   help="default priority for specs lacking one")
    p.add_argument("--cap", type=int, default=None)
    p.add_argument("--repo", default=DEFAULT_REPO)
    p.add_argument("--branch", default=DEFAULT_BRANCH)
    p.add_argument("--confirm", action="store_true")
    p.set_defaults(fn=cmd_submit_batch)

    p = sub.add_parser("poll")
    p.add_argument("--project", required=True)
    p.add_argument("--job-id", required=True)
    p.add_argument("--repo", default=DEFAULT_REPO)
    p.add_argument("--timeout", type=int, default=180,
                   help="minutes to wait")
    p.set_defaults(fn=cmd_poll)

    p = sub.add_parser("fetch")
    p.add_argument("--project", required=True)
    p.add_argument("--job-id", required=True)
    p.add_argument("--dest", required=True)
    p.add_argument("--repo", default=DEFAULT_REPO)
    p.set_defaults(fn=cmd_fetch)

    p = sub.add_parser("aggregate")
    p.add_argument("--project", required=True)
    p.add_argument("--job-id", required=True)
    p.add_argument("--dest", required=True)
    p.add_argument("--script", required=True)
    p.add_argument("--repo", default=DEFAULT_REPO)
    p.add_argument("script_args", nargs="*")
    p.set_defaults(fn=cmd_aggregate)

    p = sub.add_parser("budget")
    p.add_argument("--cap", type=int, default=None)
    p.set_defaults(fn=cmd_budget)

    p = sub.add_parser("queue")
    p.set_defaults(fn=cmd_queue)

    a = ap.parse_args()
    a.fn(a)


if __name__ == "__main__":
    main()
