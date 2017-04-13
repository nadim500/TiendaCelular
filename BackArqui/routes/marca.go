package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

/*SetMarcaRoutes son las rutas para manipular el modelo Category*/
func SetMarcaRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/marca", controllers.GetMarcas).Methods("GET")
	router.HandleFunc("/marca/{id}", controllers.GetMarcaByID).Methods("GET")
	router.HandleFunc("/marca", controllers.CreateMarca).Methods("POST")
	router.HandleFunc("/marca/{id}", controllers.UpdateMarca).Methods("PUT")
	router.HandleFunc("/marca/{id}", controllers.DeleteMarca).Methods("DELETE")
	return router
}
