package pkg

import (
	"strconv"
	"strings"

	"github.com/anton-uvarenko/suqences_alg/core"
)

type inputData struct {
	StartValue int
	EndValue   int
}

func GetStartsByArgs(args []string, nodes []*core.Node) ([]*core.Node, []int) {
	result := []*core.Node{}
	pathIds := []int{}
	data := parseArgs(args)
	for i, d := range data {
		node := getRandomStart(nodes, i)
		node.Value = d.StartValue
		node.EndValue = d.EndValue
		result = append(result, node)
		pathIds = append(pathIds, i+1)
	}
	return result, pathIds
}

func parseArgs(args []string) []inputData {
	result := []inputData{}
	for _, v := range args {
		rawData := strings.Split(v, "-")
		startVal, _ := strconv.Atoi(rawData[0])
		endVal, _ := strconv.Atoi(rawData[1])

		result = append(result, inputData{
			StartValue: startVal,
			EndValue:   endVal,
		})
	}
	return result
}
