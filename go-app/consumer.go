package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka_kafka_1:9092",
		"topic":             "my-topic-three"}

	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("Error consumer", err.Error())
	}

	topics := []string{"teste"}

	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
