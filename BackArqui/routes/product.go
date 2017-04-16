package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

/*SetProductRoutes son las rutas para manipular el modelo Product*/
func SetProductRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/product", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProductByID).Methods("GET")
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	return router
}
