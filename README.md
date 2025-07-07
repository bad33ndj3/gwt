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

## Shell helper

To quickly switch worktrees from your shell, a helper script is provided. Append
it to your shell configuration so the `gwtcd` function and `gwtls` alias are
available. If you have the repository cloned, you can append the file directly:

```bash
cat gwt_shell.sh >> ~/.bashrc   # or ~/.zshrc

# If you installed with `go install` and don't have the file locally, fetch it
# from the repository instead:
curl -fsSL https://raw.githubusercontent.com/bad33ndj3/gwt/main/gwt_shell.sh >> ~/.bashrc   # or ~/.zshrc
```

Reload your shell or source the file to start using it.

The helper defines:

```bash
gwtcd <branch>   # cd into the worktree for <branch>
gwtls            # shorthand for 'gwt list'
```
