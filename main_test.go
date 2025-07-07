package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

const sampleWorktreeOutput = `worktree /tmp/main
HEAD 1234567
branch refs/heads/main

worktree /tmp/feature
HEAD 89abcde
branch refs/heads/feature
`

func TestFindWorktree(t *testing.T) {
	if err := os.Setenv("GWT_WORKTREE_LIST", sampleWorktreeOutput); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("GWT_WORKTREE_LIST"); err != nil {
			t.Fatalf("failed to unset env: %v", err)
		}
	}()

	cmd := exec.Command("bash", "./gwt.sh", "switch", "feature")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	path := strings.TrimSpace(string(out))
	if path != "/tmp/feature" {
		t.Fatalf("expected /tmp/feature, got %s", path)
	}
}

func TestListWorktrees(t *testing.T) {
	cmd := exec.Command("bash", "./gwt.sh", "list")
	cmd.Env = append(os.Environ(), "GWT_WORKTREE_LIST="+sampleWorktreeOutput)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if !strings.Contains(string(out), "feature") {
		t.Fatalf("expected output to contain feature, got %q", string(out))
	}
}
