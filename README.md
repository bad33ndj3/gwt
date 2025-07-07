# gwt

`gwt` is a shell function for quickly switching Git worktrees.
It changes the current directory to the worktree for a branch.

```
Usage:
  gwt <branch>
```

## Installation

Download the script and source it from your shell configuration:

```bash
curl -fsSL https://raw.githubusercontent.com/bad33ndj3/gwt/main/gwt.sh -o ~/.gwt.sh
```

Then add the following line to your `~/.bashrc` (or `~/.zshrc`):

```bash
source ~/.gwt.sh
```

Reload your shell and run `gwt <branch>` to jump to that worktree.

### Tab Completion

When sourced, `gwt.sh` now installs a Bash completion function. Press `Tab`
after typing `gwt` to see available branches that already have worktrees.
