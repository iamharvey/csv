package csv

import (
	"encoding/csv"
	"os"
)

/*
Writer is for writing a CSV file.
*/
type Writer struct {}

/*
Write writes a [][]string to a csv file
 */
func (writer *Writer) Write(data [][]string, saveAs string) error {
	f, err := os.Create( saveAs)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	err = w.WriteAll(data)
	if err != nil {
		return err
	}

	return f.Close()
}
