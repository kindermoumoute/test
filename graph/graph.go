package main

import "fmt"

var (
	ErrArcExist     = fmt.Errorf("An arc already exists with this ID")
	ErrNodeExist    = fmt.Errorf("A node already exists with this ID")
	ErrNodeNotExist = fmt.Errorf("There is no node with this ID")
)

const (
	MaxInt = 1<<63 - 1
)

type Graph struct {
	ID    string
	Nodes map[string]*Node
	Arcs  map[string]*Arc
}

func NewGraph(id string) Graph {
	return Graph{ID: id, Nodes: make(map[string]*Node), Arcs: make(map[string]*Arc)}
}

func (g Graph) ResetTempWeight() {
	for id := range g.Nodes {
		g.Nodes[id].tempW = MaxInt
	}
}

func (g Graph) AddNode(id string) error {
	_, exist := g.Nodes[id]
	if exist {
		return ErrNodeExist
	}
	g.Nodes[id] = NewNode(id)
	return nil
}

func (g Graph) AddArc(idArc, idIn, idOut string, weight int) error {
	// check if in and out nodes exists
	_, exist := g.Nodes[idIn]
	if !exist {
		return ErrNodeNotExist
	}
	_, exist = g.Nodes[idOut]
	if !exist {
		return ErrNodeNotExist
	}

	// verify an arc doesn't already exists with this id
	_, exist = g.Arcs[idArc]
	if exist {
		return ErrArcExist
	}
	_, exist = g.Nodes[idIn].ArcOut[idArc]
	if exist {
		return ErrArcExist
	}
	_, exist = g.Nodes[idOut].ArcIn[idArc]
	if exist {
		return ErrArcExist
	}

	// Add the new arc to the map of arc and the new node
	g.Arcs[idArc] = NewArc(idArc, weight, g.Nodes[idIn], g.Nodes[idOut])
	g.Nodes[idIn].ArcOut[idArc] = g.Arcs[idArc]
	g.Nodes[idOut].ArcIn[idArc] = g.Arcs[idArc]
	return nil
}

func (g Graph) Delete(id string) error {
	_, exist := g.Nodes[id]
	if !exist {
		return ErrNodeNotExist
	}
	delete(g.Nodes, id)
	return nil
}

func (g Graph) StringByNode() string {
	s := fmt.Sprintf("\nGraph %s", g.ID)
	for _, n := range g.Nodes {
		s += fmt.Sprintf("\n\t%s", n.String())
	}
	return s
}

func (g Graph) StringByArc() string {
	s := fmt.Sprintf("\nGraph %s", g.ID)
	for idArc, a := range g.Arcs {
		s += fmt.Sprintf("\n\t%s - %d - %s (%s)", a.NodeIn.ID, a.Weight, a.NodeOut.ID, idArc)
	}
	return s
}

func (g Graph) findShortestPath(idIn, idOut string) ([]*Node) {
	return g.Nodes[idIn].FindShortestPathTo(idOut, 0)
}
