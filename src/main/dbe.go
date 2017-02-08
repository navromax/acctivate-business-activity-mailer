package main

import (
   _ "github.com/denisenkom/go-mssqldb"
  "database/sql"
  "fmt"
)

func indexOf(a string, list []string) int {
  for i, b := range list {
    if b == a {
      return i
    }
  }
  return -1
}

func GetColumn(rows *sql.Rows, columnName string, dest interface{}) error {
  if cols, err := rows.Columns(); err != nil {
    return err
  } else {
    index := indexOf(columnName, cols)
    if index < 0 {
      return fmt.Errorf("Column %s not found", columnName)
    }

    pointers := make([]interface{}, len(cols))
    for i, _ := range pointers {
      pointers[i] = new(interface{})
    }
    pointers[index] = dest

    return rows.Scan(pointers...)
  }
}