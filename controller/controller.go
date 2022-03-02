package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tjaime/go-api-rest/database"
	"github.com/tjaime/go-api-rest/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func PagePersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []model.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func FindPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idConvertito, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	var personalidade model.Personalidade
	database.DB.First(&personalidade, idConvertito)
	json.NewEncoder(w).Encode(personalidade)
}

func InsertPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade model.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idConvertito, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	var personalidade model.Personalidade
	database.DB.Delete(&personalidade, idConvertito)

	json.NewEncoder(w).Encode(personalidade)
}
