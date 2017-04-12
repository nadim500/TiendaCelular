package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

/*SetCategoryRoutes son las rutas para manipular el modelo Category*/
func SetCategoryRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/category", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/category/{id}", controllers.GetCategoryByID).Methods("GET")
	router.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", controllers.DeleteCategory).Methods("DELETE")
	return router
}
