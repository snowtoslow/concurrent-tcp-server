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

	header := []string{} // holds first row (header)

	for lineNum, record := range rawCSVData {

		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
		} else {
			// for each cell, map[string]string k=header v=value
			line := make(map[string]interface{})
			for i := 0; i < len(record); i++ {
				line[header[i]] = record[i]
			}

			groupedData.FullMap = append(groupedData.FullMap, line)
		}
	}

	return
}
