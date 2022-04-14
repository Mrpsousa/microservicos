package main

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consume(ctx context.Context) {
	const (
		topic          = "teste"
		broker1Address = "localhost:9092"
	)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}

func main() {

	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	consume(ctx)
}

// configMap := &kafka.ConfigMap{
// 	"bootstrap.servers": "kafka_kafka_1:9092",
// 	"client.id":         "gokafka-consumer",
// 	"group.id":          "gokafka-group",
// 	"auto.offset.reset": "earliest",
// }
// c, err := kafka.NewConsumer(configMap)

// if err != nil {
// 	fmt.Println("Error consumer", err.Error())
// }

// topics := []string{"teste"}

// c.SubscribeTopics(topics, nil)
// for {
// 	msg, err := c.ReadMessage(-1)
// 	if err == nil {
// 		fmt.Println(string(msg.Value), msg.TopicPartition)
// 	}
// }
// }
