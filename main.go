package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fireworkweb/godotenv"
	kafka2 "github.com/rafaelyokota/imersaofc-simulator/application/kafka"
	"github.com/rafaelyokota/imersaofc-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consumer()

	for msg := range msgChan {
		fmt.Printf("%s \n", string(msg.Value))
		go kafka2.Produce(msg)
	}

}
