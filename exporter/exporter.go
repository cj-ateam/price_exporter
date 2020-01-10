package exporter

import (
	"fmt"
	"time"
	"go.uber.org/zap"

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



				metric.SetMetric(currentBlockHeight, rpcData, log)
				metricData := metric.GetMetric()


				gaugesValue := [...]float64{
					float64(metricData.Network.BlockHeight),

					metricData.Validator.Commit.VoteType,
					metricData.Validator.Commit.PrecommitStatus,
				}

				for i := 0; i < len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
				}

			time.Sleep(500 * time.Millisecond)
		}()
	}
}
