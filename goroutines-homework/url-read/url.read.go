package main

import (
	"fmt"
	"github.com/k3a/html2text"
	"golang.org/x/net/html"
	"goroutines-homework/words"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

type Words struct {
	Key   string
	Value int
}

var Sw []Words

var allWords string

var urls []string

var baseUrl string

//func newUrls(urls string) string { TODO
//	u := strings.Fields(urls)
//	for _, u := range u {
//		nu := baseUrl + u
//		fmt.Println(nu)
//		z := 5
//		fmt.Println(z)
//		return nu
//	}
//	return "no string"
//}

func urlGen(u string) {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r := resp.Body
	for _, l := range getLinks(r) {
		if string(l[0]) == "/" {
			urls = append(urls, l)
		}
	}

}

func ulrRead(u string) []string {

	fmt.Printf("HTML code of %s:\n", u)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r := resp.Body

	html, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	plain := html2text.HTML2Text(string(html))
	w := strings.Fields(plain)

	return w
}

func getLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}
}

func addWords(ws []string) {
	for _, w := range ws {
		if len(allWords) < 5000 {
			allWords += " " + w
		}
	}
}

func sortWords(s string) {

	wc := words.WordCount(s)
	for k, v := range wc {
		Sw = append(Sw, Words{k, v})
	}
	sort.Slice(Sw, func(i, j int) bool {
		return Sw[i].Value > Sw[j].Value
	})
}


//func ExampleScrape() {
//	// Request the HTML page.
//	res, err := http.Get("http://service1001.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//	if res.StatusCode != 200 {
//		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
//	}
//
//	// Load the HTML document
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Find the review items
//	doc.Find("body").Each(func(i int, s *goquery.Selection) {
//		// For each item found, get the title
//		title := s.Find("p").Text()
//		fmt.Printf("Review %d: %s\n", i, title)
//	})
//}

func main() {
	//ExampleScrape()
	//baseUrl := os.Args[1:]
	baseUrl = "http://service1001.com"

	addWords(ulrRead(baseUrl))
	urlGen(baseUrl)

	for _, subUrl := range urls {
		addWords(ulrRead(baseUrl + subUrl))
	}
	sortWords(allWords)

	fmt.Printf("HTML code of %s:\n", baseUrl)
	for _, v := range Sw {
		fmt.Printf("%s: %d times\n", v.Key, v.Value)
	}
}
