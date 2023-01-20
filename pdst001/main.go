package main

import (
	"github.com/chayaninpia/go-backend-microservice/pdst001/handler"
	"github.com/chayaninpia/go-backend-microservice/pdst001/util"
)

func main() {

	util.InitConfig()

	handler.Messenger()

}
