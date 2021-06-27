package gbcontext

import (
	"context"
)

func NewContext(ctx ...context.Context) *Context {
	o := new(Context)
	if len(ctx) > 0 {
		o.Context = ctx[0]
	} else {
		o.Context = context.TODO()
	}
	return o
}
