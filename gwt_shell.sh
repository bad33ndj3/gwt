# Helper to quickly switch worktrees
# Usage: gwtcd <branch>
# Changes directory to the worktree for <branch>

function gwtcd() {
    cd "$(gwt switch "$1")"
}

# Alias for listing worktrees
alias gwtls='gwt list'

