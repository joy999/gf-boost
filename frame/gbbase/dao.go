package gbbase

import (
	"github.com/joy999/gf-boost/frame/gbcontext"
	"github.com/joy999/gf-boost/frame/gbdb"
)

type (
	Dao struct { //此需要在使用时被继承
		gbcontext.ContextManage
	}
	DaoManage struct {
		gbdb.TransactionManage
	}
)
