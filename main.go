package main

import (
	"fmt"
	"os"
	"net/http"
	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	exporter "github.com/node-a-team/price_exporter/exporter"
	cfg "github.com/node-a-team/price_exporter/config"
)

var (
	port string = "26661"
)

func main() {

	log,_ := zap.NewDevelopment()
        defer log.Sync()

	cfg.ConfigPath = os.Args[1]
	cfg.Init(log)

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start(log)

	err := http.ListenAndServe(":" +port, nil)

	// log
        if err != nil {
                // handle error
                log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: " +port +")"),)
        }
}
