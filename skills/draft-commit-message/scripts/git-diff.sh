#!/usr/bin/env bash

set -euo pipefail
export GIT_PAGER=cat

ARGS="$@"

section() {
    echo "<<<SECTION:$1>>>"
}

end_section() {
    echo "<<<END:$1>>>"
    echo
}

section "STATUS"
git --no-pager status --short $ARGS || true
end_section "STATUS"

section "DIFF_STAT_STAGED"
if git diff --cached --quiet $ARGS; then
    echo "NO_CHANGES"
else
    git --no-pager diff --cached --stat $ARGS
fi
end_section "DIFF_STAT_STAGED"

section "DIFF_STAT_UNSTAGED"
if git diff --quiet $ARGS; then
    echo "NO_CHANGES"
else
    git --no-pager diff --stat $ARGS
fi
end_section "DIFF_STAT_UNSTAGED"

section "DIFF_STAGED"
git --no-pager diff --cached $ARGS || true
end_section "DIFF_STAGED"

section "DIFF_UNSTAGED"
git --no-pager diff $ARGS || true
end_section "DIFF_UNSTAGED"

section "UNTRACKED_FILES"
git ls-files --others --exclude-standard $ARGS || true
end_section "UNTRACKED_FILES"
