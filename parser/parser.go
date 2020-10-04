package parser

type ParseStrategy interface {
	Parse(message string)
}

var ParsePlatforms = map[string]ParseStrategy{"text/csv": CsvData{}, "json": JsonData{},
	"application/xml": XmlData{}, "application/x-yaml": YamlData{}}
