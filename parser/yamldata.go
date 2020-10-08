package parser

import (
	"concurrent-tcp-server/utils"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"reflect"
)

type YamlData []struct {
	ID           int    `yaml:"id"`
	FirstName    string `yaml:"first_name"`
	LastName     string `yaml:"last_name"`
	CardNumber   string `yaml:"card_number"`
	CardBalance  string `yaml:"card_balance,omitempty"`
	CardCurrency string `yaml:"card_currency,omitempty"`
}

func (yamlData YamlData) Parse(dataToBeParsed string, groupedData *GroupedData) (err error) {
	log.Println("Logic to parse yaml")
	y := YamlData{}

	err = yaml.Unmarshal([]byte(dataToBeParsed), &y)
	if err != nil {
		return
	}
	y.createYamlMap()
	groupedData.FullMap = append(groupedData.FullMap, y.createYamlMap()) //DONE
	return
}

func (yamlData YamlData) createYamlMap() (testMap map[string]interface{}) {
	testMap = make(map[string]interface{})
	for i := 0; i < len(yamlData); i++ {
		v := reflect.ValueOf(yamlData[i])
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			testMap[utils.ToSnakeCase(typeOfS.Field(i).Name)] = fmt.Sprintf("%v", v.Field(i).Interface())
		}
	}
	return testMap
}
