package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CreditCard struct {
	CreditCardNumber string
	CardHolder       string
	ExpirationDate   string
	CCV              string
}

type CreditCardValidation struct {
	CreditCardNumber string
	IsValid          bool
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	var card CreditCard
	var cardValidation CreditCardValidation

	body, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal([]byte(body), &card); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid JSON Request"))
	} else {
		// fmt.Printf("Card Number: %s\nCard Holder's Name: %s\nExpiration Date: %s\nCCV: %s\n", card.CreditCardNumber, card.CardHolder, card.ExpirationDate, card.CCV)
		fmt.Println("Validation Request Initialized...")
		io.WriteString(w, "Hi "+card.CardHolder+" Validation Complete\n")

		cardValidation.CreditCardNumber = card.CreditCardNumber
		cardValidation.IsValid = validate(card.CreditCardNumber)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cardValidation)

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
