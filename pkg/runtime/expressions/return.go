package expressions

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"github.com/pkg/errors"
)

type (
	ReturnExpression struct {
		src       core.SourceMap
		predicate core.Expression
	}
)

func NewReturnExpression(
	src core.SourceMap,
	predicate core.Expression,
) (*ReturnExpression, error) {
	if predicate == nil {
		return nil, errors.Wrap(core.ErrMissedArgument, "expression")
	}

	return &ReturnExpression{
		src,
		predicate,
	}, nil
}

func (e *ReturnExpression) Exec(ctx context.Context, scope *core.Scope) (core.Value, error) {
	val, err := e.predicate.Exec(ctx, scope)

	if err != nil {
		return values.None, core.SourceError(e.src, err)
	}

	return val, nil
}
