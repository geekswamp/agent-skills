#!/usr/bin/env bash
set -e

LAST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || true)

if [ -z "$LAST_TAG" ]; then
    LAST_TAG="v0.0.0"
    COMMITS=$(git log --pretty=format:"%s")
else
    COMMITS=$(git log "$LAST_TAG"..HEAD --pretty=format:"%s")
fi

if [ -z "$COMMITS" ]; then
    echo "No changes since last release."
    exit 0
fi

VERSION=${LAST_TAG#v}
IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION"

BUMP_MAJOR=0
BUMP_MINOR=0
BUMP_PATCH=0

COMMITS_JSON="[]"

while IFS= read -r COMMIT; do
    TYPE=$(echo "$COMMIT" | cut -d':' -f1)

    CATEGORY="other"
    BREAKING=false

    case "$TYPE" in
        feat* )
        CATEGORY="added"
        BUMP_MINOR=1
        ;;
    fix* )
        CATEGORY="fixed"
        BUMP_PATCH=1
        ;;
    perf* )
        CATEGORY="changed"
        BUMP_PATCH=1
        ;;
    refactor* )
        CATEGORY="changed"
        BUMP_PATCH=1
        ;;
    revert* )
        CATEGORY="reverted"
        BUMP_PATCH=1
        ;;
    bump*|test*|ci*|chore*|docs* )
        CATEGORY="ignore"
        ;;
esac

    if [ "$CATEGORY" = "ignore" ]; then
        continue
    fi

    if [[ "$COMMIT" =~ ^[a-zA-Z]+! ]] || [[ "$COMMIT" == *"BREAKING CHANGE"* ]]; then
        BREAKING=true
        BUMP_MAJOR=1
        CATEGORY="breaking"
    fi

    PR=$(echo "$COMMIT" | grep -oE '#[0-9]+' | head -n1 | tr -d '#' || true)

    MESSAGE=$(echo "$COMMIT" \
    | sed -E 's/^[^:]+: //' \
    | sed -E 's/\(#?[0-9]+\)//g' \
    | xargs)

    COMMITS_JSON=$(echo "$COMMITS_JSON" | jq \
        --arg category "$CATEGORY" \
        --arg message "$MESSAGE" \
        --arg pr "$PR" \
        --argjson breaking "$BREAKING" \
        '. + [{
            category:$category,
            message:$message,
            pr:$pr,
            breaking:$breaking
        }]')

done <<< "$COMMITS"

if [ "$BUMP_MAJOR" -eq 1 ]; then
    MAJOR=$((MAJOR+1))
    MINOR=0
    PATCH=0
elif [ "$BUMP_MINOR" -eq 1 ]; then
    MINOR=$((MINOR+1))
    PATCH=0
elif [ "$BUMP_PATCH" -eq 1 ]; then
    PATCH=$((PATCH+1))
fi

NEXT_VERSION="v$MAJOR.$MINOR.$PATCH"
DATE=$(date +%Y-%m-%d)

jq -n \
    --arg last_tag "$LAST_TAG" \
    --arg next_version "$NEXT_VERSION" \
    --arg date "$DATE" \
    --argjson commits "$COMMITS_JSON" \
    '{
        last_tag: $last_tag,
        next_version: $next_version,
        date: $date,
        commits: $commits
    }'