---
name: draft-commit-message
description: Draft a conventional commit message when the user asks for help writing a commit message.
compatibility: Requires git
metadata:
  short-description: Draft an informative commit message.
---

Draft a conventional commit message that matches the change summary provided by the user.

## Requirements
- Use the Conventional Commits format: `type(scope): summary`.
- Use the imperative mood in the summary (for example, `feat`, `fix`, `refactor`).
- The supported types are `bump`, `feat`, `fix`, `docs`, `refactor`, `test`, `ci`, `chore`, `perf`, and `revert`.
- The **entire first line (including type and scope)** must not exceed 72 characters.
- Do not wrap the summary line.
- If there are breaking changes, include a `BREAKING CHANGE:` footer.
- Always use English.

## Script
- Always run `scripts/git-diff.sh` (or `bash scripts/git-diff.sh`) before generating the commit message.
- Use the script output as the primary source of truth for determining the commit type and summary.
- Run `scripts/git-diff.sh` or `bash scripts/git-diff.sh` to show both unstaged and staged full diffs.
- Pass optional file paths or flags as args, e.g. `scripts/git-diff.sh <path>` or `bash scripts/git-diff.sh <path>`.
- Output order is unstaged diff first, then staged diff; add separators if needed.
- Prefer analyzing staged changes when drafting the commit message.
- If the script fails, fallback to `git --no-pager diff --cached`.

## When to load references
- Detailed technical reference: [REFERENCE.md](./references/REFERENCE.md).
