package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reverse_proxy/config"
	"reverse_proxy/internal"
)

func ServeApp(address string) {
	http.HandleFunc("/", ServeHTTP)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error while listening error: %v", err))
	}
}

func ServeHTTP(response http.ResponseWriter, request *http.Request) {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	matchedRoute, err := internal.Matcher(request, cfg)

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		_, _ = response.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	remote, err := url.Parse(fmt.Sprintf("%s:%d", matchedRoute.Address, matchedRoute.Port))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))

	response.Header().Set("X-Proxy", "SimpleHTTP Proxy Server")

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(response, request)
}
