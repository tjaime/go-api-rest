package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tjaime/go-api-rest/controller"
)

func HandleRequest() {
	route := mux.NewRouter()

	route.HandleFunc("/", controller.Home)
	route.HandleFunc("/api/personalidades", controller.PagePersonalidades).Methods("Get")
	route.HandleFunc("/api/personalidades/{id}", controller.FindPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", route))
}
