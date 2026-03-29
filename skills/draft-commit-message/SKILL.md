---
name: draft-commit-message
description: Draft a Conventional Commit message from current repo changes. Use when user asks for commit message help based on staged/unstaged git diff.
compatibility: Requires git
metadata:
  short-description: Draft an informative commit message.
  version: 1.0.0
---

Draft a conventional commit message that matches the change summary provided by the user.

## Requirements
- Use the Conventional Commits format: `type(scope): summary`.
- Use the imperative mood in the summary (for example, `feat`, `fix`, `refactor`).
- The supported types are `bump`, `feat`, `fix`, `docs`, `refactor`, `test`, `ci`, `chore`, `perf`, and `revert`.
- The **entire first line (including type and scope)** must not exceed 72 characters.
- Do not wrap the summary line.
- If there are breaking changes, include a `BREAKING CHANGE:` footer.
- Always use **English**.

## Script
- Alway run [git-diff.sh](./scripts/git-diff.sh) script before generating a commit message.
- Run:
  - `bash scripts/git-diff.sh`
  - Optional scoped analysis: `bash scripts/git-diff.sh <path>`
- Use the script output as the primary source of truth for determining the commit type and summary.
- The script shows both staged and unstaged status, diff stats, full diffs, and untracked files.

## Output Contract
- Return commit message text only (no explanation) unless user asks for reasoning.
- Follow the exact output format in [REFERENCE.md](./references/REFERENCE.md):
  - Subject: `type(scope): summary` (<=72 chars)
  - Optional body: concise bullet points
  - Optional footer: `BREAKING CHANGE: ...`