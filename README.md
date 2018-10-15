# go-csv-tag
Read csv file from go using tags

[![godoc for artonge/go-csv-tag](https://godoc.org/github.com/artonge/go-csv-tag?status.svg)](http://godoc.org/github.com/artonge/go-csv-tag)

[![Build Status](https://travis-ci.org/artonge/go-csv-tag.svg?branch=master)](https://travis-ci.org/artonge/go-csv-tag)
![cover.run go](https://cover.run/go/github.com/artonge/go-csv-tag.svg)
[![goreportcard for artonge/go-csv-tag](https://goreportcard.com/badge/github.com/artonge/go-csv-tag)](https://goreportcard.com/report/artonge/go-csv-tag)

[![Sourcegraph for artonge/go-csv-tag](https://sourcegraph.com/github.com/artonge/go-csv-tag/-/badge.svg)](https://sourcegraph.com/github.com/artonge/go-csv-tag?badge)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) 

# Install
`go get github.com/artonge/go-csv-tag`

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
type Demo struct {                         // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

tab := []Demo{}                             // Create the slice where to put the file content
err  := csvtag.Load(csvtag.Config{          // Load your csv with the appropriate configuration
  Path: "file.csv",                         // Path of the csv file
  Dest: &tab,                               // A pointer to the create slice
  Separator: ';',                           // Optional - if your csv use something else than ',' to separate values
  Header: []string{"name", "ID", "number"}, // Optional - if your csv does not contains a header
})
```

## Dump
You go code:
```go
type Demo struct {                         // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

tab := []Demo{                             // Create the slice where to put the file content
	Demo{
		Name: "some name",
		ID: 1,
		Num: 42.5,
	},
}

err := csvtag.DumpToFile(tab, "csv_file_name.csv")
```
You can also dump the data into an io.Writer with
```go
err := csvtag.Dump(tab, yourIOWriter)
```
The csv file written:
```csv
name,ID,number
some name,1,42.5
```
