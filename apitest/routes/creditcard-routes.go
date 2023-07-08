package routes

import (
	"github.com/dzoxploit/eratani-test-case/apitest/controllers"
	"github.com/gorilla/mux"
)

var RegisterCreditCardRoutes = func(router *mux.Router) {
	router.HandleFunc("/credit-card/", controllers.CreateCreditCard).Methods("POST")
	router.HandleFunc("/credit-cards", controllers.GetCreditCard).Methods("GET")
	router.HandleFunc("/get-highest-credit-card-countries", controllers.GetCountryHighSpend).Methods("GET")
	router.HandleFunc("/get-highest-credit-card-type", controllers.GetHighestCreditCardType).Methods("GET")
	router.HandleFunc("/credit-card/{Id}", controllers.GetCreditCardById).Methods("GET")
	router.HandleFunc("/credit-card/{Id}", controllers.UpdateCreditCard).Methods("PUT")
	router.HandleFunc("/credit-card/{Id}", controllers.DeleteCreditCard).Methods("DELETE")
}