package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Chain struct {
	order       int
	suffix      map[string][]string
	capitalized int
}

func newChainFromFile(filename string, order int) (*Chain, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	return newChain(file, order)
}

func newChain(reader io.Reader, order int) (*Chain, error) {
	chain := &Chain{
		order:  order,
		suffix: make(map[string][]string),
	}

	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	window := make([]string, order)
	for sc.Scan() {
		word := sc.Text()
		if len(window) > 0 {
			prefix := strings.Join(window, " ")
			chain.suffix[prefix] = append(chain.suffix[prefix], word)
			if isCapital(prefix) {
				chain.capitalized++
			}
		}
		window = append(window[1:], word)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return chain, nil
}

func isCapital(word string) bool {
	r, _ := utf8.DecodeRuneInString(word)
	return unicode.IsUpper(r)
}

func (c *Chain) genSentence(n int, startWithCapital bool) string {
	var i int
	var words []string
	var prefix string

	if startWithCapital {
		i = rand.Intn(c.capitalized)
	} else {
		i = rand.Intn(len(c.suffix))
	}

	for prefix = range c.suffix {
		if startWithCapital && !isCapital(prefix) {
			continue
		}
		if i == 0 {
			break
		}
		i--
	}

	words = append(words, prefix)
	prefixWords := strings.Split(prefix, " ")
	n -= len(prefixWords)
	for {
		wordChoices := c.suffix[prefix]
		if len(wordChoices) == 0 {
			break
		}
		i = rand.Intn(len(wordChoices))
		suffix := wordChoices[i]
		words = append(words, suffix)

		n--
		if n < 0 || isSentenceEnd(suffix) {
			break
		}
		prefixWords = append(prefixWords[1:], suffix)
		prefix = strings.Join(prefixWords, " ")
	}
	return strings.Join(words, " ")

}

func isSentenceEnd(word string) bool {
	w, _ := utf8.DecodeLastRuneInString(word)
	return w == '.' || w == '?' || w == '!'
}

func main() {
	inputFile := flag.String("in", "lyrics.txt", "input file")
	numWords := flag.Int("n", 1, "number of words to use as a prefix")
	numRuns := flag.Int("runs", 20, "number of runs to generate")
	wordsPerRun := flag.Int("words", 20, "number of words per run")
	startOnCapital := flag.Bool("capital", false, "start output with a capitalized prefix")
	flag.Parse()

	c, errChain := newChainFromFile(*inputFile, *numWords)
	if errChain != nil {
		log.Fatal(errChain)
	}

	f, errFile := os.OpenFile("generatedLyrics.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	os.Truncate("generatedLyrics.txt", 0)
	if errFile != nil {
		log.Fatal(errFile)
	}

	defer f.Close()

	for i := 0; i < *numRuns; i++ {
		out := c.genSentence(*wordsPerRun, *startOnCapital)
		fmt.Println(out)
		f.WriteString(out + "\n")
	}
}
