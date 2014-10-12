package main

import (
	"html/template"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var templates = template.Must(template.ParseFiles("public/error.html"))

type Proxy struct {
}

func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check redis first
	id := r.Header.Get("DockerID")
	port := r.Header.Get("Port")

	container, err := inspectCachedContainer(id)
	if err != nil || container == nil {
		renderError(w, 503)
		return
	}

	scheme := "http://"
	if r.URL.Scheme != "" {
		scheme = r.URL.Scheme
	}

	backend, err := url.Parse(scheme + container.NetworkSettings.IPAddress + ":" + port)
	if err != nil {
		renderError(w, 503)
		return
	}

	// Reverse proxy the request.
	// (Need special code for websockets, courtesy of bradfitz)
	if r.Header.Get("Upgrade") == "websocket" {
		proxyWebsocket(w, r, backend.Host)
	} else {
		p := httputil.NewSingleHostReverseProxy(backend)
		p.ServeHTTP(w, r)
	}
}

// proxyWebsocket copies data between websocket client and server until one side
// closes the connection.  (ReverseProxy doesn't work with websocket requests.)
func proxyWebsocket(w http.ResponseWriter, r *http.Request, host string) {
	d, err := net.Dial("tcp", host)
	if err != nil {
		http.Error(w, "Error contacting backend server.", 500)
		return
	}
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Not a hijacker?", 500)
		return
	}
	nc, _, err := hj.Hijack()
	if err != nil {
		return
	}
	defer nc.Close()
	defer d.Close()

	err = r.Write(d)
	if err != nil {
		return
	}

	errc := make(chan error, 2)
	cp := func(dst io.Writer, src io.Reader) {
		_, err := io.Copy(dst, src)
		errc <- err
	}
	go cp(d, nc)
	go cp(nc, d)
	<-errc
}

func removePort(s string) string {
	return s[:strings.LastIndex(s, ":")]
}

func getHost(s string) string {
	return s[:strings.Index(s, ".")]
}

func renderError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	templates.Execute(w, nil)
}
