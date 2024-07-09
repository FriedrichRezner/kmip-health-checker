package domain

import "context"

type (
	// KMIPRepository is the interface for the KMIP repository
	// The implementation should be able to create and destroy keys
	KMIPRepository interface {
		Create(ctx context.Context) (string, error)
		Destroy(ctx context.Context, id string) error
	}
)
