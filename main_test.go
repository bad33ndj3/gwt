package main

import (
	"os"
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

	path, err := findWorktree("feature")
	if err != nil {
		t.Fatalf("findWorktree returned error: %v", err)
	}
	if path != "/tmp/feature" {
		t.Fatalf("expected /tmp/feature, got %s", path)
	}
}

func TestListWorktrees(t *testing.T) {
	if err := os.Setenv("GWT_WORKTREE_LIST", sampleWorktreeOutput); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}
	defer func() {
		if err := os.Unsetenv("GWT_WORKTREE_LIST"); err != nil {
			t.Fatalf("failed to unset env: %v", err)
		}
	}()

	tmp, err := os.CreateTemp("", "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Remove(tmp.Name()); err != nil {
			t.Fatalf("failed to remove temp file: %v", err)
		}
	}()
	old := os.Stdout
	os.Stdout = tmp
	defer func() { os.Stdout = old }()

	if err := listWorktrees(); err != nil {
		t.Fatalf("listWorktrees returned error: %v", err)
	}
	if err := os.Stdout.Close(); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "feature") {
		t.Fatalf("expected output to contain feature, got %q", string(data))
	}
}
