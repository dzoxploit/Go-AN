package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dzoxploit/eratani-test-case/apitest/models"
	"github.com/dzoxploit/eratani-test-case/apitest/utils"
	"github.com/gorilla/mux"
)

var NewCreditCard models.Creditcard

func CreateCreditCard(w http.ResponseWriter, r *http.Request) {

	CreateCreditCard := &models.Creditcard{}
	utils.ParseBody(r, CreateCreditCard)
	b:= CreateCreditCard.CreateCreditCard()
	
	fmt.Print(b)
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCountryHighSpend(w http.ResponseWriter, r *http.Request) {
	// Call the models function to retrieve the country with the highest spend
	jsonData, err := models.GetCountryHighSpend()
	if err != nil {
		http.Error(w, "Failed to retrieve country with highest spend", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetHighestCreditCardType(w http.ResponseWriter, r *http.Request) {
	// Call the model function to retrieve the highest credit card types
	jsonData, err := models.GetHighestCreditCardType()
	if err != nil {
		http.Error(w, "Failed to retrieve highest credit card types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}


func GetCreditCard(w http.ResponseWriter, r *http.Request) {
	// Extract the page number and page size from query parameters
	pageNumber, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	pageSearch := r.URL.Query().Get("search")

	// Call the model function to get contacts with pagination
	contacts, err := models.GetAllCreditCard(pageNumber, pageSize, pageSearch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the contacts to JSON
	response, err := json.Marshal(contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetCreditCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creditCardIdStr := vars["Id"]

	if creditCardIdStr == "" {
		http.Error(w, "Credit card ID is missing", http.StatusBadRequest)
		return
	}

	creditCardId, err := strconv.ParseInt(creditCardIdStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid credit card ID", http.StatusBadRequest)
		return
	}

	creditcardDetails, _:= models.GetCreditCardById(creditCardId)
	
	res, _ := json.Marshal(creditcardDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCreditCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creditCardIDStr := vars["Id"]
	
	var updateCreditCard models.Creditcard

	// Validate creditCardID for null or empty
	if creditCardIDStr == "" {
		http.Error(w, "Credit card ID is missing", http.StatusBadRequest)
		return
	}

	creditCardID, err := strconv.ParseInt(creditCardIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid credit card ID", http.StatusBadRequest)
		return
	}

	utils.ParseBody(r, &updateCreditCard)

	creditcardDetails, db := models.GetCreditCardById(creditCardID)

	if updateCreditCard.Country != "" {
		creditcardDetails.Country = updateCreditCard.Country
	}

	if updateCreditCard.CreditCardType != "" {
		creditcardDetails.CreditCardType = updateCreditCard.CreditCardType
	}

	if updateCreditCard.CreditCard != 0 {
		creditcardDetails.CreditCard = updateCreditCard.CreditCard
	}

	if updateCreditCard.FirstName != "" {
		creditcardDetails.FirstName = updateCreditCard.FirstName
	}

	if updateCreditCard.LastName != "" {
		creditcardDetails.LastName = updateCreditCard.LastName
	}

	db.Save(&creditcardDetails)

	res, _ := json.Marshal(creditcardDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteCreditCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creditCardIDStr := vars["Id"]

	// Validate creditCardID for null or empty
	if creditCardIDStr == "" {
		http.Error(w, "Credit card ID is missing", http.StatusBadRequest)
		return
	}

	creditCardId, err := strconv.ParseInt(creditCardIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid credit card ID", http.StatusBadRequest)
		return
	}

	creditcard:= models.DeleteCreditCard(creditCardId)
	res, _ := json.Marshal(creditcard)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}