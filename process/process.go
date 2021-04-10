package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func RemoveDuplicatesUnordered(elements []string, url ...string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        if (url != nil) {
			result = append(result, url[0]+strings.Trim(key, "=\""))
		} else {
			result = append(result, strings.Trim(key, "=\""))
		}
    }
    return result
}

func PageRequest(url *string) string {
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func RegEx(url string, expression ...string) []string {
	var re *regexp.Regexp
	var filtered_links []string
	homepage := PageRequest(&url)

	if (expression != nil) {
		re = regexp.MustCompile(expression[0])
		links_to_visit := re.FindAllString(homepage, -1)
		filtered_links = RemoveDuplicatesUnordered(links_to_visit)
	} else {
		re = regexp.MustCompile(`[^a-z]"/[[:alpha:]].*?"`)
		links_to_visit := re.FindAllString(homepage, -1)
		filtered_links = RemoveDuplicatesUnordered(links_to_visit, url)
	}
	
	
	return filtered_links
}

func SubLinks(s []string) []string {
	s_tmp := s
	for _, l := range s_tmp {
		sublinks := RegEx(l)
		s = append(s, sublinks...)
	}

	return s
}

func VisitLinks(keyword *string,links []string) int {
	var sum int
	for _, l := range links {
		count := FindKeyword(keyword, PageRequest(&l), l)
		sum = sum + count
	}
	return sum
}

func FindKeyword(keyword *string, source string, link string) int {
	// <(?:h1|p).*?>(.*)</(?:h1|p)>
	re := regexp.MustCompile(*keyword)
	match := re.FindAllString(source, -1)
	if match != nil {
		fmt.Println(link, len(match))
	}
	return len(match)
}