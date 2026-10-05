#!/usr/bin/env python3
"""Campaign pipeline: ONE job spec fans out to a matrix over candidates.

APPROVAL ECONOMY: a single push (one user approval) covers dozens of
candidates end-to-end. Each matrix shard runs the FULL S0->S4 battery
for its candidate(s) with early stop (battery/campaign.py). Per-stage
dispatches (one approval per stage) are banned — the stages are chained
inside the job.

  pipeline.py run --project mix-sbox-search --candidates c1,c2,...
      --image ghcr.io/OWNER/fanout-mix-sbox-search:TAG [--confirm]
      -> builds one campaign spec, submits once, polls, fetches,
         aggregates -> campaign_summary.json + per-candidate verdicts

  pipeline.py sync-candidates --project P --src DIR
      -> stage VM candidate bundles into battery/candidates/ (local)

  pipeline.py aggregate --project P --job-id J --dest DIR
      -> fetch + run battery/aggregate.py locally

Rules:
  * no GHA job exceeds 5h (timeout_minutes=300 hard cap);
  * S3 on GHA stops at 1GB (campaign.py --max-bytes); breaking-point
    runs (>1GB) stay on the VM;
  * candidates must be committed under
    projects/<project>/battery/candidates/ before the run.
"""
import argparse
import datetime
import json
import math
import os
import shutil
import subprocess
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import fanout as F
import budget as budget_mod

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
MAX_CONCURRENT = 20
CAMPAIGN_TIMEOUT = 300  # 5h hard cap


def project_battery_dir(project):
    return os.path.join(REPO_ROOT, "projects", project, "battery")


def campaign_command(project, only):
    cands = os.path.join("projects", project, "battery", "candidates")
    cmd = (f"python3 projects/{project}/battery/campaign.py "
           f"--candidates-dir {cands} --out $OUT_DIR")
    if only:
        cmd += f" --only {','.join(only)}"
    return cmd


def campaign_env(project):
    return (f"BATTERY_HOME=/work/projects/{project}/battery\n"
            f"FLOW_DIR=/work/projects/{project}/flow")


def build_campaign_spec(project, campaign_id, candidates, image,
                        per_shard, per_candidate_min, priority):
    n = len(candidates)
    shards = max(1, min(MAX_CONCURRENT, math.ceil(n / per_shard)))
    est = shards * per_shard * per_candidate_min
    return {
        "project": project,
        "job_id": campaign_id,
        "image": image,
        "command": campaign_command(project, candidates),
        "shards": shards,
        "timeout_minutes": CAMPAIGN_TIMEOUT,
        "env": campaign_env(project),
        "cache_key": "third-party",
        "cache_paths": (f"projects/{project}/battery/"
                        f"third_party/install"),
        "registry_user": "",
        "registry_password": "",
        "priority": priority,
        "est_min": est,
        "per_shard_min": per_shard * per_candidate_min,
        "campaign": campaign_id,
        "candidates": candidates,
        "submitted_at": datetime.datetime.now(
            datetime.timezone.utc).isoformat(),
    }


def cmd_sync_candidates(a):
    """Copy VM candidate bundles into battery/candidates/ (local staging)."""
    dest = os.path.join(project_battery_dir(a.project), "candidates")
    os.makedirs(dest, exist_ok=True)
    new, updated = [], []
    for name in sorted(os.listdir(a.src)):
        s = os.path.join(a.src, name)
        d = os.path.join(dest, name)
        if not os.path.isdir(s):
            continue
        for f in ("impl.c", "model.py", "meta.json"):
            if not os.path.exists(os.path.join(s, f)):
                sys.exit(f"bundle {name}: missing {f} in {a.src}")
        (new if not os.path.exists(d) else updated).append(name)
        shutil.copytree(s, d, dirs_exist_ok=True)
    print(f"staged into {dest}: {len(new)} new, {len(updated)} updated")
    for n in new:
        print(f"  new: {n}")
    print("next: commit + push these through the parent, then run the "
          "campaign")


