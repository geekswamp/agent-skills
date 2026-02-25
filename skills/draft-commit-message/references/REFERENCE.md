# Allowed Type Values
- `feat`: new feature for the user, not a new feature for build script.
- `fix`: bug fix for the user, not a fix to a build script.
- `docs`: changes to the documentation (for example, `README.md`).
- `refactor`: refactoring production code, e.g. renaming a variable.
- `test`: adding missing tests, refactoring tests; no production code change.
- `chore`: updating grunt tasks etc. no production code change.
- `ci`: changes to CI configuration files or scripts.
- `perf`: performance improvements.
- `revert`: reverting a previous commit.
- `bump`: update/increment version of dependencies.

## Invocation Guide
Run the helper script from repository root before drafting:

```bash
bash skills/draft-commit-message/scripts/git-diff.sh
```

Optional scoped run:

```bash
bash skills/draft-commit-message/scripts/git-diff.sh path/to/file_or_dir
```

## Script Output Sections
The script emits structured sections:
- `STATUS`
- `DIFF_STAT_STAGED`
- `DIFF_STAT_UNSTAGED`
- `DIFF_STAGED`
- `DIFF_UNSTAGED`
- `UNTRACKED_FILES`

Selection priority:
1. `DIFF_STAGED` and `DIFF_STAT_STAGED`
2. Fallback to unstaged sections when nothing is staged

## Commit Message Output Format
Use Conventional Commit format:

```text
type(scope): summary
```

Rules:
- Use one of allowed `type` values.
- Keep first line <=72 chars.
- Use imperative mood (`add`, `fix`, `refactor`, not `added`/`fixed`).
- Keep scope short and meaningful (`auth`, `api`, `deps`, `test`).

Extended format when needed:

```text
type(scope): summary

- bullet point 1
- bullet point 2

BREAKING CHANGE: describe migration impact
```

If there is no breaking change, omit the footer.

## Concrete Output Examples
```text
fix(auth): handle token refresh failure on startup
```

```text
test(auth): add bloc tests for login error transitions
```

```text
refactor(api): split user sync into dedicated service

- extract sync logic from handler
- keep API behavior unchanged
```

```text
feat(config): replace env var names for runtime settings

BREAKING CHANGE: rename APP_PORT to SERVER_PORT
```

## Conventional Branch Reference
- A specification for adding human and machine-readable meaning to branch: [conventional_branch.md](./conventional_branch.md).
