package gbhttp

import (
	"time"

	"github.com/gogf/gf/net/ghttp"
)

type (
	Server struct {
		*ghttp.Server
	}

	HandlerFunc = ghttp.HandlerFunc
	Request     = ghttp.Request

	// errorStack is the interface for Stack feature.
	errorStack interface {
		Error() string
		Stack() string
	}
)

const (
	HOOK_BEFORE_SERVE     = "HOOK_BEFORE_SERVE"  // Deprecated, use HookBeforeServe instead.
	HOOK_AFTER_SERVE      = "HOOK_AFTER_SERVE"   // Deprecated, use HookAfterServe instead.
	HOOK_BEFORE_OUTPUT    = "HOOK_BEFORE_OUTPUT" // Deprecated, use HookBeforeOutput instead.
	HOOK_AFTER_OUTPUT     = "HOOK_AFTER_OUTPUT"  // Deprecated, use HookAfterOutput instead.
	HookBeforeServe       = "HOOK_BEFORE_SERVE"
	HookAfterServe        = "HOOK_AFTER_SERVE"
	HookBeforeOutput      = "HOOK_BEFORE_OUTPUT"
	HookAfterOutput       = "HOOK_AFTER_OUTPUT"
	ServerStatusStopped   = 0
	ServerStatusRunning   = 1
	supportedHttpMethods  = "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE"
	defaultServerName     = "default"
	defaultDomainName     = "default"
	defaultMethod         = "ALL"
	handlerTypeHandler    = 1
	handlerTypeObject     = 2
	handlerTypeController = 3
	handlerTypeMiddleware = 4
	handlerTypeHook       = 5
	exceptionExit         = "exit"
	exceptionExitAll      = "exit_all"
	exceptionExitHook     = "exit_hook"
	routeCacheDuration    = time.Hour
)

var (
	methodsMap = make(map[string]struct{})
)
