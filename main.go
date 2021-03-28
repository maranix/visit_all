package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)


func main() {
	main_url := "https://www.moneycontrol.com/"
	
	homepage := page_request(main_url)

	re := regexp.MustCompile(`\s*(?i)https://www[.]moneycontrol[.]com(\"([^"]*\")|'[^']*'|([^'">\s]+))`)
	links_to_visit := re.FindAllString(homepage, -1)
	visit_links(links_to_visit)
}

func page_request(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func visit_links(links []string) {
	for _, l := range links {
		page := page_request(l)
	}
}

func count_match() int{


	return sum
}