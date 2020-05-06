package csvtag

import (
	"bytes"
	"fmt"
	"testing"
)

var tabTest = []test{
	{"name", 1, 0.000001},
}

var tabTestNoID = []testNoID{
	{"name", 1, 0.000001},
}

func TestDumpToFileEmptyName(t *testing.T) {
	err := DumpToFile(tabTest, "")
	if err == nil {
		t.Fail()
	}
}

func TestDumpTestStruct(t *testing.T) {
	buffer := bytes.Buffer{}

	err := DumpToWriter(tabTest, &buffer)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if buffer.String() != "header1,header2,header3\nname,1,0.000001\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestDumpTestNoIdStruct(t *testing.T) {
	buffer := bytes.Buffer{}

	err := DumpToWriter(tabTestNoID, &buffer)
	if err != nil {
		t.Fail()
	}

	if buffer.String() != "header1,header\nname,0.000001\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestDumpTestStructPointer(t *testing.T) {
	buffer := bytes.Buffer{}

	err := DumpToWriter(&tabTest, &buffer)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if buffer.String() != "header1,header2,header3\nname,1,0.000001\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestDumpTestNoIdStructPointer(t *testing.T) {
	buffer := bytes.Buffer{}

	err := DumpToWriter(&tabTestNoID, &buffer)
	if err != nil {
		t.Fail()
	}

	if buffer.String() != "header1,header\nname,0.000001\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}
}

func TestDumpNilPointer(t *testing.T) {
	buffer := bytes.Buffer{}

	var n *test
	err := DumpToWriter(n, &buffer)
	if err == nil {
		t.Fail()
	}
}

func TestEmptyDumpToWriter(t *testing.T) {
	buffer := bytes.Buffer{}

	err := DumpToWriter([]test{}, &buffer)
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
	err := DumpToWriter(2, &buffer)
	if err == nil {
		t.Fail()
	}
}

func TestBigFloat(t *testing.T) {
	buffer := bytes.Buffer{}
	err := DumpToWriter(2, &buffer)
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
		{
			Name: "some name",
			ID:   1,
			Num:  0.000001,
		},
	}

	err := DumpToFile(tab, "csv_files/csv_file_name.csv")

	if err != nil {
		t.Fail()
	}
}

func TestDumpWithSemicolon(t *testing.T) {
	buffer := bytes.Buffer{}

	err := DumpToWriter(tabTest, &buffer, CsvOptions{Separator: ';'})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if buffer.String() != "header1;header2;header3\nname;1;0.000001\n" {
		fmt.Println(buffer.String())
		t.Fail()
	}

}
