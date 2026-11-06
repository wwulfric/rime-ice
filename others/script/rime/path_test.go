package rime

import (
	"path/filepath"
	"testing"
)

func TestRimeDirOverride(t *testing.T) {
	directory := t.TempDir()
	t.Setenv("RIME_DIR", directory+string(filepath.Separator))
	if got := getRimeDir(); got != directory {
		t.Fatalf("getRimeDir() = %q, want %q", got, directory)
	}
}
