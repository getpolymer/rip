package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}

func handler(res http.ResponseWriter, req *http.Request) {
	proxyMapping := map[string]string{
		"bin-1.nbrn.xyz": "http://localhost:9001",
		"bin-2.nbrn.xyz": "http://localhost:9002",
		"bin-3.nbrn.xyz": "http://localhost:9003",
	}

	fmt.Printf("%+v \n", req.Header)

	host := req.Header.Get("Host")
	log.Println(host)
	if host == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	backend, ok := proxyMapping[host]
	if !ok {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serveReverseProxy(backend, res, req)
}

// StartReverseProxy starts the reverse proxy
func StartReverseProxy() {
	log.Println("Starting Reverse Proxy on port 8080...")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
