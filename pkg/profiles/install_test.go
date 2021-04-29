package profiles

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-logr/zapr"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

func TestInstall(t *testing.T) {
	l := zapr.NewLogger(zap.NewNop())
	ref := ProfileRef{RepoURL: "https://github.com/weaveworks/dokuwiki-profile.git", Branch: "copy-on-install"}
	cloner := newFakeCloner(withRepo(ref, mustAbs(t, "testdata/profile")))
	outputPath := testDir(t)

	if err := Install(context.TODO(), l, outputPath, ref, cloner); err != nil {
		t.Fatal(err)
	}

	want := []string{
		"dokuwiki/helmrelease.yaml",
		"dokuwiki/helmrepository.yaml",
		"profile.yaml",
	}
	if diff := cmp.Diff(want, listTree(t, outputPath)); diff != "" {
		t.Fatalf("failed to copy files:\n%s", diff)
	}
}

func mustAbs(t *testing.T, s string) string {
	t.Helper()
	absPath, err := filepath.Abs(s)
	if err != nil {
		t.Fatalf("could not calculate absolute path for %s", s)
	}
	return absPath
}

func listTree(t *testing.T, root string) []string {
	t.Helper()
	var dirs []string
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to WalkDir %s: %w", p, err)
		}
		if d.IsDir() {
			return nil
		}
		dirs = append(dirs, strings.TrimPrefix(p, root+"/"))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return dirs
}

func testDir(t *testing.T) string {
	t.Helper()
	dir, err := ioutil.TempDir("", "kustom")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("failed to cleanup %s: %s", dir, err)
		}
	})
	return dir
}
