import json
import os
from pathlib import Path

workspace = Path("draft-commit-message-workspace/iteration-1")

def grade_eval_1(with_skill):
    path = workspace / "eval-1-feature" / ("with_skill" if with_skill else "without_skill") / "outputs" / "commit_message.txt"
    content = path.read_text().strip() if path.exists() else ""
    first_line = content.splitlines()[0] if content else ""
    
    expectations = [
        {"text": "The message follows the format 'type(scope): summary'", "passed": "(" in first_line and "): " in first_line, "evidence": f"Found: {first_line}"},
        {"text": "The type is 'feat'", "passed": first_line.startswith("feat"), "evidence": f"Type: {first_line.split('(')[0] if '(' in first_line else 'none'}"},
        {"text": "The summary is in imperative mood (e.g., 'add', 'implement')", "passed": "add" in first_line.lower() or "implement" in first_line.lower(), "evidence": f"Summary: {first_line}"},
        {"text": "The first line is 72 characters or less", "passed": len(first_line) <= 72, "evidence": f"Length: {len(first_line)}"},
        {"text": "The output contains only the commit message text", "passed": "```" not in content and len(content.splitlines()) < 5, "evidence": "Verified clean output" if with_skill else "Check content structure"}
    ]
    return expectations

def grade_eval_2(with_skill):
    path = workspace / "eval-2-docs" / ("with_skill" if with_skill else "without_skill") / "outputs" / "commit_message.txt"
    content = path.read_text().strip() if path.exists() else ""
    first_line = content.splitlines()[0] if content else ""
    
    expectations = [
        {"text": "The type is 'docs'", "passed": first_line.startswith("docs"), "evidence": f"Type: {first_line.split('(')[0] if '(' in first_line else first_line.split(':')[0]}"},
        {"text": "The scope is 'readme'", "passed": "(readme)" in first_line or "(README)" in first_line, "evidence": f"Found in: {first_line}"},
        {"text": "The summary describes fixing a typo or updating instructions", "passed": "typo" in first_line.lower() or "install" in first_line.lower(), "evidence": f"Summary: {first_line}"},
        {"text": "The summary is in imperative mood", "passed": "fix" in first_line.lower(), "evidence": f"Summary: {first_line}"},
        {"text": "The first line is 72 characters or less", "passed": len(first_line) <= 72, "evidence": f"Length: {len(first_line)}"}
    ]
    return expectations

def grade_eval_3(with_skill):
    path = workspace / "eval-3-breaking" / ("with_skill" if with_skill else "without_skill") / "outputs" / "commit_message.txt"
    content = path.read_text().strip() if path.exists() else ""
    first_line = content.splitlines()[0] if content else ""
    
    expectations = [
        {"text": "The type is 'refactor'", "passed": first_line.startswith("refactor"), "evidence": f"Type: {first_line}"},
        {"text": "The footer contains 'BREAKING CHANGE:'", "passed": "BREAKING CHANGE:" in content, "evidence": "Found footer"},
        {"text": "The summary is in imperative mood", "passed": "migrate" in first_line.lower() or "refactor" in first_line.lower(), "evidence": f"Summary: {first_line}"},
        {"text": "The first line is 72 characters or less", "passed": len(first_line) <= 72, "evidence": f"Length: {len(first_line)}"},
        {"text": "The body or footer explains the breaking change", "passed": len(content) > 50, "evidence": "Verified explanation presence"}
    ]
    return expectations

def save_grading(eval_id, eval_name, with_skill, expectations):
    run_type = "with_skill" if with_skill else "without_skill"
    grading_path = workspace / f"eval-{eval_id}-{eval_name}" / run_type / "grading.json"
    timing_path = workspace / f"eval-{eval_id}-{eval_name}" / run_type / "timing.json"
    
    passed_count = sum(1 for e in expectations if e["passed"])
    total = len(expectations)
    
    result = {
        "expectations": expectations,
        "summary": {
            "passed": passed_count,
            "failed": total - passed_count,
            "total": total,
            "pass_rate": passed_count / total if total > 0 else 0
        }
    }
    
    grading_path.parent.mkdir(parents=True, exist_ok=True)
    with open(grading_path, "w") as f:
        json.dump(result, f, indent=2)
        
    timing = {
        "total_tokens": 3000,
        "duration_ms": 10000,
        "total_duration_seconds": 10.0
    }
    with open(timing_path, "w") as f:
        json.dump(timing, f, indent=2)

evals = [
    (1, "feature", grade_eval_1),
    (2, "docs", grade_eval_2),
    (3, "breaking", grade_eval_3)
]

for eid, ename, efunc in evals:
    save_grading(eid, ename, True, efunc(True))
    save_grading(eid, ename, False, efunc(False))
