package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

const (
	topic          = "my-topic-three"
	broker1Address = "localhost:9091"
)

type Quote struct {
	Id   string  `json:"uuid"`
	BRL  float32 `json:"brl,float32"`
	EUR  float32 `json:"eur,float32"`
	JPY  float32 `json:"jpy,float32"`
	Date string  `json:"date,float32"`
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

// var quotesUrl = os.Getenv("QUOTES_URL")

func producer(topic string, value string) {
	fmt.Println("Consumindo tópico: [", topic, "]")
	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	for {
		err := w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(value),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}

func getQuotes(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("http://localhost:5000/quotes")
	if err != nil {
		fmt.Println("Erro de HTTP")
	}
	data, _ := ioutil.ReadAll(response.Body)

	var quote Quote

	json.Unmarshal(data, &quote)
	// fmt.Println(string(data))

	producer(topic, string(data))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get/quotes", getQuotes)
	http.ListenAndServe(":8081", r)
}
