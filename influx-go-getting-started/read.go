package main

import (
	"context"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func ReadData(client influxdb2.Client) {
	org := "influxdata"

	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "default")
				|> range(start: -10m)
				|> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}
