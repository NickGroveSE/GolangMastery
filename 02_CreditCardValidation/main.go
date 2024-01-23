package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
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
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Validation Request:\n%s", string(reqDump))
	io.WriteString(w, "Is Credit Card Number Valid? "+strconv.FormatBool(validate())+"\n")
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
