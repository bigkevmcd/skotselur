package profiles

import (
	"context"

	"github.com/go-logr/logr"
)

type ProfileRef struct {
	RepoURL string
	Branch  string // TODO: This should be a tag
}

const (
	DEBUG = 1
)

// Install installs a profile by duplicating all the files into a path.
func Install(ctx context.Context, l logr.Logger, outputPath string, ref ProfileRef, cloner RepoCloner) error {
	l.V(DEBUG).Info("installing profile", "ref", ref, "output", outputPath)
	if err := cloner.Clone(ctx, ref, outputPath); err != nil {
		return err
	}
	return nil
}
