# go-csv-tag
Read csv file from go using tags

# Install
`go get github.com/artonge/go-csv-tag`

# Usage
```go
tabT := []Test{}
err  := Load(Config{
  path: "csv_files/valid.csv",
  dest: &tabT,
  separator: ';',
  header: []string{"header1", "header2", "header3"}
})
```

# Contribute
Pull requests are welcome ! :)
