package routes

import (
	"github.com/gorilla/mux"
)

/*InitRoutes inicializa todas las rutas para la aplicación*/
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetCategoryRoutes(router)
	return router
}
