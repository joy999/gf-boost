package gbhttp

import (
	"reflect"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

type (
	Server struct {
		*ghttp.Server
	}

	HandlerFunc = ghttp.HandlerFunc
	Request     = ghttp.Request
	Router      = ghttp.Router

	// errorStack is the interface for Stack feature.
	errorStack interface {
		Error() string
		Stack() string
	}

	handlerItem struct {
		itemId     int                // Unique handler item id mark.
		itemName   string             // Handler name, which is automatically retrieved from runtime stack when registered.
		itemType   int                // Handler type: object/handler/controller/middleware/hook.
		itemFunc   HandlerFunc        // Handler address.
		initFunc   HandlerFunc        // Initialization function when request enters the object(only available for object register type).
		shutFunc   HandlerFunc        // Shutdown function when request leaves out the object(only available for object register type).
		middleware []HandlerFunc      // Bound middleware array.
		ctrlInfo   *handlerController // Controller information for reflect usage.
		hookName   string             // Hook type name.
		router     *Router            // Router object.
		source     string             // Source file path:line when registering.
	}

	handlerController struct {
		name    string       // Handler method name.
		reflect reflect.Type // Reflect type of the controller.
	}

	serviceObjectClassInfo struct {
		rVal    reflect.Value
		methods *gmap.StrAnyMap // map[string]func(*Request)

		initFunc func(*Request)
		shutFunc func(*Request)
	}
)

const (
	// HOOK_BEFORE_SERVE     = "HOOK_BEFORE_SERVE"  // Deprecated, use HookBeforeServe instead.
	// HOOK_AFTER_SERVE      = "HOOK_AFTER_SERVE"   // Deprecated, use HookAfterServe instead.
	// HOOK_BEFORE_OUTPUT    = "HOOK_BEFORE_OUTPUT" // Deprecated, use HookBeforeOutput instead.
	// HOOK_AFTER_OUTPUT     = "HOOK_AFTER_OUTPUT"  // Deprecated, use HookAfterOutput instead.
	// HookBeforeServe       = "HOOK_BEFORE_SERVE"
	// HookAfterServe        = "HOOK_AFTER_SERVE"
	// HookBeforeOutput      = "HOOK_BEFORE_OUTPUT"
	// HookAfterOutput       = "HOOK_AFTER_OUTPUT"
	// ServerStatusStopped   = 0
	// ServerStatusRunning   = 1
	supportedHttpMethods = "GET,PUT,POST,DELETE,PATCH,HEAD,CONNECT,OPTIONS,TRACE"
	// defaultServerName     = "default"
	// defaultDomainName     = "default"
	defaultMethod      = "ALL"
	handlerTypeHandler = 1
	// handlerTypeObject     = 2
	// handlerTypeController = 3
	// handlerTypeMiddleware = 4
	// handlerTypeHook       = 5
	// exceptionExit         = "exit"
	// exceptionExitAll      = "exit_all"
	// exceptionExitHook     = "exit_hook"
	// routeCacheDuration    = time.Hour
)

var (
	methodsMap = make(map[string]struct{})

	serviceObjectClassInfoCache *gmap.StrAnyMap // map[string]*gpool.Pool
)
