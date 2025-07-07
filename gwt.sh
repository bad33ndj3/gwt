#!/usr/bin/env bash
# shellcheck shell=bash
set -euo pipefail

# gwt - change to the worktree for the given branch
# Source this file in your shell configuration to use the function.

gwt() {
    if [[ $# -ne 1 ]]; then
        echo "Usage: gwt <branch>" >&2
        return 1
    fi

    if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
        echo "Error: Not a git repository" >&2
        return 1
    fi

    local branch="$1"
    local path
    path=$(git worktree list --porcelain | awk -v target="$branch" '
        /^worktree / {
            path = substr($0, 10);
            getline; getline;
            br = substr($0, 8);
            sub(/^refs\/heads\//, "", br);
            if (br == target) {
                print path;
                found = 1;
                exit;
            }
        }
        END { if (!found) exit 1 }
    ') || {
        echo "Error: no worktree for branch $branch" >&2
        return 1
    }

    cd "$path" || return
}

# Allow execution as a script for backward compatibility
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
    gwt "$@"
fi
