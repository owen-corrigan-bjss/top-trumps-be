package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	portNum := ":8000"
	fmt.Printf("Server is running at http://localhost%s", portNum)
	log.Fatal(http.ListenAndServe(portNum, nil))
}
