# Agent Skills

A curated set of reusable Codex skills for common engineering tasks.

This repository contains skill packages that guide Codex with consistent workflows, quality standards, and references.

## Included Skills

### `go-test`
Generate high-quality Go unit tests with table-driven patterns and strong coverage targets.

- Focus: deterministic tests, `testify` assertions/mocks, success/error/edge cases.
- Main file: `skills/go-test/SKILL.md`

### `flutter-test`
Generate high-quality Flutter tests for BLoC/Cubit-based applications.

- Focus: `bloc_test`, `mocktail`, unit + widget test separation, behavior-first assertions.
- Main file: `skills/flutter-test/SKILL.md`

### `draft-commit-message`
Draft conventional commit messages from your repository changes.

- Focus: Conventional Commits format, concise subject line, optional breaking-change footer.
- Includes helper script: `skills/draft-commit-message/scripts/git-diff.sh`
- Main file: `skills/draft-commit-message/SKILL.md`

## How Skills Are Used

Each skill is defined by a `SKILL.md` file and can optionally include:

- `agents/`: provider-specific metadata (display name, short description).
- `references/`: deeper guidance loaded only when needed.
- `scripts/`: helper utilities used by the skill workflow.

When a skill is invoked in Codex, the model follows `SKILL.md` as the source of truth for behavior and output style.

## Local Development

Use this repository to author and maintain skills, then make them available to your Codex setup.

Typical workflow:

1. Edit `SKILL.md` (and related references/scripts).
2. Validate instructions by running the skill in real prompts.
3. Version changes in Git.

## Requirements

- Git (required for `draft-commit-message` diff workflow).
- A Codex environment that supports custom/local skills.