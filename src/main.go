package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/astraeus-lab/astraeus-cmdb/src/common"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/constant"
	"github.com/astraeus-lab/astraeus-cmdb/src/web"
)

func main() {
	config, err := common.InitCommonDepend()
	if err != nil {
		panic("init denpend err: " + err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	go web.StartWeb(ctx, &config.Config.Web)

	// TODO: start API Server
	// go apiserver.StartAPIServer(ctx, &config.Config.APIServer)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
		cancel()
		time.Sleep(constant.GraceShutdownTimeout * time.Second)
		if errMsg := common.CloseCommonDepend(); errMsg != "" {
			panic(errMsg)
		}
		os.Exit(0)
	}
}
