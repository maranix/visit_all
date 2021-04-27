package source

import (
	"fmt"
	"time"

	"github.com/ramanverma2k/visit_all/process"
)

func EconomicTimes(url string, keyword string) {
	var channel = make(chan []string)
	fmt.Println("Yahoo Finance: ")
	t0 := time.Now()
	go process.RegEx(url, channel)
	links := <-channel
	fmt.Println("Parent links: ", len(links))
	sub_links := process.SubLinks(links, url)
	fmt.Println("SubLinks: ", len(sub_links))
	fmt.Printf("It took %v to run\n", time.Since(t0))
	// fmt.Println(process.VisitLinks(&keyword, sub_links))
	// fmt.Println(process.VisitLinks(&keyword, links))
}