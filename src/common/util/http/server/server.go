package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartSrvWithGracefulShutdown start the incoming http.Server with
// goroutine and listen for shutdown signals to achieve elegant exit.
//
// It will block the caller to enable listening
// and processing of the shutdown signal.
func StartSrvWithGracefulShutdown(srv *http.Server) (err error) {
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		return
	}

	return
}
