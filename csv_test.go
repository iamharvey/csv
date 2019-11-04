package csv

import (
	"reflect"
	"testing"
)

func TestCSVReader_Row_IncludeHead(t *testing.T) {
	var reader Reader
	err := reader.Read("sample/weights.csv", ',', false, nil)
	if err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := []string{"date", "weight-in-kg", "time-of-last-meal"}
	actual := reader.Row(1)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestCSVReader_Row_ExcludeHead(t *testing.T) {
	var reader Reader
	err := reader.Read("sample/weights.csv", ',', true, nil)
	if err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := []string{"2019-10-01", "80.5", "2000"}
	actual := reader.Row(1)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestCSVReader_Rows(t *testing.T) {
	var reader Reader
	err := reader.Read("sample/weights.csv", ',', true, nil)
	if err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := [][]string{
		{"2019-10-01", "80.5", "2000"},
		{"2019-10-02", "80", "1930"},
	}

	actual := reader.Rows(1, 3)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestCSVReader_Col_ByIndex(t *testing.T) {
	var reader Reader
	err := reader.Read("sample/weights.csv", ',', true, nil)
	if err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := []string{"80.5", "80", "79.5", "80.4"}

	actual, err := reader.Col(2)
	if err != nil {
		t.Errorf("error occurs %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestCSVReader_Col_ByName(t *testing.T) {
	var reader Reader
	err := reader.Read("sample/weights.csv", ',', true, []string{"date", "weight-in-kg",
		"time-of-last-meal"})
	if err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := []string{"80.5", "80", "79.5", "80.4"}

	actual, err := reader.Col("weight-in-kg")
	if err != nil {
		t.Errorf("error occurs %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestCSVReader_Cols_ByIndices(t *testing.T) {
	var reader Reader
	if err := reader.Read("sample/weights.csv", ',', true, nil); err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := [][]string{
		{"2019-10-01", "2000"},
		{"2019-10-02", "1930"},
		{"2019-10-03", "1730"},
		{"2019-10-04", "2005"},
	}

	if actual, err := reader.Cols(1, 3); err != nil {
		t.Errorf("error occurs %v", err.Error())
	} else {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}
}

func TestCSVReader_Cols_ByNames(t *testing.T) {
	var reader Reader

	if err := reader.Read("sample/weights.csv", ',', true, []string{"date", "weight-in-kg",
		"time-of-last-meal"}); err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := [][]string{
		{"2019-10-01", "2000"},
		{"2019-10-02", "1930"},
		{"2019-10-03", "1730"},
		{"2019-10-04", "2005"},
	}

	if actual, err := reader.Cols("date", "time-of-last-meal"); err != nil {
		t.Errorf("error occurs %v", err)
	} else {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}
}

func TestCSVReader_RowNCol_ByIndices(t *testing.T) {
	var reader Reader

	if err := reader.Read("sample/weights.csv", ',', true, nil); err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := [][]string{
		{"2019-10-02", "1930"},
		{"2019-10-03", "1730"},
	}

	if actual, err := reader.Slice(2, 4, 1, 3); err != nil {
		t.Errorf("error occurs %v", err)
	} else {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}
}

func TestCSVReader_RowNCol_ByNames(t *testing.T) {
	var reader Reader

	if err := reader.Read("sample/weights.csv", ',', true, []string{"date", "weight-in-kg",
		"time-of-last-meal"}); err != nil {
		t.Errorf("error occurs %v", err)
	}

	expected := [][]string{
		{"2019-10-02", "1930"},
		{"2019-10-03", "1730"},
	}

	if actual, err := reader.Slice(2, 4, "date", "time-of-last-meal"); err != nil {
		t.Errorf("error occurs %v", err)
	} else {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}
}
