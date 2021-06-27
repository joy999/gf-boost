package gbcontext

import (
	"context"
)

type (
	Context struct {
		context.Context
	}
	// data *gmap.StrAnyMap

	ContextManage struct {
		*Context
	}
)

const (
	ContextVarKey = "gb_context_ext_data"
)
