package gbdb

import (
	"context"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/joy999/gf-boost/frame/gbcontext"
)

type TransactionManage struct {
	gbcontext.ContextManage
}

func NewTransactionManage(ctx ...context.Context) *TransactionManage {
	o := new(TransactionManage)
	return o
}

func (this *TransactionManage) Transaction(f func(ctx context.Context, tx *gdb.TX) error) error {
	ctx := this.Ctx()
	defer func() {
		this.SetCtx(ctx)
	}()

	return g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		defer func() {
			if re := recover(); re != nil {
				err = gerror.New("SYSTEM ERROR: " + g.NewVar(re).String())
			}
		}()
		this.SetCtx(ctx)
		err = f(ctx, tx)
		return
	})
}

func (this *TransactionManage) SetCtx(ctx context.Context) {
	this.ContextManage.SetCtx(ctx)
}

func (this *TransactionManage) Close() {
	this.ContextManage.Close()
}
