package parser

import (
	"encoding/json"
	"fmt"
	"log"
)

func (jsonData JsonData) Parse(dataToBeParsed string) {
	log.Println("Logic to parse json")

	var myResp []JsonData

	if dataToBeParsed[len(dataToBeParsed)-2] == 44 {
		dataToBeParsed = dataToBeParsed[:len(dataToBeParsed)-2] + "]"
	}

	err := json.Unmarshal([]byte(dataToBeParsed), &myResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(myResp)
}

type JsonData struct {
	Id           int8   `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	CardNumber   string `json:"card_number,omitempty"`
	CardBalance  string `json:"card_balance,omitempty"`
	CardCurrency string `json:"card_currency,omitempty"`
	Email        string `json:"email,omitempty"`
	Organization string `json:"organization,omitempty"`
	FullName     string `json:"full_name,omitempty"`
	EmployeeId   string `json:"employee_id,omitempty"`
}
