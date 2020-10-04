package parser

import (
	"encoding/xml"
	"log"
	"regexp"
)

func (xmlData XmlData) Parse(dataToBeParsed string) {
	log.Println("Logic to parse xml")

	data := XmlData{}

	regexToDeleteNeLines := regexp.MustCompile(`\r?\n`)
	rawXmlWithOutNewLines := regexToDeleteNeLines.ReplaceAllString(dataToBeParsed, "")

	if err := xml.Unmarshal([]byte(rawXmlWithOutNewLines), &data); err != nil {
		log.Println(err)
	} else {
		log.Println(data.RecordList)
	}
}

type XmlData struct {
	XMLName    xml.Name `xml:"dataset"`
	RecordList []Record `xml:"record"`
}

type Record struct {
	XMLName        xml.Name `xml:"record"`
	Id             int8     `xml:"id"`
	FirstName      string   `xml:"first_name"`
	LastName       string   `xml:"last_name"`
	BitcoinAddress string   `xml:"bitcoin_address"`
}
