# gwt

`gwt` is a tiny helper for managing Git worktrees. It can list existing worktrees,
print the path of a worktree for quick switching, and add new worktrees.

```
Usage:
  gwt list
  gwt switch <branch>
  gwt <branch> <path>
```

Use `cd $(gwt switch <branch>)` to change directories quickly.

## Installation

You can install the command directly from the repository using `go install`:

```bash
go install github.com/bad33ndj3/gwt@latest
```

Make sure `$GOBIN` (or `$GOPATH/bin`) is in your `PATH` so the `gwt` binary is
available.
