package gb

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
