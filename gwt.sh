#!/usr/bin/env bash
set -euo pipefail

usage() {
    cat <<USAGE
Usage: gwt <branch>
USAGE
}



find_worktree() {
    local branch="$1"
    git worktree list --porcelain | awk -v target="$branch" '
        /^worktree / {
            path = substr($0, 10);
            getline; getline;
            br = substr($0, 8);
            sub(/^refs\/heads\//, "", br);
            if (br == target) {
                print path;
                found=1;
                exit;
            }
        }
        END { if (!found) exit 1 }
    '
}

main() {
    if [[ $# -ne 1 ]]; then
        usage >&2
        exit 1
    fi

    if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
        echo "Error: Not a git repository" >&2
        exit 1
    fi

    if path=$(find_worktree "$1"); then
        echo "$path"
    else
        echo "Error: no worktree for branch $1" >&2
        exit 1
    fi
}

main "$@"
