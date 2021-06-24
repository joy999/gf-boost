package gb

import (
	"reflect"
	"time"

	"github.com/gogf/gf/container/gpool"
)

type (
	serviceClassInfo struct {
		pool   *gpool.Pool
		object interface{}
		// v      reflect.Value
		// t      reflect.Type

	}
)

var (
	serviceClasses map[string]*serviceClassInfo = make(map[string]*serviceClassInfo)
)

//设置服务类
func SetServiceClassObjectOfFactory(key string, o interface{}) {
	ci := new(serviceClassInfo)
	ci.object = o
	v := reflect.ValueOf(o)
	t := v.Type()
	ci.pool = gpool.New(time.Minute*5, func() (interface{}, error) {
		o := reflect.New(t)
		o.Elem().Set(v)

		return o.Interface(), nil
	})

	serviceClasses[key] = ci
}

//生成一个新的服务对象
func NewServiceObjectOfFactory(key string) interface{} {
	ci, ok := serviceClasses[key]
	if !ok {
		return nil
	}

	o, _ := ci.pool.Get()

	return o
}

func CloseServiceObjectOfFactory(key string, o interface{}) {
	ci, ok := serviceClasses[key]
	if !ok {
		return
	}
	ci.pool.Put(o)
}
