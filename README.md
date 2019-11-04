![](https://github.com/iamharvey/csv/workflows/Go/badge.svg) ![Go report](https://goreportcard.com/badge/github.com/iamharvey/csv)

# csv - A package reads and slices data from a CSV file

## Features
- reads a csv with or without skipping header
- retrieves a single row or rows
- retrieves a single column or cols
- retrieves a slice of data by specifiying both the row range and column indices

## Example

### Read a CSV file, header skipped
```
err := reader.Read("sample/weights.csv", true)
if err != nil {
  fmt.Println(err)
}
```

### Retrieves the first two rows
```
err := reader.Read("sample/weights.csv", true)
if err != nil {
  fmt.Println(err)
}
data := reader.Rows(1, 3)
```

### Retrieves the data of the first and the third column
```
err := reader.Read("sample/weights.csv", true)
if err != nil {
  fmt.Println(err)
}
data := reader.Cols(1, 3)
```

We also support retrieving columns by their names, E.g.,
```
err := reader.Read("sample/weights.csv", ',', true, []string{"date","weight-in-kg",
		"time-of-last-meal"})
if err != nil {
  fmt.Println(err)
}
data := reader.Cols("date", "time-of-last-meal")
```

### Retrieve the rows (2, 3, 4) of the data of the first and the third column
```
err := reader.Read("sample/weights.csv", true)
if err != nil {
  fmt.Println(err)
}
data := reader.Slice(2, 4, 1, 3)
```

In this case, we also support retrieving columns by their names. E.g., 
```
err := reader.Read("sample/weights.csv", ',', true, []string{"date","weight-in-kg",
       		"time-of-last-meal"})
if err != nil {
  fmt.Println(err)
}
data := reader.Slice(2, 4, "date", "time-of-last-meal")
```

## Benchmark
Task | Rounds | Speed | Mem Alloc | Allocs 
  --- | --- | --- | --- | --- 
| Get column data (random cols) |  980545  |            3710 ns/op   |         2672 B/op   |      54 allocs/op
