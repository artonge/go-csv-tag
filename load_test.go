package csvtag

import "testing"

type test struct {
	Name string  `csv:"header1"`
	ID   int     `csv:"header2"`
	Num  float64 `csv:"header3"`
}

type testNoID struct {
	Name string `csv:"header1"`
	ID   int
	Num  float64 `csv:"header"`
}

type testBool struct {
	Name  string `csv:"header1"`
	ABool bool   `csv:"header2"`
}

// Check the values are correct
func checkValues(tabT []test) bool {
	return false ||
		tabT[0].Name != "line1" || tabT[0].ID != 1 || tabT[0].Num != 1.2 ||
		tabT[1].Name != "line2" || tabT[1].ID != 2 || tabT[1].Num != 2.3 ||
		tabT[2].Name != "line3" || tabT[2].ID != 3 || tabT[2].Num != 3.4
}

func checkBoolValues(tabT []testBool) bool {
	return false ||
		tabT[0].Name != "line1" || tabT[0].ABool != true ||
		tabT[1].Name != "line2" || tabT[1].ABool != false ||
		tabT[2].Name != "line3" || tabT[2].ABool != false
}

func TestValideFile(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/valid.csv", &tabT)
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestBool(t *testing.T) {
	tabT := []testBool{}
	err := LoadFromPath("csv_files/bool.csv", &tabT)
	if err != nil || checkBoolValues(tabT) {
		t.Fail()
	}
}

func TestBadHeader(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/badHeader.csv", &tabT)
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestMissATag(t *testing.T) {
	tabT := []testNoID{}
	err := LoadFromPath("csv_files/valid.csv", &tabT)
	if err != nil ||
		tabT[0].Name != "line1" || tabT[0].ID != 0 || tabT[0].Num != 0 ||
		tabT[1].Name != "line2" || tabT[1].ID != 0 || tabT[1].Num != 0 ||
		tabT[2].Name != "line3" || tabT[2].ID != 0 || tabT[2].Num != 0 {
		t.Fail()
	}
}

func TestEmptyFile(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/empty.csv", &tabT)
	if err != nil || len(tabT) != 0 {
		t.Fail()
	}
}

func TestNoHeader(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath(
		"csv_files/noHeader.csv",
		&tabT,
		CsvOptions{Header: []string{"header1", "header2", "header3"}})
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestWithSemicolon(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath(
		"csv_files/semicolon.csv",
		&tabT,
		CsvOptions{Separator: ';'})
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestWithTrailingSpaces(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/trailingSpaces.csv", &tabT)
	if err != nil || checkValues(tabT) {
		t.Fail()
	}
}

func TestToMutchOptions(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath(
		"csv_files/semicolon.csv",
		&tabT,
		CsvOptions{Separator: ';'},
		CsvOptions{Separator: ';'})
	if err == nil {
		t.Fail()
	}
}

func TestBadFormat(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/badFormat.csv", &tabT)
	if err == nil {
		t.Fail()
	}
}

func TestNonexistingFile(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/nonexistingfile.csv", &tabT)
	if err == nil {
		t.Fail()
	}
}

func TestBadInt(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/badInt.csv", &tabT)
	if err == nil {
		t.Fail()
	}
}

func TestBadFloat(t *testing.T) {
	tabT := []test{}
	err := LoadFromPath("csv_files/badFloat.csv", &tabT)
	if err == nil {
		t.Fail()
	}
}

func TestNoDist(t *testing.T) {
	err := LoadFromPath("csv_files/valid.csv", &test{})
	if err == nil {
		t.Fail()
	}
}

func TestLazyQuotes(t *testing.T) {

	type testQuote struct {
		Name  string `csv:"name"`
		ABool bool   `csv:"flag"`
	}

	type quotes []testQuote

	data := quotes{}
	err := LoadFromPath("csv_files/lazy_quotes.csv", &data,
		CsvOptions{Separator: '|',
			Header: []string{
				"name",
				"flag",
			},
			LazyQuote: true})
	if err != nil {
		t.Errorf("read failed with an error: %s", err.Error())
	}
	if len(data) != 1 {
		t.Error("Expected one row read")
	}
	if data[0].Name != "a content with \" \" chars" {
		t.Error("Received unexpected file content")
	}
}
