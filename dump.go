package csvtag

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

// DumpToWriter - writes a slice content into an io.Writer.
// @param slice: an object typically of the form []struct, where the struct is using csv tags.
// @param writer: the location of where you will write the slice content to. Example: File, Stdout, etc.
// @param options (optional): options for the csv parsing.
// @return an error if one occures.
func DumpToWriter(slice interface{}, writer io.Writer, options ...CsvOptions) error {
	// If slice is a pointer, get the value it points to.
	// (if it isn't, Indirect() does nothing and returns the value it was called with).
	reflectedValue := reflect.Indirect(reflect.ValueOf(slice))

	if reflectedValue.Kind() != reflect.Array && reflectedValue.Kind() != reflect.Slice {
		return errors.New("slice is not a slice")
	}

	option := CsvOptions{}
	if len(options) == 1 {
		option = options[0]
	}

	// Generate the header.
	if option.Header == nil {
		for i := 0; i < reflectedValue.Type().Elem().NumField(); i++ {
			name := reflectedValue.Type().Elem().Field(i).Tag.Get("csv")
			if name != "" {
				option.Header = append(option.Header, name)
			}
		}
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	if option.Separator != 0 {
		csvWriter.Comma = option.Separator
	}

	err := csvWriter.Write(option.Header)
	if err != nil {
		return err
	}

	for i := 0; i < reflectedValue.Len(); i++ {
		line := []string{}

		for j := 0; j < reflectedValue.Type().Elem().NumField(); j++ {
			valueRv := reflectedValue.Index(i)
			value := valueRv.Field(j)
			tag := valueRv.Type().Field(j).Tag.Get("csv")

			if tag == "" {
				continue
			}

			switch valueRv.Type().Field(j).Type.Kind() {
			case reflect.Float64, reflect.Float32:
				line = append(line, strconv.FormatFloat(value.Float(), 'f', -1, 64))
			default:
				line = append(line, fmt.Sprint(value))
			}
		}

		err = csvWriter.Write(line)
		if err != nil {
			return err
		}
	}

	return nil
}

// DumpToFile - writes a slice content into a file specified by path.
// @param slice: An object typically of the form []struct, where the struct is using csv tag.
// @param path: The file path string of where you want the file to be created.
// @param options (optional): options for the csv parsing.
// @return an error if one occures.
func DumpToFile(slice interface{}, path string, options ...CsvOptions) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	err = DumpToWriter(slice, file, options...)

	file.Close()

	return err
}

// DumpToString - writes a slice content into a string.
// @param slice: An object typically of the form []struct, where the struct is using csv tag.
// @param options (optional): options for the csv parsing.
// @return a string and an error if one occures.
func DumpToString(slice interface{}, options ...CsvOptions) (string, error) {

	writer := new(bytes.Buffer)

	err := DumpToWriter(slice, writer, options...)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}
