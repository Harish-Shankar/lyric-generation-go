package main

import (
	"bufio"
	"flag"
	"io"
	"log"
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
}
