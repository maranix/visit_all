package source

import (
	"fmt"

	"github.com/ramanverma2k/visit_all/process"
)

func Investing(url string, keyword string) {
	
	links := process.RegEx(url)
	// fmt.Println(process.VisitLinks(&keyword, links))
	// sub_links := process.SubLinks(links)

	fmt.Println(len(links))

	// fmt.Println(process.VisitLinks(keyword, filtered_links))
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