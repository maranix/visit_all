package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func EconomicTimes(url string, keyword string) {
	links := process.RegEx(url)
	sub_links := process.SubLinks(links)
	fmt.Println(len(links))
	fmt.Println(len(sub_links))
	// for _, l := range sub_links {
	// 	fmt.Println(l)
	// }
}