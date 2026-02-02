#!/usr/bin/env bash

# Conventional Commits regex
# type(scope): description
COMMIT_REGEX="^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([a-z0-9-]+\))?!?: .+"

COMMIT_MSG_FILE=$1
COMMIT_MSG=$(cat "$COMMIT_MSG_FILE")

if [[ ! $COMMIT_MSG =~ $COMMIT_REGEX ]]; then
    echo "❌ ERROR: Invalid commit message format."
    echo "Expected: <type>(<scope>): <description>"
    echo "Allowed types: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert"
    echo "Example: feat(auth): add JWT validation"
    exit 1
fi

echo "✅ Commit message follows conventional commits specification."
exit 0
