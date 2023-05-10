package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	log "github.com/sirupsen/logrus"
)

/*
server run:

docker run -d -e CLICKHOUSE_DB=my_database -e CLICKHOUSE_USER=username \
-e CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=1 \
-e CLICKHOUSE_PASSWORD=password \
-p 9000:9000/tcp clickhouse/clickhouse-server
*/

func initTable(tableName string) error {
	ctx := context.Background()
	if err := conn.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)); err != nil {
		return err
	}
	err := conn.Exec(ctx, fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (Col1 UInt8,Col2 String,Col3 DateTime) engine=Memory",
		tableName))
	if err != nil {
		return err
	}

	batch, err := conn.PrepareBatch(ctx, fmt.Sprintf(
		"INSERT INTO %s (Col1, Col2, Col3)", tableName))
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		if err := batch.Append(uint8(i), fmt.Sprintf("value_%d", i), time.Now()); err != nil {
			return err
		}
	}
	return batch.Send()
}

var conn driver.Conn

func init() {
	var err error
	conn, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{"172.16.42.112:9000"}, // 使用容器启动，映射服务端口到主机，此处的地址不能用127,该特性不同于mysql
		Auth: clickhouse.Auth{
			Database: "my_database",
			Username: "username",
			Password: "password",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		//Debug: true,
	})
	if err != nil {
		log.Error(err)
		return
	}
}
