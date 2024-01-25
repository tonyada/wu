package structEx

import (
	"fmt"
	"reflect"
	"testing"
)

type Users struct {
	Id     int
	Name   string
	Age    int
	Market map[int]string
	Source *Sfrom
	Ext    Info
}
type Info struct {
	Detail string
}
type Sfrom struct {
	Area string
}

func (u Users) Login() {
	fmt.Println("login")
}

type StockBasic struct {
	AdaCode   string `json:"-" sql:"-"`                 //! ada系统索引代码 = 格式 SH600298 SZ000333
	StockCode string `json:"stock_code" em_clist:"f12"` // f12 证券代码: 600298 000333
	//= 证券名称和拼音索引 (有可能需要更新，需要检查 每天开盘前更新)
	StockName  string `json:"stock_name" em_clist:"f14"`                                             // f14 证券名称: 安琪酵母
	MarketCode int    `json:"market_code" em_clist:"f13" gorm:"type:tinyint(1);comment:'1=沪市 0=深市'"` //= f13 沪市 = 1  深市 = 0
	MarketStr  string `json:"market_str"`                                                            // 沪市 = SH 深市 = SZ
	// StockPinyin string `json:"stock_pinyin" gorm:"type:varchar(4);"`                                  // 证券中文拼音简单代码 AQJM=安琪酵母 WLY=五粮液
	Status int `json:"status" gorm:"type:tinyint(1);comment:'0=正常交易 1=停牌'"`
	// StockBaseUpdateTime string `json:"stock_base_update_time"` //! 每天 9：10 之后更新一次 "910"
}

var jsonOneData = `{
	"abc":11,
	"data":{
			"f12": "000333", 
			"f14": "美的集团",
			"f3": 0,
			"f4": "SZ"
		}
	}`
var jsonNestedData = `{
	"abc":11,
	"data":[
		{
			"f12": "000333",
			"f14": "美的集团",
			"f3": 0,
			"f4": "SZ"
		},
		{
			"f12": "600298",
			"f14": "安琪酵母",
			"f3": 1,
			"f4": "SH"
		}
]}`

func TestGetFieldNames(t *testing.T) {
	stockA := &StockBasic{
		AdaCode:    "SH600298",
		StockCode:  "600298",
		StockName:  "安琪酵母",
		MarketCode: 1,
		MarketStr:  "SH",
	}
	fmt.Println(stockA)

	fmt.Println("TestGetFieldNames:", GetFieldNames(stockA))
}

func TestGetTagNamesByArrayPos(t *testing.T) {
	stockA := &StockBasic{
		AdaCode:    "SH600298",
		StockCode:  "600298",
		StockName:  "安琪酵母",
		MarketCode: 1,
		MarketStr:  "SH",
	}

	fmt.Println("TestGetTagNamesByArrayPos:", GetTagsByArrayPos(2, stockA))
}
func TestGetFieldTagValue(t *testing.T) {
	stockA := &StockBasic{
		AdaCode:    "SH600298",
		StockCode:  "600298",
		StockName:  "安琪酵母",
		MarketCode: 1,
		MarketStr:  "SH",
	}
	fmt.Println("TestGetFieldTagValue:", GetFieldTagValue("MarketCode", "em_clist", stockA))
}

func TestExplicit(t *testing.T) {
	stockA := &StockBasic{
		AdaCode:    "SH600298",
		StockCode:  "600298",
		StockName:  "安琪酵母",
		MarketCode: 1,
		MarketStr:  "SH",
	}
	m := map[int]string{1: "abc"}
	s := &Sfrom{Area: "beijing"}
	i := Info{Detail: "detail"}
	u := &Users{Id: 12, Market: m, Ext: i, Source: s}
	v := reflect.ValueOf(u)
	Explicit(v, 0)
	Explicit(reflect.ValueOf(stockA), 0)

}

func TestPrint(t *testing.T) {

	m := map[int]string{1: "abc"}
	s := &Sfrom{Area: "beijing"}
	i := Info{Detail: "detail"}
	u := &Users{Id: 12, Market: m, Ext: i, Source: s}
	Print(u)

}
