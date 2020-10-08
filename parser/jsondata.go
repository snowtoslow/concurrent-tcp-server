package parser

import (
	"encoding/json"
	"log"
)

type JsonData struct{}

func (jsonData JsonData) Parse(dataToBeParsed string, data *GroupedData) (err error) {

	log.Println("Logic to parse json")

	var m []map[string]interface{}

	/*var arrayOfMaps []map[string]interface{}*/

	if dataToBeParsed[len(dataToBeParsed)-3] == 44 {
		dataToBeParsed = dataToBeParsed[:len(dataToBeParsed)-3] + "]"
	}
	err = json.Unmarshal([]byte(dataToBeParsed), &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		data.FullMap = append(data.FullMap, v)
		/*arrayOfMaps = append(arrayOfMaps,v)*/
	}

	return
}
