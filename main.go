package main

import (
	"fmt"
	"log"

	kafkaApplication "github.com/jailtonjunior94/fullcycle-simulator/application/kafka"
	"github.com/jailtonjunior94/fullcycle-simulator/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApplication.Produce(msg)
	}
}
