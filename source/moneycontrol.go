package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func MoneyControl(url string, keyword string, expression string) {
	links := process.RegEx(url, expression)
	sub_links := process.SubLinks(links)
	fmt.Println("MoneyControl: ")
	fmt.Println("Parent links: ", len(links))
	fmt.Println("SubLinks: ", len(sub_links))
	fmt.Println(process.VisitLinks(&keyword, sub_links))
}