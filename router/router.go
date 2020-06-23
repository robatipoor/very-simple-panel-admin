package router

import (
	"github.com/gorilla/mux"
	"github.com/robatipoor/very-simple-panel-admin/handler"
)

// GetRouters app
func GetRouters() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handler.Home).Methods("GET")
	mux.HandleFunc("/AddProduct", handler.AddProduct).Methods("POST")
	mux.HandleFunc("/GetProduct/{id}", handler.GetProduct).Methods("GET")
	mux.HandleFunc("/GetAllProduct", handler.GetAllProduct).Methods("GET")
	mux.HandleFunc("/DeleteProduct/{id}", handler.DeleteProduct).Methods("DELETE")
	mux.HandleFunc("/UpdateProduct/{id}", handler.UpdateProduct).Methods("PUT")
	return mux
}
