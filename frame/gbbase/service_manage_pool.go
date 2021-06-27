package gbbase

import (
	"time"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gpool"
)

func init() {
	serviceManagePool = gpool.New(time.Minute*5, func() (interface{}, error) {
		o := new(ServiceManage)
		// o.instances = make(map[string]IService, 0)
		o.instances = gmap.NewStrAnyMap(true)
		return o, nil
	}, func(i interface{}) {
		s := i.(*ServiceManage)
		s.Close()
	})
}

func GetServiceManageObject() *ServiceManage {
	i, _ := serviceManagePool.Get()
	return i.(*ServiceManage)
}

func CloseServiceManageObject(s *ServiceManage) {
	s.Close()
}
