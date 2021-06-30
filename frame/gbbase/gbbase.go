package gbbase

import (
	"context"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gpool"
	"github.com/joy999/gf-boost/frame/gbcontext"
	"github.com/joy999/gf-boost/frame/gbdb"
)

type (
	Initter interface {
		Init()
	}

	ICtx interface {
		SetCtx(context.Context)
		Ctx() context.Context
		//Transaction(func(context.Context, *gdb.TX) error)
	}

	IServiceCloser interface {
		Close()
	}

	IService interface {
		ICtx
		IServiceCloser
		SetServiceManage(*ServiceManage)
	}
)

type (
	// ServiceCtx struct {
	// 	ctx context.Context
	// }

	Service struct {
		gbcontext.ContextManage
		*ServiceManage
	}

	//使用ServiceManage时，应该继承自该类并重新实现
	ServiceManage struct {
		gbdb.TransactionManage
		instances *gmap.StrAnyMap // map[string]IService
		isClosing bool
		//Request   *ghttp.Request
	}
)

var (
	serviceManagePool *gpool.Pool
)
