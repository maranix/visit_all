package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func CNBC(url string, keyword string, expression string) {
	links := process.RegEx(url, expression)
	fmt.Println("CNBC: ")
	fmt.Println("Parent links: ", len(links))
	sub_links := process.SubLinks(links)
	fmt.Println("SubLinks: ", len(sub_links))
}