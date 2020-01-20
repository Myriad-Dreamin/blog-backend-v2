package main

import (
	"flag"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-template/server"
	_ "net/http/pprof"
	"os"
	"strings"
)

const (
	defaultPort   = ":23336"
	defaultConfig = "./config"
)

var (
	_config  = flag.String("config", "", "config path")
	_port    = flag.String("port", "", "serve on port")
	_isDebug = flag.Bool("debug", false, "serve with debug mode")
	https    = flag.Bool("https", false, "serve with https")
	_crtFile = flag.String("crt", "", "crt file path")
	_keyFile = flag.String("key", "", "private key file path")
)

func config() string {
	return fstring("config", "BLOG_CONFIG", defaultConfig, _config)
}

func port() string {
	return fstring("port", "BLOG_PORT", defaultPort, _port)
}

func crtFile() string {
	return fstring("crt", "BLOG_CRT_FILE", "myd.crt", _crtFile)
}

func keyFile() string {
	return fstring("key", "BLOG_KEY_FILE", "myd.pri", _keyFile)
}

func main() {
	srv := server.New(config())
	if srv == nil {
		return
	}

	// srv.Inject(myPlugins...)

	if isDebug() {
		srv = srv.WithPProf(port())
	}
	if isHttps() {
		fmt.Println("crtFile", crtFile())
		fmt.Println("keyFile", keyFile())
		srv.ServeTLS(port(), crtFile(), keyFile())
	} else {
		srv.Serve(port())
	}

}

func fstring(cname, eName, defaultValue string, cValue *string) string {
	if len(*cValue) == 0 {
		p := os.Getenv(eName)
		if len(p) == 0 {
			return defaultValue
		}
		return p
	}
	return *cValue
}

func isDebug() bool {
	if *_isDebug {
		return true
	}
	x := os.Getenv("BLOG_DEBUG")
	if strings.ToLower(x) == "true" {
		return true
	}
	return false
}

func isHttps() bool {
	if *https {
		return true
	}
	x := os.Getenv("BLOG_HTTPS")
	if strings.ToLower(x) == "true" {
		return true
	}
	return false
}
