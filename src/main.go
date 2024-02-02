package main

import (
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

}
