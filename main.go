package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando servidor rest com go")
	HandleRequest()
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
