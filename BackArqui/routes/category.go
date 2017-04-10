package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

/*SetCategoryRoutes son las rutas para manipular el modelo Category*/
func SetCategoryRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/category", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	return router
}
