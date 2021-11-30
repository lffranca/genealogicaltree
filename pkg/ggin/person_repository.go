package ggin

import (
	"context"
	"github.com/lffranca/genealogicaltree/internal"
)

type PersonRepository interface {
	All(ctx context.Context) ([]*internal.Person, error)
	ByName(ctx context.Context, name *string) ([]*internal.Person, error)
	ShortestPath(ctx context.Context, name1, name2 *string) ([]*internal.Person, error)
}
