# gwt

`gwt` is a tiny **shell script** for managing Git worktrees. It prints the path of a worktree
for quick switching.

```
Usage:
  gwt <branch>
```

Use `cd $(gwt <branch>)` to change directories quickly.

## Installation

Download the script somewhere on your `PATH` and make it executable. Create the
directory if it doesn't already exist:

```bash
mkdir -p ~/bin
curl -fsSL https://raw.githubusercontent.com/bad33ndj3/gwt/main/gwt.sh -o ~/bin/gwt
chmod +x ~/bin/gwt
```

You can also place it directly in a system directory:

```bash
curl -fsSL https://raw.githubusercontent.com/bad33ndj3/gwt/main/gwt.sh \
  -o /usr/local/bin/gwt
chmod +x /usr/local/bin/gwt
```

Adjust the destination path as needed and ensure the directory is in your `PATH`.

## Shell helper

To quickly switch worktrees from your shell, a helper script is provided. Append
it to your shell configuration so the `gwtcd` function is available. If you have the repository cloned, you can append the file directly:

```bash
cat gwt_shell.sh >> ~/.bashrc   # or ~/.zshrc

# If you downloaded only the script, fetch the helper from the repository:
curl -fsSL https://raw.githubusercontent.com/bad33ndj3/gwt/main/gwt_shell.sh >> ~/.bashrc   # or ~/.zshrc
```

Make sure the `gwt` script is installed and on your `PATH` before sourcing
`gwt_shell.sh`. If you saved the script as `gwt.sh`, rename it to `gwt` so the
helper can find it.

Reload your shell or source the file to start using it.

The helper defines:

```bash
gwtcd <branch>   # cd into the worktree for <branch>
```
