# shellcheck shell=bash
# Helper to quickly switch worktrees
# Usage: gwtcd <branch>
# Changes directory to the worktree for <branch>

function gwtcd() {
    cd "$(gwt "$1")" || return
}


