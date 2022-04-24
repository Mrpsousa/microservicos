package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Quote struct {
	BRL float32 `json:"price,string"`
	EUR float32 `json:"price,string"`
	JPY float32 `json:"price,string"`
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

var quotesUrl string

func init() {
	quotesUrl = os.Getenv("QUOTES_URL")
}

func loadQuotes() []Quote {
	response, err := http.Get(quotesUrl + "/quotes")
	if err != nil {
		fmt.Println("Erro de HTTP")
	}
	data, _ := ioutil.ReadAll(response.Body)

	var quotes Quotes
	json.Unmarshal(data, &quotes)

	return quotes.Quotes
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ListQuotes)
	r.HandleFunc("/product/{id}", ShowQuote)
	http.ListenAndServe(":8081", r)
}

func ListQuotes(w http.ResponseWriter, r *http.Request) {
	quotes := loadQuotes()
	t := template.Must(template.ParseFiles("templates/catalog.html"))
	t.Execute(w, quotes)
}

func ShowQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(quotesUrl + "/product/" + vars["id"])
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	var quote Quote
	json.Unmarshal(data, &quote)

	t := template.Must(template.ParseFiles("templates/view.html"))
	t.Execute(w, quote)
}
