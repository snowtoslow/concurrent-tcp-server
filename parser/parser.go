package parser

import "fmt"

type ParseStrategy interface {
	Parse(message string)
}

var ParsePlatforms = map[string]ParseStrategy{"text/csv": CsvData{}, "json": JsonData{},
	"application/xml": XmlData{}, "application/x-yaml": YamlData{}}

func (csvData CsvData) Parse(message string) { fmt.Println("Logic to parse csv") }

func (jsonData JsonData) Parse(message string) { fmt.Println("Logic to parse json") }

func (yamlData YamlData) Parse(message string) { fmt.Println("Logic to parse yaml") }

func (xmlData XmlData) Parse(message string) { fmt.Println("Logic to parse xml") }

type CsvData struct{}

type JsonData struct{}

type YamlData struct{}

type XmlData struct{}
