package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tjaime/go-api-rest/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func PagePersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Personalidades)
}
