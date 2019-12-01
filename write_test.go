package csv

import (
	"reflect"
	"testing"
	)

func TestWriter_Write(t *testing.T) {
	expected := [][]string{
		{"Jack", "Ryan"},
		{"Tommy", "Mendez"},
	}
	var w Writer
	err := w.Write(expected, "./sample/write_test.csv")

	if err != nil {
		t.Errorf("error unexpected, %s\n", err.Error())
	}

	var r Reader
	r.Read("./sample/write_test.csv", ',', false, nil)
	actual := r.Data()
	if !reflect.DeepEqual(expected, actual) {
		t.Error("the result does not meet the expected on\n")
	}

}
