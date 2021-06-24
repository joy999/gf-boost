package gbbase

import "context"

type (
	ICtx interface {
		SetCtx(context.Context)
		Ctx() context.Context
	}

	IServiceDestory interface {
		Destory()
	}

	IService interface {
		ICtx
		IServiceDestory
		SetServiceManage(*ServiceManage)
	}
)

type (
	ServiceCtx struct {
		ctx context.Context
	}

	Service struct {
		ServiceCtx
		*ServiceManage
	}

	//使用ServiceManage时，应该继承自该类并重新实现
	ServiceManage struct {
		ServiceCtx
		instances map[string]IService
	}
)
