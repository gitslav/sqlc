// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package querytest

import (
	"context"
)

type Querier interface {
	Bar(ctx context.Context) error
	Bars(ctx context.Context) error
}

var _ Querier = (*Queries)(nil)