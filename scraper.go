package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

func ExampleScrape() {
	docv4 := getDoc("https://www.countryipblocks.net/acl.php")
	ipv4 := collectCIDRs(docv4)
	docv6 := getDoc("https://www.countryipblocks.net/ipv6_acl.php")
	ipv6 := collectCIDRs(docv6)
	save(ipv4, "ipv4")
	save(ipv6, "ipv6")
}

func collectCIDRs(doc *goquery.Document) string {
	text := doc.Find("#textareaAll").Text()
	cidrs := strings.Split(text, "\n")
	for i, cidr := range cidrs {
		cidrs[i] = strings.TrimSpace(cidr)
	}
	return strings.Join(cidrs, "\n")
}

func getDoc(url string) *goquery.Document {
	var body url2.Values = map[string][]string{
		"countries[]": {"RU"},
		"format1":     {"1"},
		"get_acl":     {"Create ACL"},
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func save(cidrs string, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, cidrs)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ExampleScrape()
}
