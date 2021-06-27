package gbcontext

import (
	"context"

	"github.com/gogf/gf/container/gmap"
)

func (this *Context) SetCtx(ctx context.Context) {
	this.Context = ctx
}

func (this Context) Ctx() context.Context {
	return this.Context
}

func (this *Context) getDataMap() *gmap.StrAnyMap {
	var m *gmap.StrAnyMap
	m, ok := this.Context.Value(ContextVarKey).(*gmap.StrAnyMap)
	if !ok {
		m = gmap.NewStrAnyMap(true)
		this.Context = context.WithValue(this.Context, ContextVarKey, m)
	}
	return m
}

func (this *Context) SetData(key string, val interface{}) {
	this.getDataMap().Set(key, val)
}

func (this *Context) GetData(key string) interface{} {
	m := this.getDataMap()
	return m.Get(key)
}

func (this *Context) RemoveData(key string) {
	this.getDataMap().Remove(key)
}

func (this *Context) Contains(key string) bool {
	return this.getDataMap().Contains(key)
}
