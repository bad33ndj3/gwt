#!/usr/bin/env bash
# shellcheck shell=bash
# Enable strict mode only when executed directly, not when sourced
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
    set -euo pipefail
fi

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

# Autocomplete branch names with existing worktrees
__gwt_complete() {
    local cur branches
    cur="${COMP_WORDS[COMP_CWORD]}"
    if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
        branches=$(git worktree list --porcelain 2>/dev/null | awk '/^branch /{sub("^refs\/heads/","",$2); print $2}')
        COMPREPLY=( $(compgen -W "$branches" -- "$cur") )
    else
        COMPREPLY=()
    fi
}

complete -o default -F __gwt_complete gwt

# Allow execution as a script for backward compatibility
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
    gwt "$@"
fi
