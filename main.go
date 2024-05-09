package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/anton-uvarenko/suqences_alg/core"
	"github.com/anton-uvarenko/suqences_alg/pkg"
)

func main() {
	rawWidthHeight := strings.Split(os.Args[1], "-")
	width, _ := strconv.Atoi(rawWidthHeight[0])
	height, _ := strconv.Atoi(rawWidthHeight[1])

	arguments := os.Args[2:]
	nodes := pkg.CreateSequencesGraph(width, height)
	starts, pathIds := pkg.GetStartsByArgs(arguments, nodes)

	for i, v := range starts {
		v.Walk(v.EndValue, v.Value, pathIds[i])
	}

	// for _, v := range nodes {
	// 	fmt.Print(v.Value, " ")
	// 	if v.Column == width-1 {
	// 		fmt.Println()
	// 	}
	// }

	core.PrintSequnce(nodes, pathIds)
}
