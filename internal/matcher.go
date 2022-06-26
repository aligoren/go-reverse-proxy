package internal

import (
	"errors"
	"net/http"
	"reverse_proxy/config"
	"strings"
)

func Matcher(request *http.Request, proxyConfig *config.ProxyConfig) (*config.Route, error) {
	var route *config.Route

	headers := request.Header

	for name, _ := range headers {
		for _, proxyRoute := range proxyConfig.Routes {
			// TODO: Make this section clean. It's not reusable code. So, it should be refactored later!
			routeType := strings.ToLower(proxyRoute.Type)
			if routeType == "header" {
				for key, routeHeader := range proxyRoute.Headers {
					if strings.ToLower(key) == strings.ToLower(name) {
						for _, value := range routeHeader.Values {
							if strings.ToLower(headers.Get(key)) == strings.ToLower(value) {
								route = &proxyRoute
								break
							}
						}
					}
				}
			} else if routeType == "path" {
				for _, routePath := range proxyRoute.Paths {
					if routePath == strings.ToLower(request.URL.Path) {
						route = &proxyRoute
						break
					}
				}
			}
		}
	}

	if route == nil {
		return route, errors.New("there is no route found")
	}

	return route, nil
}
