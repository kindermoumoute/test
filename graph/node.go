package main

import "fmt"

type Node struct {
	ID            string
	ArcIn, ArcOut map[string]*Arc
	tempW         int
}

func NewNode(id string) *Node {
	return &Node{ID: id, ArcOut: make(map[string]*Arc), ArcIn: make(map[string]*Arc), tempW: MaxInt}
}

//Return string representation of the node
func (n *Node) String() string {
	s := fmt.Sprintf("NodeOut %s - %d", n.ID, n.tempW)
	for idArc, a := range n.ArcOut {
		s += fmt.Sprintf("\n\t\t%s - %d - %s", idArc, a.Weight, a.NodeOut.ID)
	}
	return s
}

func (n *Node) FindShortestPathTo(id string, w int) []*Node {
	if w <= n.tempW {
		n.tempW = w
		if n.ID == id {
			return []*Node{n}
		}
	} else {
		return nil
	}

	var currentPath []*Node
	for _, a := range n.ArcOut {
		path := a.NodeOut.FindShortestPathTo(id, w+a.Weight)
		if path != nil {
			currentPath = append(path, n)
		}
	}

	return currentPath
}
