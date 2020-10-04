package parser

type ParseStrategy interface {
	Parse(message string, data *GroupedData) (err error)
}

var ParsePlatforms = map[string]ParseStrategy{"text/csv": CsvData{}, "json": JsonData{},
	"application/xml": XmlData{}, "application/x-yaml": YamlData{}}
