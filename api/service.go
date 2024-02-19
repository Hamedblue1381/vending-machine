package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiService struct {
	router *mux.Router
}

func NewApiService() *ApiService {
	router := mux.NewRouter()
	return &ApiService{
		router: router,
	}
}
func (a *ApiService) Start(h VendingMachineHandler) {

	a.router.HandleFunc("/item/{id}", h.ChooseItem).Methods("GET")

	a.router.HandleFunc("/add-machine", h.CreateVendingMachine).Methods("POST")
	a.router.HandleFunc("/machines", h.GetVendingMachines).Methods("GET")
	a.router.HandleFunc("/insert", h.InsertCoin).Methods("GET")
	log.Println("Running on port 8080..")
	http.ListenAndServe(":8080", a.router)

}
