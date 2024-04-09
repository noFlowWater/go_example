package main

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://155.230.34.51:32145"
	client := influxdb2.NewClient(url, token)

	// 데이터 쓰기
	WriteData(client)

	// 데이터 읽기
	ReadData(client)

	// 클라이언트 닫기
	client.Close()
}
