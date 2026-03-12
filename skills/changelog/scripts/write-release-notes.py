#!/usr/bin/env python3

import argparse
import re
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
            text = item.group(1)

            text = re.sub(r"\s*\(#\d+\)", "", text)

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

        line = line.strip()

        if line:
            items.append(line)

    return items


def capitalize_first(text):

    if not text:
        return text

    return text[0].upper() + text[1:]


def lint_items(items):

    clean = []

    for item in items:

        item = item.strip()

        item = re.sub(r"\.$", "", item)

        item = capitalize_first(item)

        clean.append(item)

    return clean


def detect_item_category(item, categories):

    item_lower = item.lower()

    for cat, commits in categories.items():

        for commit in commits:

            words = commit.lower().split()

            if any(w in item_lower for w in words):
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
        raise RuntimeError("Release notes must contain at least one change.")

    text = intro.strip() + "\n\n"

    for item in items:
        text += f"- {item}\n"

    text = text.strip()

    if len(text) > MAX_CHARS:
        raise RuntimeError(
            f"Release note section exceeds {MAX_CHARS} characters."
        )

    return text


def main():

    parser = argparse.ArgumentParser()

    parser.add_argument("--intro-id", required=True)
    parser.add_argument("--intro-en", required=True)

    parser.add_argument("--items-id", required=True)
    parser.add_argument("--items-en", required=True)

    args = parser.parse_args()

    if not CHANGELOG.exists():
        raise FileNotFoundError("CHANGELOG.md not found")

    changelog_text = CHANGELOG.read_text()

    latest = latest_release_block(changelog_text)

    categories = extract_changelog_by_category(latest)

    items_id = normalize_bullets(args.items_id)
    items_en = normalize_bullets(args.items_en)

    items_id = lint_items(items_id)
    items_en = lint_items(items_en)

    items_id = validate_against_changelog(items_id, categories)
    items_en = validate_against_changelog(items_en, categories)

    if not items_id:
        raise RuntimeError("Indonesian release notes must contain at least one valid change.")

    if not items_en:
        raise RuntimeError("English release notes must contain at least one valid change.")

    items_id = rank_items(items_id, categories)
    items_en = rank_items(items_en, categories)

    section_id = build_section(args.intro_id, items_id)
    section_en = build_section(args.intro_en, items_en)

    content = f"""# Release Notes

## Bahasa Indonesia
{section_id}

## English
{section_en}
"""

    OUTPUT.write_text(content)

    print(f"RELEASE_NOTES generated/updated: {OUTPUT}")


if __name__ == "__main__":
    main()