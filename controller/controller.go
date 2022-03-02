package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tjaime/go-api-rest/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func PagePersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Personalidades)
}

func FindPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idConvertito, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}

	for _, personalidade := range model.Personalidades {
		if personalidade.Id == idConvertito {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
