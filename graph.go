package main

import (
	"math/rand"

	"github.com/anton-uvarenko/suqences_alg/core"
)

func CreateSequencesGraph(width, height int) []*core.Node {
	nodeSlice := createNodes(width, height)
	setupConnections(nodeSlice, width, height)
	return nodeSlice
}

func createNodes(width, height int) []*core.Node {
	var nodes []*core.Node
	for i := range height {
		for j := range width {
			nodes = append(nodes, &core.Node{
				Row:    i,
				Column: j,
			})
		}
	}

	return nodes
}

func setupConnections(nodes []*core.Node, width, height int) {
	for i, v := range nodes {
		if v.Column != 0 {
			v.Connections = append(v.Connections, nodes[i-1])
		}
		if v.Column != width-1 {
			v.Connections = append(v.Connections, nodes[i+1])
		}
		if v.Row != height-1 {
			v.Connections = append(v.Connections, nodes[i+width])
		}
		if v.Row != 0 {
			v.Connections = append(v.Connections, nodes[i-width])
		}
	}
}

func getRandomStart(nodes []*core.Node, pathID int) *core.Node {
	node := nodes[rand.Intn(len(nodes)-1)]
	for node.Value != 0 && pathID != 1 {
		node = nodes[rand.Intn(len(nodes)-1)]
	}
	return node
}

func clearNodes(nodes []*core.Node) {
	for _, v := range nodes {
		v.IsMarked = []core.Mark{}
		v.Value = 0
	}
}
