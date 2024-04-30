package main

import (
	"os"
	"strings"
)

const (
	MAXW      = 6
	MAXH      = 10
	SeqLength = 10
)

func main() {
	arguments := strings.Split(os.Args[0], ";")
	CreateSequencesGraph(MAXW, MAXH)
}
