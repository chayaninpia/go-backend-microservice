package handler

import (
	"encoding/json"
	"log"

	"github.com/chayaninpia/go-backend-microservice/models"
	"github.com/chayaninpia/go-backend-microservice/pdst001/util"
	"github.com/spf13/viper"
)

func Messenger() {

	h := Pdst001Handler{}

	topic := viper.GetString("serviceId")
	ch := make(chan []byte)
	go util.Consumer(topic, 0, ch)

	for {
		req := models.Pdst001I{}
		msg, recived := <-ch
		if recived {
			if err := json.Unmarshal(msg, &req); err != nil {
				log.Fatalln(err.Error())
			}

			res, err := h.Pdst001(&req)

			msg := make([]byte, 0)
			if err != nil {
				msg, err = json.Marshal(err)
				if err != nil {
					log.Fatalln(err.Error())
				}
			} else {
				msg, err = json.Marshal(res)
				if err != nil {
					log.Fatalln(err.Error())
				}
			}

			if err := util.Producer(msg, topic, 0); err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}
