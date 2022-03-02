package main

import (
	"fmt"

	"github.com/tjaime/go-api-rest/route"
)

func main() {
	fmt.Println("Iniciando servidor rest com go")
	route.HandleRequest()
}
