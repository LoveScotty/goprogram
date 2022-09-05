package bookstore

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/LoveScotty/goprogram/internal/store"
	"github.com/LoveScotty/goprogram/server/middleware"
)

type Server struct {
	s   store.Factory
	srv *http.Server
}

func NewServer(addr string, s store.Factory) *Server {
	srv := &Server{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := srv.Router()
	_ = router.Run(addr)
	srv.srv.Handler = middleware.Logging(middleware.Validating(router))

	return srv
}

func (srv *Server) Router() *gin.Engine {
	router := gin.New()
	bookStoreGroup := router.Group("/bookstore")

	bookGroup := bookStoreGroup.Group("/book")
	srv.bookRouter(bookGroup)

	return router
}

func (srv *Server) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)

	go func() {
		err = srv.srv.ListenAndServe()
		errChan <- err
	}()
	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.srv.Shutdown(ctx)
}
