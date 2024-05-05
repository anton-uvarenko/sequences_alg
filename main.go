package main

import (
	"os"

	"github.com/anton-uvarenko/suqences_alg/core"
	"github.com/anton-uvarenko/suqences_alg/pkg"
)

const (
	MAXW      = 6
	MAXH      = 10
	SeqLength = 10
)

func main() {
	arguments := os.Args[1:]
	nodes := pkg.CreateSequencesGraph(MAXW, MAXH)
	starts, pathIds := pkg.GetStartsByArgs(arguments, nodes)
	for _, v := range starts {
		v.Walk(v.EndValue, v.Value, v.IsMarked[0].PathId)
	}
	core.PrintSequnce(nodes, pathIds)
}
