import json
from pathlib import Path

evals_source = json.loads(Path("changelog/evals/evals.json").read_text())
expectations_map = {e["id"]: e["expectations"] for e in evals_source["evals"]}

workspace = Path("changelog-workspace/iteration-1")
for eval_dir in workspace.iterdir():
    if not eval_dir.is_dir(): continue
    metadata_path = eval_dir / "eval_metadata.json"
    if metadata_path.exists():
        metadata = json.loads(metadata_path.read_text())
        eid = metadata.get("eval_id")
        if eid in expectations_map:
            metadata["assertions"] = expectations_map[eid]
            metadata_path.write_text(json.dumps(metadata, indent=2))
            print(f"Updated {metadata_path}")
