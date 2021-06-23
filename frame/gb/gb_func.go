package gb

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/joy999/gf-boost/net/gbhttp"
)

func NewServer(s *ghttp.Server) *gbhttp.Server {
	o := new(gbhttp.Server)
	o.Server = s
	return o
}
