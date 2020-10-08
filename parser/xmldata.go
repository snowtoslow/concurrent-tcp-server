package parser

import (
	"concurrent-tcp-server/utils"
	"encoding/xml"
	"fmt"
	"log"
	"reflect"
	"regexp"
)

type XmlData struct {
	XMLName    xml.Name `xml:"dataset"`
	RecordList []Record `xml:"record"`
}

type Record struct {
	Id             int8   `xml:"id"`
	FirstName      string `xml:"first_name"`
	LastName       string `xml:"last_name"`
	BitcoinAddress string `xml:"bitcoin_address"`
}

func (xmlData XmlData) Parse(dataToBeParsed string, groupedData *GroupedData) (err error) {
	log.Println("Logic to parse xml")

	data := XmlData{}

	regexToDeleteNeLines := regexp.MustCompile(`\r?\n`)
	rawXmlWithOutNewLines := regexToDeleteNeLines.ReplaceAllString(dataToBeParsed, "")

	if err := xml.Unmarshal([]byte(rawXmlWithOutNewLines), &data); err != nil {
		return err
	} else {
		/*groupedData.fullMap = append(groupedData.fullMap, data.RecordList)*/ //PASSED
		for i := 0; i < len(data.RecordList); i++ {
			groupedData.FullMap = append(groupedData.FullMap, data.RecordList[i].createSingleMap())
		}
	}
	return
}

func (r Record) createSingleMap() (testMap map[string]interface{}) {
	testMap = make(map[string]interface{})
	v := reflect.ValueOf(r)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		testMap[utils.ToSnakeCase(typeOfS.Field(i).Name)] = fmt.Sprintf("%v", v.Field(i).Interface())
	}

	return
}
