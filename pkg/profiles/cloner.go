package profiles

import "context"

type RepoCloner interface {
	Clone(ctx context.Context, ref ProfileRef, dst string) error
}
