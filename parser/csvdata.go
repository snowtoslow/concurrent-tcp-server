package parser

import (
	"encoding/csv"
	"log"
	"strings"
)

type CsvData struct{}

func (csvData CsvData) Parse(dataToBeParsed string, groupedData *GroupedData) (err error) {
	log.Println("Logic to parse csv")

	reader := csv.NewReader(strings.NewReader(dataToBeParsed))
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // see the Reader struct information below
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		return
	}
	groupedData.CsvData = append(groupedData.CsvData, rawCSVData[1:])
	return
}

// logic for search
/*for _, row := range rawCSVdata {
	for _, col := range row {
		if col == stringToFindColumnBy{
			log.Println(row)
		}
	}
}*/
