package main

import (
	"log"
	"strings"
)

// 外匯商品瀏覽功能，可查詢各種貨幣對，CcyPair由BaseCcy、TermCcy組成
// 另開發了投資組合功能，可管理股票部位
// 股票的全名為由股票代碼、股票簡稱組成
// 股票部位的Key為股票代碼
// 接著，想要讓投資組合功能也能用來管理外匯部位

type IForex interface {
	CcyPair() string
}
type Forex struct {
	BaseCcy string
	TerCcy  string
}

func (fx *Forex) CcyPair() string {
	return fx.BaseCcy + fx.TerCcy
}

type IStock interface {
	Ticker() string
}

// Ticker TickerName = "2330-台積電"
func (s *Stock) Ticker() string {
	return strings.Split(s.TickerName, "-")[0]
}

type Stock struct {
	TickerName string
	IStock
}

type FXAdapter struct {
	IForex
	IStock
}

func (fxa *FXAdapter) Ticker() string {
	return fxa.IForex.CcyPair()
}

func main() {
	f := Forex{
		BaseCcy: "USD",
		TerCcy:  "TWD",
	}

	// st := Stock{
	// 	TickerName: "2330-台積電",
	// }

	st := FXAdapter{
		IForex: &f,
	}

	// log.Println(f.CcyPair())
	log.Println(st.Ticker())
}
