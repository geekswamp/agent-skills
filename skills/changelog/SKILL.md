---
name: changelog
description: Generate and update project changelog entries from git commit history. Use this skill whenever the user mentions changelogs, release notes, summarizing commits, or preparing for a new version. Trigger it for App Store, Play Store, or web app announcement requests, even if they do not explicitly ask for a 'changelog' file. This skill is also specialized for creating friendly, bilingual release notes for non-technical users.
compatibility: Requires Git, jq, and Python 3.9+
metadata:
  short-description: Generate changelog from git commits
  version: 1.0.0
---

# Generate Changelog
Follow this workflow in order. Do not skip step 1.

## 1. Generate Source Data First
Always run [gen-changelog.sh](./scripts/gen-changelog.sh) script:

```bash
bash scripts/gen-changelog.sh
```

Treat this command output as the single source of truth for:
- `last_tag`
- `next_version`
- `date`
- categorized commit list

If output is `No changes since last release.`, report no changelog update is needed and stop.

## 2. Use the Changelog Template
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

Remove empty sections so the output does not include blank category headings, and ensure each changelog entry starts with a capital letter.

## 3. Update CHANGELOG.md
If `CHANGELOG.md` does not exist, create it from the rendered template content.

If `CHANGELOG.md` already exists:
- Keep the existing title and intro paragraph.
- Insert the new release section directly below the intro, before older releases.
- Preserve older release entries below the new one.

## 4. Verify Final Result
Before finishing, review the generated changelog to ensure the following:

- The version number and release date exactly match the values returned by the script output.
- Section headings appear in the correct order: **Added**, **Breaking Changes**, **Changed**, **Fixed**, **Reverted**.
- Empty sections are completely removed so no unused category headings remain.
- Each changelog entry starts with a capital letter.
- Each entry is slightly improved from the original commit message to be clearer and more informative for end users.
- Avoid raw or technical commit-style messages; rewrite them into short, user-friendly descriptions when necessary.
- The Markdown structure is valid, clean, and easy to read.
- All content is written in **English**.

## Additional Guidelines
- When release notes are requested (for example for App Store, Play Store, or announcements), generate them based on the **latest version in `CHANGELOG.md`**.
- Follow the guidelines defined in [`references/release_notes.md`](./references/release_notes.md).
- Only include the **most important user-facing improvements** from that release.
- Ensure the final release notes follow the **structure, tone, and language limits** specified in the reference file.