package parser

import (
	"encoding/json"
	"log"
)

func (jsonData JsonData) Parse(dataToBeParsed string, data *GroupedData) (err error) {
	log.Println("Logic to parse json")

	var myResp []JsonData
	if dataToBeParsed[len(dataToBeParsed)-3] == 44 {
		dataToBeParsed = dataToBeParsed[:len(dataToBeParsed)-3] + "]"
	}

	err = json.Unmarshal([]byte(dataToBeParsed), &myResp)
	if err != nil {
		return
	}
	data.JsonData = append(data.JsonData, myResp)
	return
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
