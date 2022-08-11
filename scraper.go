package main

import (
	"fmt"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const maxPartitionSize = 50 * 1024

func scrape() {
	docv4 := getDoc("https://www.countryipblocks.net/acl.php")
	ipv4 := collectCIDRs(docv4)
	docv6 := getDoc("https://www.countryipblocks.net/ipv6_acl.php")
	ipv6 := collectCIDRs(docv6)
	save(ipv4, "ipv4")
	save(ipv6, "ipv6")
}

func collectCIDRs(doc *goquery.Document) []string {
	text := doc.Find("#textareaAll").Text()
	cidrs := strings.Split(text, "\n")
	for i, cidr := range cidrs {
		cidrs[i] = strings.TrimSpace(cidr)
	}
	return cidrs
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

func save(cidrs []string, indexName string) {
	indexFile, err := os.Create(indexName)
	if err != nil {
		log.Fatal(err)
	}
	var partitionIndex int
	var partitionSize int
	defer indexFile.Close()
	partitionName := fmt.Sprintf("%s_partition%d", indexName, partitionIndex)
	partitionFile, err := os.Create(partitionName)
	if err != nil {
		log.Fatal(err)
	}
	for _, cidr := range cidrs {
		if partitionSize > maxPartitionSize {
			partitionIndex++
			partitionSize = 0
			_, err = fmt.Fprintln(indexFile, partitionName)
			if err != nil {
				log.Fatal(err)
			}
			partitionFile.Close()
			partitionName = fmt.Sprintf("%s_partition%d", indexName, partitionIndex)
			partitionFile, err = os.Create(partitionName)
			if err != nil {
				log.Fatal(err)
			}
		}
		line := cidr + "\n"
		partitionSize += len(line)
		_, err = fmt.Fprint(partitionFile, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	if partitionSize > 0 {
		_, err = fmt.Fprintln(indexFile, partitionName)
		if err != nil {
			log.Fatal(err)
		}
		partitionFile.Close()
	}
}

func main() {
	scrape()
}
