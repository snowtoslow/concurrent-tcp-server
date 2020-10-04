package parser

import (
	"encoding/csv"
	"log"
	"strings"
)

type CsvData struct{}

func (csvData CsvData) Parse(dataToBeParsed string) {
	log.Println("Logic to parse csv")

	reader := csv.NewReader(strings.NewReader(dataToBeParsed))
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // see the Reader struct information below
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	log.Println(rawCSVData)
	// logic for search
	/*for _, row := range rawCSVdata {
		for _, col := range row {
			if col == stringToFindColumnBy{
				log.Println(row)
			}
		}
	}*/

}
