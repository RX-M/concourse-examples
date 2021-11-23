package main

import (
	"fmt"
	"net/http"
)

func AddSix(i int) int {
	return i + 6
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := "Hello World!"
	fmt.Fprintln(w, response)
	fmt.Println("Processing hello request.")
}

func listenAndServe(port string) {
	fmt.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := "8080"
	go listenAndServe(port)
	select {}
}
