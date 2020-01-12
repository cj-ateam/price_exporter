package exporter

import (
	"fmt"
	"time"
	"go.uber.org/zap"

	price "github.com/node-a-team/price_exporter/price"
	metric "github.com/node-a-team/price_exporter/exporter/metric"

	"github.com/prometheus/client_golang/prometheus"
)

var (

)

func Start(log *zap.Logger) {

	gaugesNamespaceList := metric.GaugesNamespaceList

	var gauges []prometheus.Gauge = make([]prometheus.Gauge, len(gaugesNamespaceList))


	// nomal guages
	for i := 0; i < len(gaugesNamespaceList); i++ {
                gauges[i] = metric.NewGauge("price", gaugesNamespaceList[i], "")
                prometheus.MustRegister(gauges[i])
        }


	ps := price.NewPriceService()
	ps.OnStart(log)



	time.Sleep(1 * time.Second)

	for {
		func() {
/*
			defer func() {

				if r := recover(); r != nil {
					//Error Log
				}

				time.Sleep(500 * time.Millisecond)

			}()
*/
				metric.SetMetric(log, ps)
				metricData := metric.GetMetric()

				fmt.Println("usd/krw/dunamu: ", metricData.USD.KRW.Dunamu)

				fmt.Println("btc/krw/upbit: ", metricData.BTC.KRW.Upbit)
				fmt.Println("btc/usdt/upbit: ", metricData.BTC.USDT.Upbit)
				fmt.Println("btc/usdt/binance: ", metricData.BTC.USDT.Binance)
				fmt.Println("btc/usdt/huobiGlobal: ", metricData.BTC.USDT.HuobiGlobal)

				fmt.Println("\natom/krw/coinone: ", metricData.ATOM.KRW.Coinone)
				fmt.Println("atom/krw/upbit: ", metricData.ATOM.KRW.Upbit)
				fmt.Println("atom/btc/upbit: ", metricData.ATOM.BTC.Upbit)
				fmt.Println("atom/btc/binance: ", metricData.ATOM.BTC.Binance)
				fmt.Println("atom/btc/huobiGlobal: ", metricData.ATOM.BTC.HuobiGlobal)
				fmt.Println("atom/usdt/huobiGlobal: ", metricData.ATOM.USDT.HuobiGlobal)
				fmt.Println("atom/usdt/binance: ", metricData.ATOM.USDT.Binance)

				fmt.Println("\nluna/krw/coinone: ", metricData.LUNA.KRW.Coinone)
				fmt.Println("luna/krw/bithumb: ", metricData.LUNA.KRW.Bithumb)
				fmt.Println("luna/btc/upbit: ", metricData.LUNA.BTC.Upbit)

				fmt.Println("\niris/btc/huobiGlobal: ", metricData.IRIS.BTC.HuobiGlobal)
                                fmt.Println("iris/usdt/huobiGlobal: ", metricData.IRIS.USDT.HuobiGlobal)

				fmt.Println("\nkava/btc/binance: ", metricData.KAVA.BTC.Binance)
				fmt.Println("kava/usdt/binance: ", metricData.KAVA.USDT.Binance)

				fmt.Println("------")

				gaugesValue := [...]float64{

					metricData.USD.KRW.Dunamu,

					metricData.BTC.KRW.Upbit,
					metricData.BTC.USDT.Upbit,
					metricData.BTC.USDT.Binance,
					metricData.BTC.USDT.HuobiGlobal,

					metricData.ATOM.KRW.Coinone,
					metricData.ATOM.KRW.Upbit,
					metricData.ATOM.USDT.Binance,
					metricData.ATOM.USDT.HuobiGlobal,
					metricData.ATOM.BTC.Binance,
					metricData.ATOM.BTC.HuobiGlobal,

					metricData.LUNA.KRW.Coinone,
					metricData.LUNA.KRW.Bithumb,
					metricData.LUNA.BTC.Upbit,

					metricData.IRIS.USDT.HuobiGlobal,
					metricData.IRIS.BTC.HuobiGlobal,

					metricData.KAVA.USDT.Binance,
					metricData.KAVA.BTC.Binance,

				}

				for i := 0; i < len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
				}

			fmt.Println(" ")
			time.Sleep(1000 * time.Millisecond)
		}()
	}
}
