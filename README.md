# go-csv-tag

Read csv file from go using tags

[![godoc for artonge/go-csv-tag](https://godoc.org/github.com/artonge/go-csv-tag?status.svg)](http://godoc.org/github.com/artonge/go-csv-tag)

![Go](https://github.com/artonge/go-csv-tag/workflows/Go/badge.svg)
[![goreportcard for artonge/go-csv-tag](https://goreportcard.com/badge/github.com/artonge/go-csv-tag)](https://goreportcard.com/report/artonge/go-csv-tag)

[![Sourcegraph for artonge/go-csv-tag](https://sourcegraph.com/github.com/artonge/go-csv-tag/-/badge.svg)](https://sourcegraph.com/github.com/artonge/go-csv-tag?badge)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

# Install

`go get github.com/artonge/go-csv-tag/v2`

# Example

## Load

The csv file:

```csv
name, ID, number
name1, 1, 1.2
name2, 2, 2.3
name3, 3, 3.4
```

Your go code:

```go
type Demo struct {                                // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

tab := []Demo{}                                   // Create the slice where to put the content
err  := csvtag.LoadFromPath(
	"file.csv",                                   // Path of the csv file
	&tab,                                         // A pointer to the create slice
	csvtag.CsvOptions{                            // Load your csv with optional options
		Separator: ';',                           // changes the values separator, default to ','
		Header: []string{"name", "ID", "number"}, // specify custom headers
})
```

You can also load the data from an io.Reader with:

```go
csvtag.LoadFromPath(youReader, &tab)
```

Or from a string with:

```go
csvtag.LoadFromString(yourString, &tab)
```

## Dump

You go code:

```go
type Demo struct {                         // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

tab := []Demo{                             // Create the slice where to put the content
	Demo{
		Name: "some name",
		ID: 1,
		Num: 42.5,
	},
}

err := csvtag.DumpToFile(tab, "csv_file_name.csv")
```

You can also dump the data into an io.Writer with:

```go
err := csvtag.DumpToWriter(tab, yourIOWriter)
```

Or dump to a string with:

```go
str, err := csvtag.DumpToString(tab)
```

The csv file written:

```csv
name,ID,number
some name,1,42.5
```
