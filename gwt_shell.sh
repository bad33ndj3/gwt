# shellcheck shell=bash
# Helper to quickly switch worktrees
# Usage: gwtcd <branch>
# Changes directory to the worktree for <branch>

function gwtcd() {
    if [[ $# -eq 0 ]]; then
        echo "Usage: gwtcd <branch>" >&2
        return 1
    fi
    cd "$(gwt "$1")" || return
}


