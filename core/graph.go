package core

import "fmt"

type Mark struct {
	PathId   int
	IsMarked bool
}

type Node struct {
	Connections []*Node
	Row         int
	Column      int
	Value       int
	// EndValue is marker for start Node to know where we gonna end
	EndValue int
	IsMarked []Mark
}

type mOperation func(x int) int

var CurrentId = 1

// m stands for Math
var (
	mPlus = func(x int) int {
		x = x + 1
		return x
	}
	mMinus = func(x int) int {
		x = x - 1
		return x
	}
)

func (n *Node) walk(endValue int, current int, pathID int, opertation mOperation) bool {
	n.IsMarked = append(n.IsMarked, Mark{
		IsMarked: true,
		PathId:   pathID,
	})
	n.Value = current

	if endValue == current {
		return true
	}

	current = opertation(current)
	RandomizeArray(n.Connections)

	for _, v := range n.Connections {
		if !v.isMarked() ||
			(v.isMarked() && v.containsPathId(pathID) && v.Value == current) {
			isWalkable := v.walk(endValue, current, pathID, opertation)
			if isWalkable {
				return true
			}
		}
	}

	n.IsMarked = n.IsMarked[:len(n.IsMarked)-1]
	n.Value = 0

	return false
}

func GetNodeAt(nodes []*Node, row, coulmn int) *Node {
	for i := range nodes {
		if nodes[i].Row == row && nodes[i].Column == coulmn {
			return nodes[i]
		}
	}

	return nil
}

func (n *Node) Walk(endValue int, startNum int, pathID int) bool {
	switch {
	case startNum > 1:
		return n.walk(endValue, startNum, pathID, mMinus)
	case startNum == 1:
		return n.walk(endValue, startNum, pathID, mPlus)
	}
	return false
}

func (n *Node) isMarked() bool {
	marked := false
	for _, v := range n.IsMarked {
		marked = marked || v.IsMarked
	}

	return marked
}

func (n *Node) containsPathId(pathID int) bool {
	for _, v := range n.IsMarked {
		if v.PathId == pathID {
			return true
		}
	}

	return false
}

func (n *Node) GetSuqence(pathID int) []*Node {
	result := []*Node{n}
	for _, v := range n.Connections {
		if v.containsPathId(pathID) {
			result = append(result, n.GetSuqence(pathID)...)
		}
	}
	return result
}

func PrintSequnce(nodes []*Node, pathIds []int) {
	result := ""
	for _, pathId := range pathIds {
		for _, v := range nodes {
			for _, m := range v.IsMarked {
				if m.PathId != pathId {
					continue
				}

				result += fmt.Sprintf("%d-%d:%d,", v.Row, v.Column, v.Value)
			}
		}
		result += ";"
	}

	result = result[:len(result)-2]
	fmt.Println(result)
}
