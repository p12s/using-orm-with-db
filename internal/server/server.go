package server

import (
	"context"
	"github.com/p12s/using-orm-with-db/internal/config"
	"github.com/p12s/using-orm-with-db/internal/transport/rest"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	SERVER_TIMEOUT = 10 * time.Second
)

// httpServer - http-server struct
type httpServer struct {
	server http.Server
	port   int
}

// NewHttpServer - constructor
func NewHttpServer(cfg config.Backend, handler rest.Handler) *httpServer {
	return &httpServer{
		port: cfg.Port,
		server: http.Server{
			Addr:         ":" + strconv.Itoa(cfg.Port),
			Handler:      handler.InitRouter(),
			ReadTimeout:  SERVER_TIMEOUT,
			WriteTimeout: SERVER_TIMEOUT,
		},
	}
}

// Run - listen and serve
func (s *httpServer) Run(log *zap.SugaredLogger) {
	go func() {
		log.Infof("http server restarted with port: %d", s.port)

		err := s.server.ListenAndServe()
		if err != nil {
			log.Infof(err.Error())
		}
	}()
}

// WaitShutdown - gracefully shutdown
func (s *httpServer) WaitShutdown() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-done
	
	return s.server.Shutdown(context.Background())
}
