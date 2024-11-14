package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("INSIDE THE RANDOM ROUTE")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HELLO THERE")
	})

	fmt.Println("STARTING THE SERVER")
	if err := http.ListenAndServe(":8080", nil); err == nil {
		panic("ERROR OCCURE WHILE RUNING SERVER")
	}
}
