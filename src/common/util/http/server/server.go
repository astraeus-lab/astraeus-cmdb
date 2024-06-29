package server

import (
	"context"
	"errors"
	"net/http"
	"time"
)

const (
	graceShutdownTimeout = 30
)

// StartSrvWithGracefulShutdown start the incoming http.Server with
// goroutine and listen for shutdown signals to achieve elegant exit.
//
// It will block the caller to enable listening
// and processing of the shutdown signal.
func StartSrvWithGracefulShutdown(ctx context.Context, srv *http.Server, graceShutdownTime int) (err error) {
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return
		}
	}()

	shutdownCtx, cancel := context.WithTimeout(ctx, time.Duration(graceShutdownTime)*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		if err = srv.Shutdown(shutdownCtx); err != nil {
			return
		}
	}

	return
}
