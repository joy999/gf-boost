package gbbase

import (
	"context"

	"github.com/gogf/gf/frame/g"
)

// 关闭当前服务管理对象
func (this *ServiceManage) Close() {
	if this.isClosing {
		return
	}
	this.isClosing = true
	this.instances.LockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			if fn, ok := v.(IService); ok {
				fn.SetServiceManage(nil)
				fn.Close()
			}
			// CloseServiceObjectOfFactory(k, v)
		}
	})
	this.TransactionManage.Close()

	serviceManagePool.Put(this)
	//g.Log().Debugf("Close ServiceMange %p %v", this, this)
	//gdebug.PrintStack()
	this.isClosing = false
}

func (this *ServiceManage) Destory() {
	this.isClosing = true
	this.instances.LockFunc(func(m map[string]interface{}) {
		for k, v := range m {
			if fn, ok := v.(IService); ok {
				fn.SetServiceManage(nil)
				fn.Close()
			}
			// if fn, ok := v.(IServiceCloser); ok {

			// 	fn.Close()
			// }
			CloseServiceObjectOfFactory(k, v)
		}
	})
	//this.instances.Clear()
	this.TransactionManage.Close()
	g.Log().Debugf("Destory ServiceMange %p %v", this, this)
	this.isClosing = false
}

func (this *ServiceManage) CloseService(key string) {
	if this.instances.Contains(key) {
		i := this.instances.Get(key)
		if fn, ok := i.(IServiceCloser); ok {
			fn.Close()
		}
		CloseServiceObjectOfFactory(key, i)
	}
	// if v, ok := this.instances[key]; ok {
	// 	var i interface{} = v
	// 	if fn, ok := i.(IServiceCloser); ok {
	// 		fn.Close()
	// 	}
	// 	CloseServiceObjectOfFactory(key, v)
	// }
}

func (this *ServiceManage) SetCtx(ctx context.Context) {
	this.TransactionManage.SetCtx(ctx)
	// this.instances.Iterator(func(k string, v interface{}) bool {
	// 	if v, ok := v.(IService); ok {
	// 		v.SetCtx(ctx)
	// 	}
	// 	return true
	// })
}

func (this *ServiceManage) GetService(key string) interface{} {
	var s IService

	i := this.instances.Get(key)

	if i == nil {
		if v := NewServiceObjectOfFactory(key); v != nil {
			s = v.(IService)
		} else {
			g.Throw("未设置服务对象于工厂中, key:" + key)
		}
		this.instances.Set(key, s)
	} else {
		s = i.(IService)
	}

	// s, ok := this.instances[key]
	// if !ok {
	// 	if v := NewServiceObjectOfFactory(key); v != nil {
	// 		s = v.(IService)
	// 	} else {
	// 		g.Throw("未设置服务对象于工厂中, key:" + key)
	// 	}
	// 	this.instances[key] = s
	// }

	s.SetServiceManage(this)

	if v, ok := s.(Initter); ok {
		v.Init()
	}

	ctx := this.Ctx()
	if ctx != nil {
		s.SetCtx(ctx)
	}
	return s

}
