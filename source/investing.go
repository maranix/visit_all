package source

import (
	"fmt"
	"regexp"

	"github.com/ramanverma2k/visit_all/process"
)

// (.*)

func Investing(url string, keyword string) {
	homepage := process.PageRequest(&url)
	re := regexp.MustCompile(`href="/news.*?">`)
	links_to_visit := re.FindAllString(homepage, -1)
	filtered_links := process.RemoveDuplicatesUnordered(links_to_visit)
	for _, v := range filtered_links {
		fmt.Println(v)
	}
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