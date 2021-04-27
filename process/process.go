package process

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func RemoveDuplicatesUnordered(elements []string, url string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key := range encountered {
        if (url != "") {
			result = append(result, url+strings.Trim(key, "=\""))
		}
    }
    return result
}

func PageRequest(url *string) string {
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func RegEx(url string, expression ...string) []string {
	var re *regexp.Regexp
	var links_to_visit []string
	homepage := PageRequest(&url)

	if (expression == nil) {
		re = regexp.MustCompile(`[^a-z]"/[[:alpha:]].*?"`)
		links_to_visit = re.FindAllString(homepage, -1)
	} else {
		re = regexp.MustCompile(expression[0])
		links_to_visit = re.FindAllString(homepage, -1)
	}
	filtered_links := RemoveDuplicatesUnordered(links_to_visit, url)
	return filtered_links
}

func SubLinks(s []string, url string) []string {
	var filtered_links []string
	s_tmp := s
	fmt.Println("Before: ", len(s))
	for _, l := range s_tmp {
		temp := RegEx(l)
		s = append(s, temp...)
	}
	
	s_tmp = s
	for _, l := range s_tmp {
		temp := RegEx(l)
		s = append(s, temp...)
	}
	fmt.Println("s_tmp: ",len(s_tmp))
	fmt.Println("s: ",len(s))
	filtered_links = RemoveDuplicatesUnordered(s, url)
	fmt.Println("After: ",len(filtered_links))
	return filtered_links
}

func VisitLinks(keyword *string,links []string) int {
	channel := make(chan int)
	var sum int
	for _, l := range links {
		go FindKeyword(keyword, PageRequest(&l), l, channel)
		sum = sum + <-channel
	}
	return sum
}

func FindKeyword(keyword *string, source string, link string, ch chan<-int) {
	// <(?:h1|h2|h3|p).*?>.*(Plastiblends).*</(?:h1|h2|h3|p)>
	re := regexp.MustCompile(`<(?:h1|h2|h3|p).*?>.*(`+*keyword+`).*</(?:h1|h2|h3|p)>`)
	match := re.FindAllString(source, -1)
	if match != nil {
		fmt.Println(link, len(match))
	}
	ch <- len(match)
}