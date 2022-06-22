package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{MsgChan: msgChan}
}

func (k *KafkaConsumer) Consumer() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	fmt.Println(os.Getenv("KafkaBootstrapServers"), os.Getenv("KafkaConsumerGroupId"))
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("Error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka Consumer has been started topic:", topics)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
