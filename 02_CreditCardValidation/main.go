package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Validation Request\n")
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
