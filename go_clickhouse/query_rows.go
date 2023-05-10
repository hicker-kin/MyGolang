package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	err := QueryRows()
	if err != nil {
		fmt.Println(err)
	}
}

func QueryRows() error {
	tableName := "example"
	if err := initTable(tableName); err != nil {
		return err
	}
	rows, err := conn.Query(context.Background(),
		"SELECT Col1, Col2, Col3 FROM example WHERE Col1 >= 2")
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			col1 uint8
			col2 string
			col3 time.Time
		)
		if err := rows.Scan(&col1, &col2, &col3); err != nil {
			return err
		}
		fmt.Printf("row: col1=%d, col2=%s, col3=%s\n", col1, col2, col3)
	}
	rows.Close()
	return rows.Err()

	/*
		row: col1=2, col2=value_2, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=3, col2=value_3, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=4, col2=value_4, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=5, col2=value_5, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=6, col2=value_6, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=7, col2=value_7, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=8, col2=value_8, col3=2023-05-10 08:13:42 +0000 UTC
		row: col1=9, col2=value_9, col3=2023-05-10 08:13:42 +0000 UTC
	*/
}
