package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 5001
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server Running on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
