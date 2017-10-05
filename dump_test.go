package csvtag

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	TESTSTRUCT     = "header1,header2,header3\nline1,1,1.2\nline2,2,2.3\nline3,3,3.4"
	TESTNOIDSTRUCT = "header1,header\nline1,0\nline2,0\nline3,0"
	FILEERROR      = "open : no such file or directory"
)

var tabTest = []test{
	test{"name", 1, 42.5},
}

var tabTestNoID = []testNoID{
	testNoID{"name", 1, 42.5},
}

func TestDumpToFileEmptyName(t *testing.T) {
	err := DumpToFile(tabTest, "")
	if err == nil {
		t.Fail()
	}
}

func TestDumpTestStruct(t *testing.T) {
	buffer := bytes.Buffer{}

	err := Dump(tabTest, &buffer)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if buffer.String() != "header1,header2,header3\nname,1,42.5\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestDumpTestNoIdStruct(t *testing.T) {
	buffer := bytes.Buffer{}

	err := Dump(tabTestNoID, &buffer)
	if err != nil {
		t.Fail()
	}

	if buffer.String() != "header1,header\nname,42.5\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestEmptyDump(t *testing.T) {
	buffer := bytes.Buffer{}

	err := Dump([]test{}, &buffer)
	if err != nil {
		t.Fail()
	}

	if buffer.String() != "header1,header2,header3\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestWrongType(t *testing.T) {
	buffer := bytes.Buffer{}
	err := Dump(2, &buffer)
	if err == nil {
		t.Fail()
	}
}

type Demo struct { // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

func TestREADMEExample(t *testing.T) {

	tab := []Demo{ // Create the slice where to put the file content
		Demo{
			Name: "some name",
			ID:   1,
			Num:  42.5,
		},
	}

	err := DumpToFile(tab, "csv_files/csv_file_name.csv")

	if err != nil {
		t.Fail()
	}
}
