package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "my-topic-three"
	broker1Address = "localhost:9091"
)

func consumer(topic string) {
	fmt.Println("Consumindo tópico: [", topic, "]")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
	r.Close()
}

func main() {

	consumer(topic)

}
