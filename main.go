// Command gwt manages Git worktrees. It can list, switch, and add
// worktrees using only standard library packages.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		if err := listWorktrees(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "switch":
		if len(os.Args) != 3 {
			fmt.Fprint(os.Stderr, "Usage: gwt switch <branch>\n")
			os.Exit(1)
		}
		path, err := findWorktree(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(path)
	default:
		if len(os.Args) != 3 {
			usage()
			os.Exit(1)
		}
		if err := addWorktree(os.Args[1], os.Args[2]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

func usage() {
	fmt.Fprint(os.Stderr, "Usage:\n  gwt list\n  gwt switch <branch>\n  gwt <branch> <path>\n")
}

// worktreeList returns `git worktree list` output. Tests can override it by
// setting the GWT_WORKTREE_LIST environment variable.
func worktreeList() ([]byte, error) {
	if v := os.Getenv("GWT_WORKTREE_LIST"); v != "" {
		return []byte(v), nil
	}
	return exec.Command("git", "worktree", "list", "--porcelain").Output()
}

// listWorktrees prints available worktrees with branch names for easy
// switching.
func listWorktrees() error {
	out, err := worktreeList()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "worktree ") {
			path := strings.TrimPrefix(line, "worktree ")
			if !scanner.Scan() { // HEAD line
				break
			}
			if !scanner.Scan() { // branch line
				break
			}
			br := strings.TrimPrefix(scanner.Text(), "branch ")
			br = strings.TrimPrefix(br, "refs/heads/")
			fmt.Printf("%-20s %s\n", br, path)
		}
	}
	return scanner.Err()
}

// findWorktree returns the path to the worktree for a given branch.
func findWorktree(branch string) (string, error) {
	out, err := worktreeList()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "worktree ") {
			path := strings.TrimPrefix(line, "worktree ")
			if !scanner.Scan() { // HEAD line
				break
			}
			if !scanner.Scan() { // branch line
				break
			}
			br := strings.TrimPrefix(scanner.Text(), "branch ")
			br = strings.TrimPrefix(br, "refs/heads/")
			if br == branch {
				return path, nil
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no worktree for branch %s", branch)
}

// addWorktree adds a new worktree if it does not already exist.
func addWorktree(branch, path string) error {
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return fmt.Errorf("worktree already exists at %s", path)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking path: %v", err)
	}

	fmt.Printf("Adding worktree for branch %s at %s\n", branch, path)
	cmd := exec.Command("git", "worktree", "add", path, branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

/*
Example usage:

    # List worktrees
    gwt list

    # Switch to the main worktree
    cd $(gwt switch main)

    # Add a worktree for a feature branch
    gwt feature /tmp/feature
*/
