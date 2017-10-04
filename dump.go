package csvtag

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

const (
	//UNSUPPORTED is default text for the unsupported cases
	UNSUPPORTED = "Unsupported data type"
)

//This function is used to obtain the csv tags labels that will be used as headers in the csv file
//@param v: reflect.Value struct
func getHeadersFromStruct(v reflect.Value) []string {
	var headers []string

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("csv")
		if name != "" {
			headers = append(headers, name)
		}
	}
	return headers
}

//The function gets the data from the reflect.Value of Slice.
//@param v: The reflect.Value or Slice
//@param headers: Slice of strings used to determine which fields needs to be extracted from the struct
func getDataFromSlice(v reflect.Value, headers []string) [][]string {
	var result [][]string
	for i := 0; i < v.Len(); i++ {
		data := []string{}
		fields := v.Index(i).NumField()
		for q := 0; q < fields; q++ {
			if inSlice(headers, v.Index(i).Type().Field(q).Tag.Get("csv")) == true {
				data = append(data, fmt.Sprint(v.Index(i).Field(q)))
			}
		}
		result = append(result, data)
	}
	return result
}

//This function is used to determine if the csv tag label is in the list of fields to be used for the csv file
//@param s: slice of string representing the headers
//@parem item: The string being checked
func inSlice(s []string, item string) bool {
	var result bool
	for _, elem := range s {
		if strings.Compare(elem, item) == 0 {
			result = true
		}
	}
	return result
}

//DumpToFile - writes dat to a file
//@param data: An object typically of the form []struct, where the struct using csv tag
//@param filePath: The file path string of where you want the file to be created
func DumpToFile(data interface{}, filePath string) error {
	//Create file object
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	//Dump data to file
	err = Dump(data, file)
	return err
}

//Dump - writes data to an io.Writer
//@param data: An object typically of the form []struct, where the struct using csv tags
//@param w: The location of where you will write the data to. Example: File, Stdout, etc
func Dump(data interface{}, w io.Writer) error {
	// A placeholder for any potential errors
	var err error
	// A placeholder for the csv header
	var header []string

	// A placeholder for the csv body
	var body [][]string

	// Determines the value of the passed in data type
	v := reflect.ValueOf(data)

	// Deals with the different cases (based on data type)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		elemStruct := v.Index(0)
		header = getHeadersFromStruct(elemStruct)
		body = getDataFromSlice(v, header)

	default:
		err = errors.New(UNSUPPORTED)
		return err
	}

	//Write to csv file
	csvW := csv.NewWriter(w)
	defer csvW.Flush() // Ensures that all data in the buffer has been written to io.Writer

	err = csvW.Write(header)
	if err != nil {
		return err
	}

	err = csvW.WriteAll(body)

	return err
}
