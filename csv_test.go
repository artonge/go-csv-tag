package csvtag

import "testing"

type Test struct {
	Name string  `csv:"header1"`
	ID   int     `csv:"header2"`
	Num  float64 `csv:"header3"`
}

type TestNoID struct {
	Name string `csv:"header1"`
	ID   int
	Num  float64 `csv:"header"`
}

// Check the values are correct
func checkValues(tabT []Test) bool {
	return false ||
		tabT[0].Name != "line1" || tabT[0].ID != 1 || tabT[0].Num != 1.2 ||
		tabT[1].Name != "line2" || tabT[1].ID != 2 || tabT[1].Num != 2.3 ||
		tabT[2].Name != "line3" || tabT[2].ID != 3 || tabT[2].Num != 3.4
}

func TestLoad_valideFile(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path: "csv_files/valid.csv",
		dest: &tabT,
	})
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestLoad_missATag(t *testing.T) {
	tabT := []TestNoID{}
	err := Load(Config{
		path: "csv_files/valid.csv",
		dest: &tabT,
	})
	if err != nil ||
		tabT[0].Name != "line1" || tabT[0].ID != 0 || tabT[0].Num != 0 ||
		tabT[1].Name != "line2" || tabT[1].ID != 0 || tabT[1].Num != 0 ||
		tabT[2].Name != "line3" || tabT[2].ID != 0 || tabT[2].Num != 0 {
		t.Fail()
	}
}

func TestLoad_emptyFile(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path: "csv_files/empty.csv",
		dest: &tabT,
	})
	if err != nil || len(tabT) != 0 {
		t.Fail()
	}
}

func TestLoad_noHeader(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path:   "csv_files/noHeader.csv",
		dest:   &tabT,
		header: []string{"header1", "header2", "header3"},
	})
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestLoad_withSemicolon(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path:      "csv_files/semicolon.csv",
		dest:      &tabT,
		separator: ';',
	})
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestLoad_badFormat(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path: "csv_files/badFormat.csv",
		dest: &tabT,
	})
	if err == nil {
		t.Fail()
	}
}

func TestLoad_nonexistingFile(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path: "csv_files/nonexistingfile.csv",
		dest: &tabT,
	})
	if err == nil {
		t.Fail()
	}
}

func TestLoad_badInt(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path: "csv_files/badInt.csv",
		dest: &tabT,
	})
	if err == nil {
		t.Fail()
	}
}

func TestLoad_badFloat(t *testing.T) {
	tabT := []Test{}
	err := Load(Config{
		path: "csv_files/badFloat.csv",
		dest: &tabT,
	})
	if err == nil {
		t.Fail()
	}
}

func TestLoad_notPath(t *testing.T) {
	err := Load(Config{
		dest: &[]Test{},
	})
	if err == nil {
		t.Fail()
	}
}

func TestLoad_noDest(t *testing.T) {
	err := Load(Config{
		path: "csv_files/valid.csv",
	})
	if err == nil {
		t.Fail()
	}
}

func TestLoad_noDist(t *testing.T) {
	err := Load(Config{
		path: "csv_files/valid.csv",
		dest: &Test{},
	})
	if err == nil {
		t.Fail()
	}
}
