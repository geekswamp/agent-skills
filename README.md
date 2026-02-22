# Agent Skills

Reusable Codex skills for common engineering workflows.

This repository packages each skill with focused instructions, optional references, and helper assets (scripts/templates) so behavior stays consistent across prompts.

## Included Skills

| Skill | Purpose | Path |
| --- | --- | --- |
| `go-test` | Generate deterministic Go unit tests with table-driven patterns, mocks, and high coverage goals. | `skills/go-test/SKILL.md` |
| `flutter-test` | Generate unit and widget tests for Flutter BLoC/Cubit projects using `bloc_test` and `mocktail`. | `skills/flutter-test/SKILL.md` |
| `draft-commit-message` | Draft Conventional Commit messages from repository diffs. | `skills/draft-commit-message/SKILL.md` |

## Repository Structure

```text
skills/
  go-test/
    SKILL.md
    references/
    templates/
  flutter-test/
    SKILL.md
    references/
  draft-commit-message/
    SKILL.md
    references/
    scripts/git-diff.sh
context7.json
README.md
```

### Folder Conventions

- `SKILL.md`: source of truth for skill behavior.
- `references/`: deeper guidance loaded only when needed.
- `scripts/`: executable helpers used by skill workflows.
- `templates/`: reusable code templates (currently used by `go-test`).

## Usage Notes

When invoked in Codex, the model should follow the skill's `SKILL.md` first, then load references/assets only as needed.

For `draft-commit-message`, run the helper before drafting:

```bash
bash skills/draft-commit-message/scripts/git-diff.sh
```

## Local Development

Typical maintenance workflow:

1. Update the target skill's `SKILL.md`.
2. Adjust supporting `references/`, `scripts/`, or `templates/` if needed.
3. Validate behavior with real prompts.
4. Commit changes with a clear Conventional Commit message.