def cmd_run(a):
    owner, repo = F.split_repo(a.repo)
    project = F.sanitize(a.project)
    candidates = [c.strip() for c in a.candidates.split(",") if c.strip()]
    if not candidates:
        sys.exit("no candidates given")
    cdir = os.path.join(project_battery_dir(project), "candidates")
    missing = [c for c in candidates
               if not os.path.isdir(os.path.join(cdir, c))]
    if missing:
        sys.exit(f"candidates not staged in {cdir}: {missing}\n"
                 f"run sync-candidates first, then commit+push")

    ts = datetime.datetime.now(datetime.timezone.utc).strftime(
        "%Y%m%d-%H%M%S")
    campaign_id = F.sanitize(a.campaign_id or f"campaign-{ts}")
    workdir = os.path.abspath(a.workdir or
                              os.path.join("campaign_runs", campaign_id))
    os.makedirs(workdir, exist_ok=True)

    spec = build_campaign_spec(
        project, campaign_id, candidates, a.image, a.per_shard,
        a.per_candidate_min, a.priority)
    print(f"campaign {campaign_id}: {len(candidates)} candidates, "
          f"{spec['shards']} shards, est ~{spec['est_min']:.0f} min")
    F.submit_spec(owner, repo, a.branch, spec, a.priority,
                  spec["est_min"], a.cap, a.confirm)
    data = F.poll_until_terminal(owner, repo, project, campaign_id,
                                 a.poll_timeout)
    dest = os.path.join(workdir, "results")
    F.fetch_results(owner, repo, project, campaign_id, dest)
    agg = os.path.join(project_battery_dir(project), "aggregate.py")
    r = subprocess.run(
        [sys.executable, agg, "--shards-dir", dest, "--out", dest],
        capture_output=True, text=True)
    print(r.stdout)
    if r.returncode != 0:
        sys.exit(f"aggregate failed:\n{r.stderr[-3000:]}")
    summary = json.load(open(os.path.join(dest, "campaign_summary.json")))
    # per-candidate result JSON {candidate, stages, final verdict, signal}
    per_candidate = {}
    for name, c in summary["campaigns"].items():
        st = c["stages"].get(c["final_stage"], {})
        per_candidate[name] = {
            "candidate": name,
            "final_stage": c["final_stage"],
            "final_verdict": c["final_verdict"],
            "signal": st.get("signal", "").splitlines()[0]
                      if isinstance(st.get("signal"), str) else "",
            "survived": c["survived"],
            "stages": {s: {"verdict": v["verdict"],
                           "signal": v.get("signal", "")}
                       for s, v in c["stages"].items()},
            "job_id": campaign_id,
        }
    json.dump(per_candidate,
              open(os.path.join(dest, "per_candidate.json"), "w"),
              indent=1)
    print(f"\ncampaign {campaign_id} done: "
          f"{summary['n_survivors']}/{summary['n_total']} survived: "
          f"{','.join(summary['survivors'])}")
    print(f"kill histogram: {summary['kill_at_stage']}")
    print(f"results: {dest}")


def cmd_aggregate(a):
    owner, repo = F.split_repo(a.repo)
    dest = os.path.abspath(a.dest)
    F.fetch_results(owner, repo, a.project, a.job_id, dest)
    agg = os.path.join(project_battery_dir(a.project), "aggregate.py")
    r = subprocess.run(
        [sys.executable, agg, "--shards-dir", dest, "--out", dest]
        + a.script_args, capture_output=True, text=True)
    print(r.stdout)
    if r.returncode != 0:
        sys.exit(f"aggregate script failed:\n{r.stderr[-3000:]}")


def main():
    ap = argparse.ArgumentParser(prog="pipeline.py")
    sub = ap.add_subparsers(dest="cmd", required=True)

    p = sub.add_parser("run")
    p.add_argument("--project", required=True)
    p.add_argument("--candidates", required=True,
                   help="comma-separated candidate names")
    p.add_argument("--image", required=True)
    p.add_argument("--per-shard", type=int, default=1,
                   help="candidates per shard (worst ~3h each; keep the "
                        "shard under the 5h cap)")
    p.add_argument("--per-candidate-min", type=float, default=60,
                   help="budget estimate per candidate (expected, with "
                        "early stop; default 60)")
    p.add_argument("--priority", default="new-leads",
                   choices=budget_mod.PRIORITIES)
    p.add_argument("--cap", type=int, default=None)
    p.add_argument("--repo", default=F.DEFAULT_REPO)
    p.add_argument("--branch", default=F.DEFAULT_BRANCH)
    p.add_argument("--campaign-id", default="")
    p.add_argument("--workdir", default="")
    p.add_argument("--poll-timeout", type=int, default=400,
                   help="minutes to wait for the campaign")
    p.add_argument("--confirm", action="store_true")
    p.set_defaults(fn=cmd_run)

    p = sub.add_parser("sync-candidates")
    p.add_argument("--project", required=True)
    p.add_argument("--src", required=True)
    p.set_defaults(fn=cmd_sync_candidates)

    p = sub.add_parser("aggregate")
    p.add_argument("--project", required=True)
    p.add_argument("--job-id", required=True)
    p.add_argument("--dest", required=True)
    p.add_argument("--repo", default=F.DEFAULT_REPO)
    p.add_argument("script_args", nargs="*")
    p.set_defaults(fn=cmd_aggregate)

    a = ap.parse_args()
    a.fn(a)


if __name__ == "__main__":
    main()
