# go-csv-tag
Read csv file from go using tags

![cover.run go](https://cover.run/go/github.com/artonge/go-csv-tag.svg)
[![godoc for artonge/go-csv-tag](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/artonge/go-csv-tag)
[![goreportcard for artonge/go-csv-tag](https://goreportcard.com/badge/github.com/artonge/go-csv-tag)](https://goreportcard.com/report/artonge/go-csv-tag)
[![Sourcegraph for artonge/go-csv-tag](https://sourcegraph.com/github.com/artonge/go-csv-tag/-/badge.svg)](https://sourcegraph.com/github.com/artonge/go-csv-tag?badge)


# Install
`go get github.com/artonge/go-csv-tag`

# Usage
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

tab := []Demo{}                            // Create the slice where to put the file content
err  := csvtag.Load(csvtag.Config{         // Load your csv with the appropriate configuration
  path: "file.csv",                        // Path of the csv file
  dest: &tab,                              // A pointer to the create slice
  separator: ';',                          // Optional - if your csv use something else than ',' to separate values
  header: []string{"name", "ID", "number"} // Optional - if your csv does not contains a header
})
```

# Contribute
Pull requests are welcome ! :)
