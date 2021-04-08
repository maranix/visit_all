package main

import "github.com/ramanverma2k/visit_all/source"

var url_map = map[string]string{
	"moneycontrol": "https://www.moneycontrol.com/",
	"investing": "https://in.investing.com/",
	"yahoo": "https://in.finance.yahoo.com/",
	"indiatimes": "https://economictimes.indiatimes.com/",
	"cnbctv18": "https://www.cnbctv18.com/",
}

func main() {
	keyword := "Corona"
	source.Investing("https://in.investing.com", keyword)
	// for k, v := range url_map {
	// 	if(k=="moneycontrol"){
	// 		source.MoneyControl(v, keyword)
	// 	}
	// }

}