package route

import (
	"log"
	"net/http"

	"github.com/tjaime/go-api-rest/controller"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
