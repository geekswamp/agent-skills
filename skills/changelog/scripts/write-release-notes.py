#!/usr/bin/env python3

import re
import sys
from pathlib import Path

CHANGELOG = Path("CHANGELOG.md")
OUTPUT = Path("RELEASE_NOTES")

MAX_CHARS = 2500


def latest_release_block(text: str):
    parts = re.split(r"\n## ", text)
    if len(parts) < 2:
        raise RuntimeError("No release section found in CHANGELOG.md")
    return "## " + parts[1]


def extract_changelog_by_category(block: str):
    categories = {}
    current = None

    for line in block.splitlines():

        header = re.match(r"^###\s+(.*)", line)
        if header:
            current = header.group(1).strip().lower()
            categories[current] = []
            continue

        item = re.match(r"^- (.+)", line)
        if item and current:
            text = re.sub(r"\s*\(#\d+\)", "", item.group(1))
            categories[current].append(text)

    return categories


def normalize_bullets(text: str):
    items = []

    for line in text.splitlines():
        line = line.strip()
        if not line:
            continue

        line = re.sub(r"^\d+\.\s*", "", line)
        line = re.sub(r"^[-*]\s*", "", line)

        if line:
            items.append(line)

    return items


def lint_items(items):
    clean = []

    for item in items:
        item = item.strip()
        item = re.sub(r"\.$", "", item)
        item = item[0].upper() + item[1:] if item else item
        clean.append(item)

    return clean


def detect_item_category(item, categories):
    item_lower = item.lower()

    for cat, commits in categories.items():
        for commit in commits:
            if any(w in item_lower for w in commit.lower().split()):
                return cat

    return "other"


def rank_items(items, categories):
    priority = {
        "breaking": 1,
        "added": 2,
        "changed": 3,
        "fixed": 4,
        "reverted": 5,
        "other": 6
    }

    ranked = []

    for item in items:
        cat = detect_item_category(item, categories)
        ranked.append((priority.get(cat, 6), item))

    ranked.sort(key=lambda x: x[0])

    return [i[1] for i in ranked]


def validate_against_changelog(items, categories):
    """
    Validates items against CHANGELOG entries by checking for shared words.
    Only applied to English items — non-English translations are skipped
    because CHANGELOG.md is written in English and word matching would fail
    for other languages.
    """
    changelog_items = []

    for commits in categories.values():
        changelog_items.extend(commits)

    changelog_lower = [c.lower() for c in changelog_items]

    valid = []

    for item in items:
        item_lower = item.lower()

        if any(word in item_lower for c in changelog_lower for word in c.split()):
            valid.append(item)

    return valid


def build_section(intro, items):
    if not items:
        raise RuntimeError("Each language must contain at least one valid change.")

    text = intro.strip() + "\n\n"

    for item in items:
        text += f"- {item}\n"

    text = text.strip()

    if len(text) > MAX_CHARS:
        raise RuntimeError("Section exceeds max character limit")

    return text


# ✅ NEW: parsing per language block
def parse_languages(argv):

    langs = []
    current = None

    i = 0
    while i < len(argv):

        arg = argv[i]

        if arg == "--lang":
            current = {
                "code": argv[i + 1],
                "intro": "",
                "items": ""
            }
            langs.append(current)
            i += 2
            continue

        if arg == "--intro":
            if not current:
                raise RuntimeError("--intro must come after --lang")
            current["intro"] = argv[i + 1]
            i += 2
            continue

        if arg == "--items":
            if not current:
                raise RuntimeError("--items must come after --lang")
            current["items"] = argv[i + 1]
            i += 2
            continue

        i += 1

    if not langs:
        raise RuntimeError("At least one --lang block is required")

    return langs


def main():

    if not CHANGELOG.exists():
        raise FileNotFoundError("CHANGELOG.md not found")

    langs_input = parse_languages(sys.argv[1:])

    changelog_text = CHANGELOG.read_text()
    latest = latest_release_block(changelog_text)
    categories = extract_changelog_by_category(latest)

    content = "# Release Notes\n\n"

    for lang_data in langs_input:

        code = lang_data["code"]
        intro = lang_data["intro"]
        items_raw = lang_data["items"]

        items = normalize_bullets(items_raw)
        items = lint_items(items)

        # Validation against CHANGELOG is English-only.
        # Non-English translations won't share words with the English CHANGELOG,
        # so we trust the agent has correctly translated the items.
        if code.lower() == "en":
            items = validate_against_changelog(items, categories)

        if not items:
            raise RuntimeError(f"{code} must contain at least one valid change")

        items = rank_items(items, categories)

        section = build_section(intro, items)

        content += f"## {code.upper()}\n{section}\n\n"

    OUTPUT.write_text(content.strip())

    print(f"RELEASE_NOTES generated: {OUTPUT}")


if __name__ == "__main__":
    main()