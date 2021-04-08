package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func MoneyControl(url string, keyword string) {
	links := process.RegEx(url)
	sub_links := process.SubLinks(links)
	fmt.Println(len(links))
	fmt.Println(len(sub_links))
	// fmt.Println(process.VisitLinks(&keyword, links))
}