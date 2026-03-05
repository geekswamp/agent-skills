# Agent Skills

Reusable Codex skills for common engineering workflows.

This repository packages each skill with focused instructions, optional references, helper assets (`scripts/`, `templates/`), and agent metadata so behavior stays consistent across prompts.

## Included Skills

| Skill | Purpose | Path |
| --- | --- | --- |
| `go-test` | Generate deterministic Go unit tests using table-driven patterns, mocks, and high coverage targets. | `skills/go-test/SKILL.md` |
| `flutter-test` | Generate unit and widget tests for Flutter BLoC/Cubit projects with `bloc_test` and `mocktail`. | `skills/flutter-test/SKILL.md` |
| `draft-commit-message` | Draft Conventional Commit messages from staged/unstaged repository diffs. | `skills/draft-commit-message/SKILL.md` |
| `changelog` | Generate release-ready changelog entries from git commit history. | `skills/changelog/SKILL.md` |

## Repository Structure

```text
skills/
  changelog/
    SKILL.md
    agents/openai.yaml
    scripts/gen_changelog.sh
    templates/CHANGELOG.md
  draft-commit-message/
    SKILL.md
    agents/openai.yaml
    references/
    scripts/git-diff.sh
  flutter-test/
    SKILL.md
    agents/openai.yaml
    references/
  go-test/
    SKILL.md
    agents/openai.yaml
    references/
    templates/
context7.json
README.md
```

## Folder Conventions

- `SKILL.md`: source of truth for skill behavior and workflow.
- `references/`: deeper guidance that should be loaded only when needed.
- `scripts/`: executable helpers used during skill execution.
- `templates/`: reusable output structures (for example changelog/test scaffolds).
- `agents/openai.yaml`: agent-level execution metadata for the skill.

## Usage Notes

When invoked in Codex, follow each skill's `SKILL.md` first, then load references/assets progressively.

Helper scripts from the repository root:

```bash
# Draft commit message context
bash skills/draft-commit-message/scripts/git-diff.sh

# Optional scoped diff analysis
bash skills/draft-commit-message/scripts/git-diff.sh <path>

# Generate changelog source data (JSON)
bash skills/changelog/scripts/gen_changelog.sh
```

## Local Development

Typical maintenance workflow:

1. Update the target skill's `SKILL.md`.
2. Update supporting `references/`, `scripts/`, `templates/`, or agent metadata.
3. Validate behavior with real prompts and script output.
4. Commit changes with a clear Conventional Commit message.
