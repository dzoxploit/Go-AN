package models

import (
	"encoding/json"
	"log"

	"github.com/dzoxploit/eratani-test-case/apitest/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Creditcard struct {
	ID       int64 `json:"id"`
	Country       string `json:"country"`
	CreditCardType     string   `json:"credit_card_type"`
	CreditCard      float64 `json:"credit_card"`
	FirstName      string `json:"first_name"`
	LastName string `json:"last_name"`
}



func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Creditcard{})
}

func (b *Creditcard) CreateCreditCard() *Creditcard {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllCreditCard(pageNumber int, pageSize int, search string) ([]Creditcard, error) {


	// Calculate the offset based on the page number and page size
	offset := (pageNumber - 1) * pageSize

	// Query the database with pagination
	var creditcards []Creditcard
	result := db.Where("credit_card_type LIKE ? OR first_name LIKE ? OR last_name LIKE ? OR country LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").Order("id DESC").Offset(offset).Limit(pageSize).Find(&creditcards)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return creditcards, nil
}


func GetCountryHighSpend() ([]byte, error) {

	var resultHighestCountry []struct {
		Country string   `json:"country"`
		TotalSpend  float64 `json:"total_spend"`
	}

	result := db.Table("creditcards").
				Select("country, SUM(REPLACE(credit_card, ',', '')) as total_spend").
				Group("country").
				Order("total_spend DESC").
				Limit(1).
				Scan(&resultHighestCountry)

	if result.Error != nil {
    	panic(result.Error)
	}

	jsonData, err := json.Marshal(resultHighestCountry)
	
	if err != nil {
		panic(err)
	}


	return jsonData, nil
}

func GetHighestCreditCardType() ([]byte, error){
	var resultHighestCreditCardType []struct {
		CreditCardType string   `json:"credit_card_type "`
		Total  float64 `json:"total"`
	}

	// Retrieve the most frequent credit card types
	result := db.Table("creditcards").
		Select("credit_card_type, COUNT(*) as total").
		Group("credit_card_type").
		Order("total DESC").
		Limit(1).
		Scan(&resultHighestCreditCardType)

	if result.Error != nil {
		panic(result.Error)
	}

	jsonData, err := json.Marshal(resultHighestCreditCardType)

	if err != nil {
		panic(err)
	}

	return jsonData, nil

}


// func GetAllContact() []Contact {
// 	var Contact []Contact
// 	db.Find(&Contact)
// 	return Contact
// }

func GetCreditCardById(Id int64) (*Creditcard , *gorm.DB){
	var getCreditCard Creditcard
	db:=db.Where("ID = ?", Id).Find(&getCreditCard)
	return &getCreditCard, db
}

func DeleteCreditCard(Id int64) Creditcard {
	var creditcards Creditcard
	db.Where("ID = ?", Id).Delete(creditcards)
	return creditcards
}