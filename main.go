package main

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices []*Vertix
}
type Vertix struct {
	key         int
	Connections []*Vertix
}

func (g *Graph) addVertix(key int) error {
	if g.contains(key) {
		err := fmt.Sprintf("Vertix %d already exists", key)
		return errors.New(err)
	}
	g.Vertices = append(g.Vertices, &Vertix{key: key})
	return nil
}

func (g *Graph) addConnection(from, to int) {

	Vfrom, err := g.getVertix(from)
	if err != nil {
		fmt.Println(err)
		return
	}
	Vto, err := g.getVertix(to)
	if err != nil {
		fmt.Println(err)
		return
	}
	Vfrom.Connections = append(Vfrom.Connections, Vto)
	Vto.Connections = append(Vfrom.Connections, Vfrom)
}

func (g *Graph) getVertix(key int) (*Vertix, error) {
	for _, v := range g.Vertices {
		if v.key == key {
			return v, nil
		}
	}
	return nil, errors.New("No vertix exists")
}

func (g *Graph) contains(key int) bool {
	for _, v := range g.Vertices {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g Graph) print() {
	for _, v := range g.Vertices {
		fmt.Printf("Vertix key: %d\n", v.key)
	}
}

// TODO: test vertices connections
// TODO: print connections
func main() {
	graph := Graph{}

	for i := 0; i < 5; i++ {
		err := graph.addVertix(i)
		if err != nil {
			fmt.Println(err)
		}
	}

	err := graph.addVertix(0)
	if err != nil {
		fmt.Println(err)
	}

	graph.print()
}
