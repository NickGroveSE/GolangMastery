package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type CreditCard struct {
	CreditCardNumber string
	CardHolder       string
	ExpirationDate   string
	CCV              string
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	var card CreditCard

	body, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal([]byte(body), &card); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid JSON Request"))
	} else {
		// fmt.Printf("Card Number: %s\nCard Holder's Name: %s\nExpiration Date: %s\nCCV: %s\n", card.CreditCardNumber, card.CardHolder, card.ExpirationDate, card.CCV)
		fmt.Println("Validation Request Initialized...")
		io.WriteString(w, "Is Credit Card Number Valid? "+strconv.FormatBool(validate(card.CreditCardNumber))+"\n")
	}

}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":5050", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
