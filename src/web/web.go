package web

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/config"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/constant"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util/http/server"
)

func StartWeb(ctx context.Context, c *config.Web) {
	srv := &http.Server{
		Addr:                         fmt.Sprintf(":%s", strconv.Itoa(c.Port)),
		Handler:                      NewEngine(),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		ConnState:                    nil,
		ErrorLog:                     log.New(io.Discard, "", 0),
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	if err := server.StartSrvWithGracefulShutdown(ctx, srv, constant.GraceShutdownTimeout); err != nil {
		return
	}
}
