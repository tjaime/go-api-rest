package main

import (
	"fmt"

	"github.com/tjaime/go-api-rest/database"
	"github.com/tjaime/go-api-rest/model"
	"github.com/tjaime/go-api-rest/route"
)

func main() {
	fmt.Println("Iniciando servidor rest com go")

	database.ConnectDB()

	model.Personalidades = []model.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	route.HandleRequest()
}
