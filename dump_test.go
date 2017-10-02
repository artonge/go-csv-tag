package csvtag

import (
	"bytes"
	"strings"
	"testing"
)

const (
	TESTSTRUCT     = "header1,header2,header3\nline1,1,1.2\nline2,2,2.3\nline3,3,3.4"
	TESTNOIDSTRUCT = "header1,header\nline1,0\nline2,0\nline3,0"
)

func TestDumpTestStruct(t *testing.T) {
	tabT := []test{}
	err := Load(Config{
		Path: "csv_files/valid.csv",
		Dest: &tabT,
	})
	if err != nil {
		t.Fail()
	}

	b := &bytes.Buffer{}

	err = Dump(tabT, b)
	if err != nil {
		t.Fail()
	}

	out := b.String()
	if strings.Contains(out, TESTSTRUCT) != true {
		t.Fail()
	}
}

func TestDumpTestNoIdStruct(t *testing.T) {
	tabT := []testNoID{}
	err := Load(Config{
		Path: "csv_files/valid.csv",
		Dest: &tabT,
	})
	if err != nil {
		t.Fail()
	}

	b := &bytes.Buffer{}

	err = Dump(tabT, b)
	if err != nil {
		t.Fail()
	}

	out := b.String()

	if strings.Contains(out, TESTNOIDSTRUCT) != true {
		t.Fail()
	}
}

func TestDumpBadInput(t *testing.T) {
	badInput := "Bad Input"
	b := &bytes.Buffer{}

	err := Dump(badInput, b)

	if strings.Compare(err.Error(), UNSUPPORTED) != 0 {
		t.Errorf("Was able to pass bad input into dump")
	}
}
