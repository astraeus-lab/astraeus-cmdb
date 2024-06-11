package main

import (
	"os"
	"os/signal"

	"github.com/astraeus-lab/astraeus-cmdb/src/common"
	"github.com/astraeus-lab/astraeus-cmdb/src/web"
)

func main() {
	config, err := common.InitCommonDepend()
	if err != nil {
		panic("init denpend err: " + err.Error())
	}

	web.StartWeb(&config.Config.Web)

	// TODO: start API Server
	// go apiserver.StartAPIServer(&config.Config.APIServer)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
		if errMsg := common.CloseCommonDepend(); errMsg != "" {
			panic(errMsg)
		}
		os.Exit(0)
	}
}
