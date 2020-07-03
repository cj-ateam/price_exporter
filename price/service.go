package price

import (
	"sync"
	"go.uber.org/zap"
//	"fmt"
)

type PriceService struct {
	mutex  *sync.RWMutex
	prices map[string]float64
}

func NewPriceService() *PriceService {
	ps := &PriceService{
		mutex:  new(sync.RWMutex),
		prices: make(map[string]float64),
	}
	return ps
}

func (ps *PriceService) OnStart(log *zap.Logger) error {
	// TODO: gracefully quit go routine
	go ps.Dunamu(log)

	go ps.Upbit(log, "krw", "btc")
	go ps.Upbit(log, "usdt", "btc")
	go ps.Binance(log, "usdt", "btc")
	go ps.HuobiGlobal(log, "usdt", "btc")

	go ps.Coinone(log, "atom")
	go ps.Upbit(log, "krw", "atom")
	go ps.Upbit(log, "btc", "atom")
	go ps.Binance(log, "btc", "atom")
	go ps.HuobiGlobal(log, "btc", "atom")
	go ps.HuobiGlobal(log, "usdt", "atom")
	go ps.Binance(log, "usdt", "atom")

	go ps.Coinone(log, "luna")
	go ps.Bithumb(log, "luna")
	go ps.Upbit(log, "btc", "luna")

	go ps.HuobiGlobal(log, "btc", "iris")
	go ps.HuobiGlobal(log, "usdt", "iris")

	go ps.Binance(log, "btc", "kava")
	go ps.Binance(log, "usdt", "kava")
	go ps.Coinone(log, "kava")

	go ps.Binance(log, "usdt", "band")

	go ps.Binance(log, "btc", "sol")
        go ps.Binance(log, "busd", "sol")

	return nil
}

func (ps *PriceService) GetPrice(market string) float64 {
	ps.mutex.RLock()
	defer func() {
		ps.mutex.RUnlock()
	}()
	return ps.prices[market]
}

func (ps *PriceService) SetPrice(market string, price float64) {
	ps.mutex.Lock()
	defer func() {
		ps.mutex.Unlock()
	}()

	ps.prices[market] = price
}
