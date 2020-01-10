package metric

import (
	"go.uber.org/zap"

//	rpc "github.com/node-a-team/iov-validator_exporter/getData/rpc"
)

var (
	metricData metric

	GaugesNamespaceList = [...]string{"atom_coinone",
//				"commitVoteType",
//				"precommitStatus",
				}
)

type metric struct {

	ATOM struct {
		Coinone         float64
	}
}



func SetMetric(log *zap.Logger) {

	//// network
	metricData.ATOM.Coinone = 5000.0 //getATOM_fromCoinone()
}

func GetMetric() *metric {

	return &metricData
}
