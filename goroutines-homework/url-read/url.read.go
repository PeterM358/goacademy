package main

import (
	"fmt"
	"github.com/k3a/html2text"
	"goroutines-homework/words"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Words struct {
	Key string
	Value int
}

var Sw []Words

var allWords string

var urls string

var url string

//func newUrls(urls string) string { TODO
//	u := strings.Fields(urls)
//	for _, u := range u {
//		nu := url + u
//		fmt.Println(nu)
//		z := 5
//		fmt.Println(z)
//		return nu
//	}
//	return "no string"
//}

func ulrRead(u string) []string{

	fmt.Printf("HTML code of %s:\n", u)
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	plain := html2text.HTML2Text(string(html))
	w := strings.Fields(plain)
	return w
}

func addWords(ws []string) {
	for _, w := range ws {
		if len(allWords) < 5000 {
			path := []rune(w)
			firstChar := string(path[0:1])
			//fmt.Println(firstChar)
			if firstChar == "/" { //adding urls
				urls += " " + w
			}
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


func main() {
	fmt.Scanln(&url)

	addWords(ulrRead(url))
	//addWords(ulrRead(newUrls(urls))) TODO read from inner urls
	sortWords(allWords)

	//fmt.Println(urls)

	fmt.Printf("HTML code of %s:\n", url)
	for _, v := range Sw {
		fmt.Printf("%s: %d times\n", v.Key, v.Value)
	}
}
