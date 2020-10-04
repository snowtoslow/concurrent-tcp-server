package parser

import (
	"gopkg.in/yaml.v2"
	"log"
)

func (yamlData YamlData) Parse(dataToBeParsed string) {
	log.Println("Logic to parse yaml")
	y := YamlData{}

	err := yaml.Unmarshal([]byte(dataToBeParsed), &y)
	if err != nil {
		log.Println(err)
	}

	log.Println(y)
}

type YamlData []struct {
	ID           int    `yaml:"id"`
	FirstName    string `yaml:"first_name"`
	LastName     string `yaml:"last_name"`
	CardNumber   string `yaml:"card_number"`
	CardBalance  string `yaml:"card_balance,omitempty"`
	CardCurrency string `yaml:"card_currency,omitempty"`
}
