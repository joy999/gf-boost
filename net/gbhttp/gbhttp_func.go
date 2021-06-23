// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gbhttp

import (
	_ "unsafe"

	"github.com/gogf/gf/net/ghttp"
)

//go:linkname niceCallFunc github.com/gogf/gf/net/ghttp.niceCallFunc
func niceCallFunc(func())

//go:linkname mergeBuildInNameToPattern github.com/gogf/gf/net/ghttp.(*Server).mergeBuildInNameToPattern
func mergeBuildInNameToPattern(*ghttp.Server, string, string, string, bool) string

//go:linkname parsePattern github.com/gogf/gf/net/ghttp.(*Server).parsePattern
func parsePattern(*ghttp.Server, string) (domain, method, path string, err error)

//go:linkname serveHandlerKey github.com/gogf/gf/net/ghttp.(*Server).serveHandlerKey
func serveHandlerKey(*ghttp.Server, string, string, string) string

//go:linkname bindHandlerByMap github.com/gogf/gf/net/ghttp.(*Server).bindHandlerByMap
func bindHandlerByMap(*ghttp.Server, map[string]*handlerItem)
