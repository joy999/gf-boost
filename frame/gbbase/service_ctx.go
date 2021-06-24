package gbbase

import "context"

func (this *ServiceCtx) SetCtx(ctx context.Context) {
	this.ctx = ctx
}

func (this ServiceCtx) Ctx() context.Context {
	return this.ctx
}
