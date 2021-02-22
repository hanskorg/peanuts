package http

import (
    "context"
    "fmt"
    "github.com/hanskorg/logkit"
    "io"
    "net/http"
	"peanuts/tools"
)

var (
	s *http.Server
)

type Server struct {
	Listen  string
	Server  *http.Server
	service interface{}
	option  *Options
}
type Options struct {
	Debug        bool
	ServerName   string
	ErrLogger    io.Writer
	AccessLogger io.Writer
}

func DefaultOptions() *Options {
	return &Options{
		Debug:        false,
		ServerName:   "tools",
		ErrLogger:    logkit.NewLogWriter(logkit.LevelError),
		AccessLogger: logkit.NewLogWriter(logkit.LevelInfo),
	}
}
func New(listen string, options ...*Options) (s *Server) {
	var (
		option *Options
	)
	if len(options) == 0 {
		option = DefaultOptions()
	} else {
		option = options[0]
	}
	s = &Server{
		Listen: listen,
		option: option,
	}

	go func() {
		s.Server = &http.Server{
			Addr:    s.Listen,
			Handler: nil,
		}

        http.HandleFunc("/qrcode", tools.Qrcode)


        if err := s.Server.ListenAndServe(); err != nil {
			logkit.Infof("http server fail, %s\n", err.Error())
		}

	}()
	logkit.Infof("http server started, %s\n", s.Listen)
	return s
}

func (s *Server) Shutdown() error {
	if s.Server == nil {
		return fmt.Errorf("not started")
	}
	return s.Server.Shutdown(context.Background())
}
