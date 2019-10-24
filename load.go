package csvtag

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// LoadFromReader - Load csv from an io.Reader and put it in a array of the destination's type using tags.
// Example:
// 	tabOfMyStruct := []MyStruct{}
// 	err  := Load(
// 				myIoReader,
// 				&tabOfMyStruct,
// 				CsvOptions{
// 					Separator: ';',
// 					Header: []string{"header1", "header2", "header3"
// 				}
// 			})
// @param file: the io.Reader.
// @param destination: object where to store the result.
// @param options (optional): options for the csv parsing.
// @return an error if one occurs.
func LoadFromReader(file io.Reader, destination interface{}, options ...CsvOptions) error {
	if len(options) > 1 {
		return fmt.Errorf("error you can only pass one CsvOptions argument")
	}

	option := CsvOptions{}
	if len(options) == 1 {
		option = options[0]
	}

	header, content, err := readFile(file, option.Separator, option.Header)
	if err != nil {
		return fmt.Errorf("error reading csv from io.Reader: %v", err)
	}

	// This means that the file is empty, so just return nil.
	if content == nil {
		return nil
	}

	err = mapToDestination(header, content, destination)
	if err != nil {
		return fmt.Errorf("error mapping the content to the destination\n	==> %v", err)
	}

	return nil
}

// LoadFromPath - Load csv from a path and put it in a array of the destination's type using tags.
// Example:
// 	tabOfMyStruct := []MyStruct{}
// 	err  := Load(
// 				"my_csv_file.csv",
// 				&tabOfMyStruct,
// 				CsvOptions{
// 					Separator: ';',
// 					Header: []string{"header1", "header2", "header3"
// 				}
// 			})
// @param path: the path of the csv file.
// @param destination: object where to store the result.
// @param options (optional): options for the csv parsing.
// @return an error if one occurs.
func LoadFromPath(path string, destination interface{}, options ...CsvOptions) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	err = LoadFromReader(file, destination, options...)
	if err != nil {
		return fmt.Errorf("error mapping csv from path %v:\n	==> %v", path, err)
	}

	return nil
}

// LoadFromString - Load csv from string and put it in a array of the destination's type using tags.
// Example:
// 	tabOfMyStruct := []MyStruct{}
// 	err  := Load(
// 				myString,
// 				&tabOfMyStruct,
// 				CsvOptions{
// 					Separator: ';',
// 					Header: []string{"header1", "header2", "header3"
// 				}
// 			})
// @param str: the string.
// @param destination: object where to store the result.
// @param options (optional): options for the csv parsing.
// @return an error if one occurs.
func LoadFromString(str string, destination interface{}, options ...CsvOptions) error {
	return LoadFromReader(strings.NewReader(str), destination, options...)
}

// Load the header and file content in memory.
// @param file: the io.Reader to read from.
// @param separator: the separator used in the csv file.
// @param header: the optional header if the file does not contain one.
func readFile(file io.Reader, separator rune, header []string) ([]string, [][]string, error) {
	// Create and configure the csv reader.
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	if separator != 0 {
		reader.Comma = separator
	}

	// We need to read it all at once to have the number of records for the array creation.
	content, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	// If file is empty, return.
	if len(content) == 0 {
		return nil, nil, nil
	}

	// If no header is provided, treat first line as the header.
	if header == nil {
		header = content[0]
		content = content[1:]
	}

	return header, content, nil
}

// Map the provided content to the destination using the header and the tags.
// @param header: the csv header to match with the struct's tags.
// @param content: the content to put in destination.
// @param destination: the destination where to put the file's content.
func mapToDestination(header []string, content [][]string, destination interface{}) error {
	if destination == nil {
		return fmt.Errorf("destination slice is nil")
	}

	if reflect.TypeOf(destination).Elem().Kind() != reflect.Slice {
		return fmt.Errorf("destination is not a slice")
	}

	// Map each header name to its index.
	headerMap := make(map[string]int)
	for i, name := range header {
		headerMap[strings.TrimSpace(name)] = i
	}

	// Create the slice to put the values in.
	sliceRv := reflect.MakeSlice(
		reflect.ValueOf(destination).Elem().Type(),
		len(content),
		len(content),
	)

	for i, line := range content {
		emptyStruct := sliceRv.Index(i)

		for j := 0; j < emptyStruct.NumField(); j++ {
			propertyTag := emptyStruct.Type().Field(j).Tag.Get("csv")
			if propertyTag == "" {
				continue
			}

			propertyPosition, ok := headerMap[propertyTag]
			if !ok {
				continue
			}

			err := storeValue(line[propertyPosition], emptyStruct.Field(j))
			if err != nil {
				return fmt.Errorf("line: %v to slice: %v:\n	==> %v", line, emptyStruct, err)
			}
		}
	}

	reflect.ValueOf(destination).Elem().Set(sliceRv)

	return nil
}

// Set the value of the valRv to rawValue.
// @param rawValue: the value, as a string, that we want to store.
// @param valRv: the reflected value where we want to store our value.
// @return an error if one occurs.
func storeValue(rawValue string, valRv reflect.Value) error {
	switch valRv.Kind() {
	case reflect.String:
		valRv.SetString(rawValue)
	case reflect.Int:
		value, err := strconv.ParseInt(rawValue, 10, 64)
		if err != nil && rawValue != "" {
			return fmt.Errorf("error parsing int '%v':\n	==> %v", rawValue, err)

		}
		valRv.SetInt(value)
	case reflect.Float64:
		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil && rawValue != "" {
			return fmt.Errorf("error parsing float '%v':\n	==> %v", rawValue, err)
		}
		valRv.SetFloat(value)
	}

	return nil
}
