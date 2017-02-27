package main

import "fmt"

func main() {
	g := NewGraph("My Map")
	g.makeMap()
	path := g.findShortestPath("Paris", "Hillsboro")

	//fmt.Println(g.StringByNode())

	// display the path
	fmt.Printf("\nResult of the shortest path between %s and %s\n", "Paris", "Hillsboro")
	for _, n := range path {
		fmt.Printf(" %s <-", n.ID)
	}
	fmt.Printf(" (%d km)", path[0].tempW)
}

func (g Graph) makeMap() {
	g.AddNode("Paris")
	g.AddNode("London")
	g.AddNode("Lens")
	g.AddNode("Rennes")
	g.AddNode("Brest")
	g.AddNode("Hillsboro")

	g.AddArc("ParisToLondon", "Paris", "London", 300)
	g.AddArc("ParisToLens", "Paris", "Lens", 100)
	g.AddArc("ParisToRennes", "Paris", "Rennes", 200)

	g.AddArc("RennesToBrest", "Rennes", "Brest", 100)
	g.AddArc("RennesToLondon", "Rennes", "London", 100)

	g.AddArc("HillsboroToLens", "Hillsboro", "Lens", 1000)

	g.AddArc("LensToHillsboro", "Lens", "Hillsboro", 2000)
	g.AddArc("LensToRennes", "Lens", "Rennes", 250)
	g.AddArc("LensToParis", "Lens", "Paris", 100)

	g.AddArc("BrestToLondon", "Brest", "London", 50)

	g.AddArc("LondonToBrest", "London", "Brest", 50)
}
