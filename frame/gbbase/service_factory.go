package gbbase

import (
	"reflect"
	"time"

	"github.com/gogf/gf/container/gmap"
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
	serviceClasses *gmap.StrAnyMap
)

func init() {
	serviceClasses = gmap.NewStrAnyMap(true) // make(map[string]*serviceClassInfo)
}

//设置服务类
func SetServiceClassObjectOfFactory(key string, o interface{}) {
	ci := new(serviceClassInfo)
	ci.object = o

	v := reflect.ValueOf(ci.object)
	t := v.Type()

	ci.pool = gpool.New(time.Minute*5, func() (interface{}, error) {

		o := reflect.New(t)
		o.Elem().Set(v)

		ret := o.Elem().Interface()
		return ret, nil
	})

	// serviceClasses[key] = ci
	serviceClasses.Set(key, ci)
}

//生成一个新的服务对象
func NewServiceObjectOfFactory(key string) interface{} {
	ci, ok := serviceClasses.Get(key).(*serviceClassInfo)

	if !ok || ci == nil {
		return nil
	}

	o, _ := ci.pool.Get()

	return o
}

func CloseServiceObjectOfFactory(key string, o interface{}) {
	ci, ok := serviceClasses.Get(key).(*serviceClassInfo) // serviceClasses[key]
	if !ok || ci == nil {
		return
	}
	ci.pool.Put(o)
}
