package gin

import (
	"context"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// context 实现 container 的几个封装
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// 实现mustMake的封装
func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

// 实现MakeNew的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}
