package csvtag

// CsvOptions - options when loading or dumping csv.
type CsvOptions struct {
	Separator rune
	UseCRLF   bool // True to use \r\n as the line terminator
	Header    []string
	TagKey    string
}

const DefaultTagKey string = "csv"
