package gbcontext

import "golang.org/x/net/context"

func (this *ContextManage) SetCtx(ctx context.Context) {
	if c, ok := ctx.(*Context); ok {
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
	return this.Context.Ctx()
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
