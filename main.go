package main

import (
	"log"

	teamodels "github.com/anton-uvarenko/suqences_alg/core/tea_models"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MAXW      = 6
	MAXH      = 10
	SeqLength = 10
)

func main() {
	nodes := CreateSequencesGraph(MAXW, MAXH)
	model := &Model{
		PathsModel: &teamodels.PathsModel{},
		NodeModel: &teamodels.NodeModel{
			Nodes: nodes,
		},
		InputModel: &teamodels.InputModel{},
	}
	model.PathsModel.InitModel()

	p := tea.NewProgram(model)
	addSuqenceParams(model)

	// var isStart1, isStart2, isStart3 bool
	//
	// for !isStart1 || !isStart2 || !isStart3 {
	// 	clearNodes(nodes)
	// 	start1 := getRandomStart(nodes, 1)
	// 	isStart1 = start1.Walk(12, 12, 1)
	//
	// 	start2 := getRandomStart(nodes, 2)
	// 	isStart2 = start2.Walk(8, 1, 2)
	//
	// 	start3 := getRandomStart(nodes, 3)
	// 	isStart3 = start3.Walk(10, 1, 3)
	// }

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
