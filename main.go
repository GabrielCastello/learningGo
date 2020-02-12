package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Docker Tutorial!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !!!!!")
		fmt.Println("Go Docker Tutorial!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
