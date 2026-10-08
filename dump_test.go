package csvtag

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

var tabTest = []test{
	{"name", 1, 0.000001},
}

var tabTestTagKey = []testTagKey{
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

type dumpFailWriter struct {
	writes int
	failAt int
	err    error
	short  bool
}

func (w *dumpFailWriter) Write(p []byte) (int, error) {
	w.writes++
	if w.writes < w.failAt {
		return len(p), nil
	}
	if w.short {
		return len(p) - 1, nil
	}
	return 0, w.err
}

func TestDumpToWriterWriteErrors(t *testing.T) {
	writeErr := errors.New("destination write failed")
	small := []Demo{{Name: "name", ID: 1, Num: 2.5}}
	large := []Demo{{Name: strings.Repeat("x", 5000), ID: 1, Num: 2.5}}

	tests := []struct {
		name    string
		rows    []Demo
		options CsvOptions
		failAt  int
		short   bool
	}{
		{name: "header only", rows: []Demo{}, failAt: 1},
		{name: "small export", rows: small, failAt: 1},
		{name: "header short write", rows: []Demo{}, failAt: 1, short: true},
		{name: "small export short write", rows: small, failAt: 1, short: true},
		{name: "final flush after successful write", rows: large, failAt: 2},
		{name: "final flush short write", rows: large, failAt: 2, short: true},
		{name: "record write", rows: large, failAt: 1},
		{name: "record short write", rows: large, failAt: 1, short: true},
		{name: "configured export", rows: small, options: CsvOptions{Separator: ';', UseCRLF: true}, failAt: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &dumpFailWriter{failAt: tt.failAt, err: writeErr, short: tt.short}
			wantErr := writeErr
			if tt.short {
				wantErr = io.ErrShortWrite
			}
			if err := DumpToWriter(tt.rows, writer, tt.options); !errors.Is(err, wantErr) {
				t.Fatalf("DumpToWriter() error = %v, want %v", err, wantErr)
			}
			if writer.writes != tt.failAt {
				t.Fatalf("destination writes = %d, want %d", writer.writes, tt.failAt)
			}
		})
	}
}

func TestDumpToWriterConfiguredOutput(t *testing.T) {
	for _, tt := range []struct {
		name    string
		options CsvOptions
		want    string
	}{
		{name: "default", want: "name,ID,number\nname,1,2.5\n"},
		{name: "separator and CRLF", options: CsvOptions{Separator: ';', UseCRLF: true}, want: "name;ID;number\r\nname;1;2.5\r\n"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := DumpToWriter([]Demo{{Name: "name", ID: 1, Num: 2.5}}, &buffer, tt.options); err != nil {
				t.Fatal(err)
			}
			if got := buffer.String(); got != tt.want {
				t.Fatalf("output = %q, want %q", got, tt.want)
			}
		})
	}
}
