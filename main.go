package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	// register giveQuote Handler for "/quote" endpoint
	http.HandleFunc("/quote", randomQuote)

	// start server on port 8080
	port := ":" + os.Getenv("PORT")
	http.ListenAndServe(port, nil)
}

// randomQuote parses the quotes.json and returns a random quote
func randomQuote(w http.ResponseWriter, r *http.Request) {
	// instantiate QuotesJson struct
	quotes := make([]QuoteJson, 0)
	data := QuotesJson{quotes}

	// unmarshal the json data into data
	file, err := ioutil.ReadFile("./quotes.json")
	if err != nil {
		log.Fatal("Could not read quotes.json")
	}
	json.Unmarshal(file, &data)

	// receive a random integer between 0 and length of QArray -1
	rQuote := rand.Intn(len(data.QArray) - 1)
	// receive the quote from the array
	q := data.QArray[rQuote]

	// return the quote in the json response body
	json.NewEncoder(w).Encode(QuoteResponseJson{q.Author, q.Text})
}

// QuoteResponseJson is the response type for a citation request
type QuoteResponseJson struct {
	Author string `json:"author"`
	QText  string `json:"text"`
}

// QuotesJson contains an array of quotes (QuoteJson)
type QuotesJson struct {
	QArray []QuoteJson `json:"quotes"`
}

// QuotesJson contains a quote with author and text
type QuoteJson struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}
