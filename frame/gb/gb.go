package gb

import (
	"github.com/joy999/gf-boost/frame/gbbase"
	"github.com/joy999/gf-boost/frame/gbcontext"
	"github.com/joy999/gf-boost/frame/gbdb"
)

type (
	BaseService struct {
		Service gbbase.ServiceManage
		Dao     gbbase.DaoManage
	}
	BaseApi struct {
		BaseService
	}
	BaseDao struct {
		gbcontext.ContextManage
	}
	ServiceManage = gbbase.ServiceManage
	// Service           = gbbase.Service
	Context           = gbcontext.Context
	TransactionManage = gbdb.TransactionManage
)
