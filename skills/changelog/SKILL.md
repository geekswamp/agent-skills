---
name: changelog
description: Generate and update project changelog entries from git commit history. Use when asked to create release notes, update CHANGELOG.md, or summarize changes between releases using the local changelog script and template.
compatibility: Requires git
metadata:
  short-description: Generate changelog from git commits
---

# Generate Changelog
Follow this workflow in order. Do not skip step 1.

## 1. Generate source data first
Always run [gen_changelog.sh](./scripts/gen_changelog.sh) script:

```bash
bash scripts/gen_changelog.sh
```

Treat this command output as the single source of truth for:
- `last_tag`
- `next_version`
- `date`
- categorized commit list

If output is `No changes since last release.`, report no changelog update is needed and stop.

## 2. Use the changelog template
Use [CHANGELOG.md](./templates/CHANGELOG.md) as the rendering template for the new release block.

Map data from JSON output:
- `{{version}}` -> `next_version` without leading `v`
- `{{date}}` -> `date`
- `{{#added}}` -> commits with `category == "added"`
- `{{#breaking}}` -> commits with `category == "breaking"` or `breaking == true`
- `{{#changed}}` -> commits with `category == "changed"`
- `{{#fixed}}` -> commits with `category == "fixed"`
- `{{#reverted}}` -> commits with `category == "reverted"`

For each commit line:
- Use `- {{message}} (#{{pr}})` only when `pr` is not empty.
- Use `- {{message}}` when `pr` is empty.

Remove empty sections so output does not include blank category headings.

## 3. Update CHANGELOG.md
If `CHANGELOG.md` does not exist, create it from the rendered template content.

If `CHANGELOG.md` already exists:
- Keep the existing title and intro paragraph.
- Insert the new release section directly below the intro, before older releases.
- Preserve older release entries below the new one.

## 4. Verify final result
Check:
- Version and date match script output exactly
- Headings are in this order: **Added**, **Breaking Changes**, **Changed**, **Fixed**, **Reverted**
- No empty sections are left
- Markdown is valid and readable
- Always use **English**.
