package file

import (
	"admin-go/core/sdk/config/private/source"
	"context"
)

type filePathKey struct{}

// WithPath sets the path to file
func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, filePathKey{}, p)
	}
}
