#!/usr/bin/env python3
"""Budget enforcement for the fanout repo (shared 2000 Linux-min/month).

Hard cap (default 1800 of 2000, configurable via --cap / FANOUT_CAP):
submit() REFUSES when the ledger's projected spend would exceed the cap.
Estimates are advisory (a shard that finishes early costs less); metered
actuals live on github.com billing — there is no billing API on the
authenticated tool surface, so no automatic reconciliation exists.

Priority queue (submit --priority; recorded in the ledger; the agent
schedules in this order when several batches are pending):
    frontier-challenger > new-leads > evolution > fuzz
"""
import datetime
import json
import os

LEDGER = os.path.expanduser("~/.fanout/ledger.json")
BUDGET_MIN = 2000
HARD_CAP_DEFAULT = 1800

PRIORITIES = ["frontier-challenger", "new-leads", "evolution", "fuzz"]
PRIORITY_RANK = {p: i for i, p in enumerate(PRIORITIES)}


class BudgetRefused(Exception):
    pass


def month_key(when=None):
    return (when or datetime.date.today()).strftime("%Y-%m")


def load_ledger():
    try:
        return json.load(open(LEDGER))
    except Exception:
        return []


def save_ledger(entries):
    os.makedirs(os.path.dirname(LEDGER), exist_ok=True)
    with open(LEDGER, "w") as f:
        json.dump(entries, f, indent=1)


def spent_this_month(entries=None, month=None):
    entries = load_ledger() if entries is None else entries
    month = month or month_key()
    return sum(e["est_min"] for e in entries if e["month"] == month)


def spent_by_project(entries=None, month=None):
    entries = load_ledger() if entries is None else entries
    month = month or month_key()
    out = {}
    for e in entries:
        if e["month"] == month:
            out[e.get("project", "?")] = \
                out.get(e.get("project", "?"), 0) + e["est_min"]
    return out


def cap():
    try:
        return int(os.environ.get("FANOUT_CAP", HARD_CAP_DEFAULT))
    except ValueError:
        return HARD_CAP_DEFAULT


def check_submit(est_min, priority="new-leads", cap_min=None,
                 entries=None, month=None):
    """Raise BudgetRefused if est_min would push projected spend over cap."""
    if priority not in PRIORITY_RANK:
        raise ValueError(f"unknown priority {priority!r}; "
                         f"one of {PRIORITIES}")
    cap_min = HARD_CAP_DEFAULT if cap_min is None else cap_min
    # env override wins unless an explicit cap was passed
    if cap_min == HARD_CAP_DEFAULT:
        cap_min = cap()
    spent = spent_this_month(entries, month)
    remaining = cap_min - spent
    if est_min > remaining:
        raise BudgetRefused(
            f"REFUSED: {est_min} min would exceed hard cap "
            f"({spent} spent of {cap_min}; {remaining} remaining). "
            f"Wait for next month, raise --cap, or shrink the dispatch.")
    return {"spent": spent, "remaining": remaining,
            "remaining_after": remaining - est_min, "cap": cap_min}


def record_submit(job_id, project, est_min, priority, shards,
                  per_shard_min, month=None):
    entries = load_ledger()
    month = month or month_key()
    entries.append({
        "month": month,
        "ts": datetime.datetime.now(datetime.timezone.utc).isoformat(),
        "job_id": job_id, "project": project,
        "priority": priority, "shards": shards,
        "per_shard_min": per_shard_min, "est_min": est_min,
    })
    save_ledger(entries)


def pending_by_priority(entries=None, month=None):
    """Ledger entries for jobs not known-terminal, highest priority first.

    Terminal state is only known after a poll; entries carry last-known
    status (default 'submitted'). The agent updates status via poll.
    """
    entries = load_ledger() if entries is None else entries
    month = month or month_key()
    pend = [e for e in entries
            if e["month"] == month
            and e.get("last_status", "submitted") not in
            ("done", "partial", "failed")]
    pend.sort(key=lambda e: (PRIORITY_RANK.get(e.get("priority",
                                                   "new-leads"), 99),
                             e["ts"]))
    return pend


def mark_status(job_id, status, month=None):
    entries = load_ledger()
    month = month or month_key()
    for e in entries:
        if e["month"] == month and e["job_id"] == job_id:
            e["last_status"] = status
    save_ledger(entries)
