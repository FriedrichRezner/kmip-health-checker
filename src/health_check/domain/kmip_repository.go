package domain

import "context"

type (
	KMIPRepository interface {
		Create(ctx context.Context) (string, error)
		Destroy(ctx context.Context, id string) error
	}
)
