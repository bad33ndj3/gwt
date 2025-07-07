#!/usr/bin/env bash
set -euo pipefail

usage() {
    cat <<USAGE
Usage:
  gwt list
  gwt switch <branch>
  gwt <branch> <path>
USAGE
}

# worktree_list prints `git worktree list` output or uses GWT_WORKTREE_LIST if set
worktree_list() {
    if [[ -n ${GWT_WORKTREE_LIST:-} ]]; then
        printf '%s' "$GWT_WORKTREE_LIST"
    else
        git worktree list --porcelain
    fi
}

list_worktrees() {
    worktree_list | awk '
        /^worktree / {
            path = substr($0, 10);
            getline; getline;
            branch = substr($0, 8);
            sub(/^refs\/heads\//, "", branch);
            printf "%-20s %s\n", branch, path;
        }'
}

find_worktree() {
    local branch="$1"
    worktree_list | awk -v target="$branch" '
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

add_worktree() {
    local branch="$1" path="$2"
    if [[ -d "$path/.git" ]]; then
        echo "worktree already exists at $path" >&2
        return 1
    fi
    echo "Adding worktree for branch $branch at $path"
    git worktree add "$path" "$branch"
}

main() {
    if [[ $# -lt 1 ]]; then
        usage >&2
        exit 1
    fi

    case "$1" in
        list)
            list_worktrees
            ;;
        switch)
            if [[ $# -ne 2 ]]; then
                echo "Usage: gwt switch <branch>" >&2
                exit 1
            fi
            if path=$(find_worktree "$2"); then
                echo "$path"
            else
                echo "Error: no worktree for branch $2" >&2
                exit 1
            fi
            ;;
        *)
            if [[ $# -ne 2 ]]; then
                usage >&2
                exit 1
            fi
            add_worktree "$1" "$2"
            ;;
    esac
}

main "$@"
