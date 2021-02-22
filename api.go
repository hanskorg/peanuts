package main

import (
	"flag"
	"github.com/hanskorg/logkit"
	"os"
	"os/signal"
	"qrcode.icool.io/http"
	"syscall"
)

var (
	hs *http.Server
)

func main() {
	flag.Parse()
	hs = http.New("0.0.0.0:9909")
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			hs.Shutdown()
			logkit.Infof("tools-server [version: %s] exit")
			logkit.Exit()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
