package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func MoneyControl(url string, keyword string) {
	links := process.RegEx(url)
	fmt.Println(process.VisitLinks(&keyword, links))
}