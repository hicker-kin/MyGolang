package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// https://github.com/ClickHouse/clickhouse-go/blob/main/examples/clickhouse_api/append_struct.go

func TestStruct(t *testing.T) {
	err := initTable2()
	if err != nil {
		t.Fatal(err)
	}

	rows, err := conn.Query(context.Background(),
		"SELECT Col1, Col2, Col3,Col4 FROM example WHERE Col1 >= 2")
	if err != nil {
		t.Fatal(err)
	}

	var list []*StructRow
	for rows.Next() {
		data := &StructRow{}
		if err := rows.ScanStruct(data); err != nil {
			t.Fatal(err)
		}
		list = append(list, data)
		// fmt.Printf("row: col1=%d, col2=%s, col3=%s, col4=%s\n", data.Col1, data.Col2, data.Col3, data.Col4)
		fmt.Printf("row: col1=%d, col2=%s, col3=%v, col4=%s\n", data.Col1, data.Col2, data.Col3, data.Col4)
	}

	/*
		row: col1=2, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=3, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=4, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=5, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=6, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=7, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=8, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
		row: col1=9, col2=Golang SQL database driver, col3=[1 2 3 4 5 6 7 8 9], col4=2023-05-10 08:47:42 +0000 UTC
	*/
	t.Log("size ==== ", len(list))
	rows.Close()
	t.Log(rows.Err())
}

type StructRow struct {
	Col1       uint64
	Col4       time.Time
	Col2       string
	Col3       []uint8
	ColIgnored string
}

func initTable2() error {
	ctx := context.Background()
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		return err
	}
	if err := conn.Exec(ctx, `
		CREATE TABLE example (
			  Col1 UInt64
			, Col2 String
			, Col3 Array(UInt8)
			, Col4 DateTime
		) Engine = Memory
		`); err != nil {
		return err
	}

	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO example")
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		err := batch.AppendStruct(&StructRow{
			Col1:       uint64(i),
			Col2:       "Golang SQL database driver",
			Col3:       []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Col4:       time.Now(),
			ColIgnored: "this will be ignored",
		})
		if err != nil {
			return err
		}
	}
	return batch.Send()
}
