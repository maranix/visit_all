package source

import (
	"fmt"
	"regexp"

	"github.com/ramanverma2k/visit_all/process"
)

func MoneyControl(url string, keyword string) {
	homepage := process.PageRequest(&url)
	re := regexp.MustCompile(`\s*(?i)https://www[.]moneycontrol[.]com(\"([^"]*\")|'[^']*'|([^'">\s]+))`)
	links_to_visit := re.FindAllString(homepage, -1)
	filtered_links := process.RemoveDuplicatesUnordered(links_to_visit)
	fmt.Println(process.VisitLinks(&keyword, filtered_links))
}