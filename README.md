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
