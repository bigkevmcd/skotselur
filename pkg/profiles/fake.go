package profiles

import (
	"context"
	"errors"
	"fmt"
	"strings"

	continuityfs "github.com/containerd/continuity/fs"
)

var _ RepoCloner = (*fakeCloner)(nil)

func withRepo(p ProfileRef, src string) func(*fakeCloner) {
	return func(f *fakeCloner) {
		f.repos[testProfileKey(p)] = src
	}
}

func newFakeCloner(opts ...func(*fakeCloner)) *fakeCloner {
	fc := &fakeCloner{
		repos: make(map[string]string),
	}
	for _, o := range opts {
		o(fc)
	}
	return fc
}

type fakeCloner struct {
	repos map[string]string
}

func (f *fakeCloner) Clone(ctx context.Context, ref ProfileRef, dst string) error {
	path, ok := f.repos[testProfileKey(ref)]
	if !ok {
		return errors.New("could not find profile")
	}
	if err := continuityfs.CopyDir(dst, path); err != nil {
		return fmt.Errorf("failed to copy directory from %s to %s", path, dst)
	}
	return nil
}

func testProfileKey(r ProfileRef) string {
	return strings.Join([]string{r.RepoURL, r.Branch}, ":")
}
