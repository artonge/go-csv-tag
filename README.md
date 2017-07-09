# go-csv-tag
Read csv file from go using tags

[![godoc for artonge/go-csv-tag](https://godoc.org/github.com/artonge/go-csv-tag?status.svg)](http://godoc.org/github.com/artonge/go-csv-tag)

[![Build Status](https://travis-ci.org/artonge/go-csv-tag.svg?branch=master)](https://travis-ci.org/artonge/go-csv-tag)
![cover.run go](https://cover.run/go/github.com/artonge/go-csv-tag.svg)
[![goreportcard for artonge/go-csv-tag](https://goreportcard.com/badge/github.com/artonge/go-csv-tag)](https://goreportcard.com/report/artonge/go-csv-tag)

[![Sourcegraph for artonge/go-csv-tag](https://sourcegraph.com/github.com/artonge/go-csv-tag/-/badge.svg)](https://sourcegraph.com/github.com/artonge/go-csv-tag?badge)


# Install
`go get github.com/artonge/go-csv-tag`

# Example
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

## TODO
- [ ] Add `Dump(data interface{}, file string) error` function to write some datas the disk with csv format
- [ ] Update `Load` to also match csv fields with property name (case sensitive and lowercases) 
