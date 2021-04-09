package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func Investing(url string, keyword string) {
	
	links := process.RegEx(url)
	sub_links := process.SubLinks(links)
	fmt.Println("Investing: ")
	fmt.Println("Parent links: ", len(links))
	fmt.Println("SubLinks: ", len(sub_links))
	fmt.Println(process.VisitLinks(&keyword, sub_links))
}