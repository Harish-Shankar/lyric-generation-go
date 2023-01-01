package main

import "flag"

func main() {
	inputFile := flag.String("in", "lyrics.txt", "input file")
	numWords := flag.Int("n", 1, "number of words to use as a prefix")
	numRuns := flag.Int("runs", 20, "number of runs to generate")
	wordsPerRun := flag.Int("words", 20, "number of words per run")
	startOnCapital := flag.Bool("capital", false, "start output with a capitalized prefix")
	flag.Parse()
}
