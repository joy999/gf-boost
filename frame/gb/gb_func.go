package gb

import (
	"bytes"
	"context"
	"reflect"
	"runtime"
	"strconv"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/joy999/gf-boost/frame/gbbase"
	"github.com/joy999/gf-boost/frame/gbcontext"
	"github.com/joy999/gf-boost/net/gbhttp"
)

func Server(s ...*ghttp.Server) *gbhttp.Server {
	o := new(gbhttp.Server)
	if len(s) > 0 {
		o.Server = s[0]
	} else {
		o.Server = g.Server()
	}
	return o
}

func GetMethodsOfObject(o interface{}) (ret map[string]interface{}) {
	val := reflect.ValueOf(o)
	v := val.Elem()
	t := val.Type()

	ret = make(map[string]interface{}, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			//说明是有继承方法
			m := GetMethodsOfObject(v.Elem().Field(i).Interface())
			for k, v := range m {
				ret[k] = v
			}
			continue
		}

		if field.Type.Kind() == reflect.Func {
			ret[field.Name] = v.Elem().Field(i).Interface()
		}
	}

	return ret
}

func NewServiceManage() *ServiceManage {
	return gbbase.GetServiceManageObject()
}

func SetServiceClassObjectOfFactory(key string, o interface{}) {
	gbbase.SetServiceClassObjectOfFactory(key, o)
}

func getGoroutineID() uint64 {
	b := make([]byte, 32)
	runtime.Stack(b, false)
	//fmt.Println(string(b))
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}

func NewContext(c context.Context) *Context {
	return gbcontext.NewContext(c)
}
