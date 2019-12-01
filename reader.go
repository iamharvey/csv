package csv

import (
	"encoding/csv"
	"errors"
	"os"
)

/*
Reader defines a reading structure for reading a CSV file.
*/
type Reader struct {
	fileName  string
	data      [][]string
	skip      bool
	separator string
	columns   map[string]int
}

/*
Read reads a CSV file by specifying:
- name: file name (full path)
- sep: separator, e.g., ',', ';'
- skip: if true first row will be skipped, otherwise not skipped
- columns: column names, use can set column names and later use column name to retrieve column data.
*/
func (reader *Reader) Read(fileName string, sep rune, skip bool, columns []string) error {

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// create a csv reader
	r := csv.NewReader(f)
	r.Comma = sep
	r.Comment = '#'

	// read everything, the result is a 2d string array
	data, err := r.ReadAll()
	if err != nil {
		return err
	}

	// check if the header (first row) should be skipped
	from := 0
	if skip {
		from = 1
	}
	reader.data = data[from:]
	reader.skip = skip

	// setup column names for later use
	colN := len(columns)
	reader.columns = make(map[string]int, colN)
	if colN > 0 {
		for i := 1; i <= colN; i++ {
			reader.columns[columns[i-1]] = i
		}
	}

	return nil
}

func (reader *Reader) Data() [][] string {
	return reader.data
}

/*
Row retrieves a single row using row number (starts from 1).
*/
func (reader *Reader) Row(index int) []string {
	return reader.data[index-1 : index-1+1][0]
}

/*
Rows retrieves rows given a range(from, to). It returns a 2d array even user only retrieves one row.
*/
func (reader *Reader) Rows(from, to int) [][]string {
	return reader.data[from-1 : to-1]
}

/*
Col retrieves a single column data given either column index or column name.
*/
func (reader *Reader) Col(colName interface{}) (data []string, err error) {
	for _, row := range reader.data {
		// check if colName is potentially an index number
		if index, ok := colName.(int); ok {
			data = append(data, row[index-1])
		} else {
			// check if colName is potentially a column name
			if name, ok := colName.(string); ok {
				data = append(data, row[reader.columns[name]-1])
			} else {
				return nil, errors.New("column name is not found")
			}
		}

	}
	return data, nil
}

/*
Cols retrieves a multi-column data given either column indices or column names.
It returns a 2d array even user only retrieves one column.
*/
func (reader *Reader) Cols(colNames ...interface{}) (data [][]string, err error) {
	for _, row := range reader.data {
		var line []string
		for _, colName := range colNames {
			// check if colName is potentially an index number
			if index, ok := colName.(int); ok {
				line = append(line, row[index-1])
			} else {
				// check if colName is potentially a column name
				if name, ok := colName.(string); ok {
					line = append(line, row[reader.columns[name]-1])
				} else {
					return nil, errors.New("column name is not found")
				}
			}
		}
		data = append(data, line)
	}
	return data, nil
}

/*
Slice retrieves a multi-row, multi-column data given a row range and columns (either indices or names).
It returns a 2d array even user only retrieves one cell.
*/
func (reader *Reader) Slice(from, to int, colNames ...interface{}) (data [][]string, err error) {
	for _, row := range reader.data[from-1 : to-1] {
		var line []string
		for _, colName := range colNames {
			// check if colName is potentially an index number
			if index, ok := colName.(int); ok {
				line = append(line, row[index-1])
			} else {
				// check if colName is potentially a column name
				if name, ok := colName.(string); ok {
					line = append(line, row[reader.columns[name]-1])
				} else {
					return nil, errors.New("column name is not found")
				}
			}
		}
		data = append(data, line)
	}
	return data, nil
}
