package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/chayaninpia/go-backend-microservice/pdst001/handler"
	"github.com/chayaninpia/go-backend-microservice/pdst001/util"
	"github.com/spf13/viper"
)

func main() {

	util.InitConfig()

	serviceId := viper.GetString("serviceId")
	log.Printf("Start Messenger [ %v ]", strings.ToUpper(serviceId))

	h := handler.NewHandler()

	//Need to implement service
	go handler.Messenger(h.Pdst001)

	signals := make(chan (os.Signal), 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	signal := <-signals
	log.Printf("Received signal %s and exit", signal.String())
}
