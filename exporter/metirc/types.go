package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

)

func NewGauge(nameSpace string, name string, help string) prometheus.Gauge {
	result := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "" + nameSpace,
			Name:      "" + name,
			Help:      "" + help,
		},
	)

	return result
}
