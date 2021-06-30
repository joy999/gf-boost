package gbcontext

import (
	"golang.org/x/net/context"
)

func (this *ContextManage) SetCtx(ctx context.Context) {
	if c, ok := ctx.(*Context); ok {
		ctx := c.Ctx()
		if ctx == nil {
			ctx = context.TODO()
		}
		this.Context = c
	} else {
		if this.Context == nil {
			this.Context = NewContext(ctx)
		} else {
			this.Context.SetCtx(ctx)
		}
	}
}

func (this *ContextManage) Ctx() context.Context {
	// if this.Context != nil {
	return this.Context.Ctx()
	// } else {
	// 	return nil
	// }
}

func (this *ContextManage) SetCtxVar(key string, val interface{}) {
	this.Context.SetData(key, val)
}

func (this *ContextManage) GetCtxVar(key string) interface{} {
	return this.Context.GetData(key)
}

func (this *ContextManage) Close() {
	this.Context = nil
}
