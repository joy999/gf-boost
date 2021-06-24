package gbbase

import (
	"github.com/gogf/gf/frame/g"
	"github.com/joy999/gf-boost/frame/gb"
)

// 关闭当前服务管理对象
func (this *ServiceManage) Close() {
	for k, v := range this.instances {
		if fn, ok := v.(IServiceDestory); ok {
			fn.Destory()
		}
		gb.CloseServiceObjectOfFactory(k, v)
	}
}

func (this *ServiceManage) CloseService(key string) {
	if v, ok := this.instances[key]; ok {
		var i interface{} = v
		if fn, ok := i.(IServiceDestory); ok {
			fn.Destory()
		}
		gb.CloseServiceObjectOfFactory(key, v)
	}
}

func (this *ServiceManage) GetService(key string) (s IService) {
	s, ok := this.instances[key]
	if !ok {
		if v := gb.NewServiceObjectOfFactory(key); s != nil {
			s = v.(IService)
		} else {
			g.Throw("未设置服务对象于工厂中")
		}
	}

	s.SetServiceManage(this)
	s.SetCtx(this.Ctx())

	return

}
