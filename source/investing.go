package source

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ramanverma2k/visit_all/process"
)

func Investing(url string, keyword *string) {
	homepage := process.PageRequest(&url)
	re := regexp.MustCompile(`[^a-z]"/[[:alpha:]].*?"`)
	sub_urls := re.FindAllString(homepage, -1)
	for i, u := range sub_urls {
		sub_urls[i] = strings.Trim(u, "=\"")
	}
	filtered_links := process.RemoveDuplicatesUnordered(sub_urls)
	
	for i, v := range filtered_links {
		filtered_links[i] = "https://in.investing.com"+v
	}

	fmt.Println(process.VisitLinks(keyword, filtered_links))
	// file, err := os.OpenFile("links.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// writer := bufio.NewWriter(file)

	// for _, data := range filtered_links {
	// 	_, _ = writer.WriteString(data + "\n")
	// }

	// writer.Flush()
	// file.Close()
	// fmt.Println(process.VisitLinks(&keyword, filtered_links))
}