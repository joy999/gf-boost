package gbhttp

import "strings"

func init() {
	// Initialize the methods map.
	for _, v := range strings.Split(supportedHttpMethods, ",") {
		methodsMap[v] = struct{}{}
	}
}
