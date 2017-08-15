package main

import (
	"app/common"
	"app/httpsvr"
	logger "github.com/shengkehua/xlog4go"
)

var cityInfoPath string = "./conf/CN_adm.txt"

func main() {
	err := common.Init(cityInfoPath)
	if err != nil {
		logger.Error("init citylist fail")
		logger.Error(err.Error())
	}
	err = httpsvr.Init()
	if err != nil {
		logger.Error("init http server fail")
	}
	select {}
}
