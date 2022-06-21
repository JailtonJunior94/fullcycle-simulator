package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChan chan *ckafka.Message
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupID"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	if err := c.SubscribeTopics(topics, nil); err != nil {
		log.Fatalf("error subscribe kafka message" + err.Error())
	}
	fmt.Println("kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MessageChan <- msg
		}
	}
}
