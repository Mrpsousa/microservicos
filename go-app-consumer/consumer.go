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

type Quote struct {
	Uuid string  `json:"uuid"`
	BRL  float32 `json:"price,string"`
	EUR  float32 `json:"price,string"`
	JPY  float32 `json:"price,string"`
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

func consumer(topic string) []string {
	fmt.Println("Consumindo tópico: [", topic, "]")
	var list_ []string

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println(msg.Value)
		list_ = append(list_, string(msg.Value))

	}
	r.Close()
	return list_
}

func loadData() {
	jsonFile := consumer(topic)
	// if err != nil {
	// 	fmt.Println("erro: ", err.Error())
	// }

	fmt.Println(jsonFile)
	// defer jsonFile.Close()
	fmt.Println(jsonFile)
	// data, err := ioutil.ReadAll(jsonFile)
	// return jsonFile
}

// func ListProducts(w http.ResponseWriter, r *http.Request) {
// 	// products := loadData()
// 	w.Write([]byte(products))
// }

// func GetProductById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	// data := loadData()

// 	var quotes Quotes
// 	json.Unmarshal(data, &quotes)

// 	for _, v := range quotes.Quotes {
// 		if v.Uuid == vars["id"] {
// 			product, _ := json.Marshal(v)
// 			w.Write([]byte(product))
// 		}
// 	}
// }

func main() {
	loadData()
	// 	r := mux.NewRouter()
	// 	r.HandleFunc("/products", ListProducts)
	// 	r.HandleFunc("/product/{id}", GetProductById)
	// 	http.ListenAndServe(":8081", r)
}