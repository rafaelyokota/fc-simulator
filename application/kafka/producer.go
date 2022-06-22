package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rafaelyokota/imersaofc-simulator/application/route"
	kafka2 "github.com/rafaelyokota/imersaofc-simulator/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka2.NewKafkaProducer()
	router2 := route.NewRoute()
	json.Unmarshal(msg.Value, &router2)
	router2.LoadPositions()
	positions, err := router2.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka2.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 1000)
	}
}
