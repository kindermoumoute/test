package main

type Arc struct {
	NodeIn, NodeOut *Node
	Weight          int
}

func NewArc(id string, weight int, nodeIn, nodeOut *Node) *Arc {
	return &Arc{NodeIn: nodeIn, NodeOut: nodeOut, Weight: weight}
}
