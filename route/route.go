package route

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tjaime/go-api-rest/controller"
	"github.com/tjaime/go-api-rest/middleware"
)

func HandleRequest() {
	route := mux.NewRouter()
	route.Use(middleware.ContentTypeJson)
	route.HandleFunc("/", controller.Home)
	route.HandleFunc("/api/personalidades", controller.PagePersonalidades).Methods("Get")
	route.HandleFunc("/api/personalidades/{id}", controller.FindPersonalidade).Methods("Get")
	route.HandleFunc("/api/personalidades", controller.InsertPersonalidade).Methods("Post")
	route.HandleFunc("/api/personalidades/{id}", controller.DeletePersonalidade).Methods("Delete")
	route.HandleFunc("/api/personalidades/{id}", controller.EditarPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(route)))
}
