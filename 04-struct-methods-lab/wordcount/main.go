package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"

	//"github.com/iproduct/coursego/03-types-lab/wordcount"
	//"strings"
)

var stopwordsList = []string{`ourselves`, `hers`, `between`, `yourself`, `but`, `again`, `there`, `about`, `once`,
	`during`, `out`, `very`, `having`, `with`, `they`, `own`, `an`, `be`, `some`, `for`, `do`, `its`, `yours`,
	`such`, `into`, `of`, `most`, `itself`, `other`, `off`, `is`, `s`, `am`, `or`, `who`, `as`, `from`, `him`,
	`each`, `the`, `themselves`, `until`, `below`, `are`, `we`, `these`, `your`, `his`, `through`, `don`, `nor`,
	`me`, `were`, `her`, `more`, `himself`, `this`, `down`, `should`, `our`, `their`, `while`, `above`, `both`,
	`up`, `to`, `ours`, `had`, `she`, `all`, `no`, `when`, `at`, `any`, `before`, `them`, `same`, `and`, `been`,
	`have`, `in`, `will`, `on`, `does`, `yourselves`, `then`, `that`, `because`, `what`, `over`, `why`, `so`,
	`can`, `did`, `not`, `now`, `under`, `he`, `you`, `herself`, `has`, `just`, `where`, `too`, `only`, `myself`,
	`which`, `those`, `i`, `after`, `few`, `whom`, `t`, `being`, `if`, `theirs`, `my`, `against`, `a`, `by`,
	`doing`, `it`, `how`, `further`, `was`, `here`, `than`}

var stopwords map[string] struct{}
var splitRegex *regexp.Regexp // pointer

type Entry struct {
	Word string
	Count int
}

func main() {
	/// filiing stopwords to hashset
	stopwords = make(map[string] struct{}, 100)
	for _, w := range stopwordsList {
		stopwords[w] = struct{} {}
	}

	// initialize regex for splitting
	splitRegex = regexp.MustCompile(`[\s,\.!?\"]`)

	// count words
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, fname := range files {
		err := processFIle(fname, counts)
		if err != nil {
			log.Printf("file '%s' not found", fname)
		}
	}
	//counts map -> []Entry
	entries := make([]Entry, len(counts))
	i := 0
	for w, c := range counts {
		entries[i] = Entry{w, c}
		i++
	}

	// sort entries by count Lambda func example
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count == entries[j].Count{
			return entries[i].Word < entries[j].Word
		} else {
			return entries[i].Count > entries[j].Count
		}
	})

	for _, e := range entries[:10] {
		//fmt.Printf("%s -> %d\n", w, n)
		fmt.Printf("%v\n", e)
	}
}

func processFIle(fname string, counts map[string]int) error {
	file, err := os.Open(fname)
	if err != nil {
		log.Printf("file '%s' not found", fname)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Entered: %s\n", line)
		//if line == "" {
		//	break
		//}
		//words := strings.Fields(line)
		words := splitRegex.Split(line, -1)
		fmt.Printf("%v\n", words)
		for _, w := range words {
			if _, ok := stopwords[w]; !ok && len(w) > 1 {
				counts[w]++
			}

		}
	}
	return nil // return nothing if no mistake
}