package handler

import (
	"encoding/json"
	"log"

	"github.com/chayaninpia/go-backend-microservice/models"
	"github.com/chayaninpia/go-backend-microservice/pdst001/util"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

//Need to implement models on type service
type Service func(req *models.Pdst001I) (*models.Pdst001O, error)

func Messenger(service Service) {

	topic := viper.GetString("serviceId")
	ch := make(chan *kafka.Message, 1)
	go util.Consumer(topic, 0, ch)

	for {

		m := <-ch
		if m == nil {
			continue
		}

		//Need to implement models on request models
		req := models.Pdst001I{}

		if err := json.Unmarshal(m.Value, &req); err != nil {
			log.Println(err.Error())
		}

		res, err := service(&req)

		msg := make([]byte, 0)
		if err != nil {
			msg, err = json.Marshal(err)
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("error : %++v", err.Error())
		} else {
			msg, err = json.Marshal(res)
			if err != nil {
				log.Println(err.Error())
			}
			log.Printf("response : %++v", res)
		}

		log.Println(string(msg))
		if err := util.Producer(msg, topic, 1); err != nil {
			log.Println(err.Error())
		}
	}
}
