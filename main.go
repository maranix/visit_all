package main

import "github.com/ramanverma2k/visit_all/source"

/*
*
* TODO: Introduce Channels and routines
*
 */

func main() {
	keyword := "Plastiblends"
	source.YahooFinance("https://in.finance.yahoo.com", keyword)
	// source.EconomicTimes("https://economictimes.indiatimes.com", keyword)
	// // source.CNBC("https://www.cnbctv18.com", keyword, `\s*(?i)https://www[.]cnbctv18[.]com(\"([^"]*\")|'[^']*'|([^'">\s]+))`)
	// source.MoneyControl("https://www.moneycontrol.com", keyword, `\s*(?i)https://www[.]moneycontrol[.]com(\"([^"]*\")|'[^']*'|([^'">\s]+))`)
	// source.Investing("https://in.investing.com", keyword)
}