package parser

type GroupedData struct {
	CsvData  [][][]string
	JsonData [][]JsonData
	XmlData  [][]Record
	YamlData []YamlData
}
