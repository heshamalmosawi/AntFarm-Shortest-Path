package pkg

import (
	"errors"
	"fmt"
	"log"
	"slices"
)

type Graph struct {
	Vertices  []*vertex
	startRoom string
	endRoom   string
	ants      int
}
type vertex struct {
	key         string
	Connections []*vertex
}

func (g *Graph) Addvertex(key string) error {
	if g.Contains(key) {
		err := fmt.Sprintf("vertex %s already exists", key)
		return errors.New(err)
	}
	g.Vertices = append(g.Vertices, &vertex{key: key})
	return nil
}

func (g *Graph) AddConnection(from, to string) {
	Vfrom, err := g.getvertex(from)
	if err != nil {
		log.Fatal(err)
	}
	Vto, err := g.getvertex(to)
	if err != nil {
		log.Fatal(err)
	}
	if slices.Contains(Vfrom.Connections, Vto) || slices.Contains(Vto.Connections, Vfrom) {
		log.Fatal("ERROR: Two rooms can't be connected with more than two lines.")
	}
	Vfrom.Connections = append(Vfrom.Connections, Vto)
	Vto.Connections = append(Vto.Connections, Vfrom)
}

func (g *Graph) getvertex(key string) (*vertex, error) {
	for _, v := range g.Vertices {
		if v.key == key {
			return v, nil
		}
	}
	return nil, errors.New("no vertex exists")
}

func (g *Graph) Contains(key string) bool {
	for _, v := range g.Vertices {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	fmt.Println("Ants:", g.ants, "Start:", g.startRoom, "End:", g.endRoom)
	for _, v := range g.Vertices {
		fmt.Printf("vertex key:" + v.key)
		fmt.Print("\tIts connections: ")
		for _, connection := range v.Connections {
			fmt.Print(connection.key, " ")
		}
		fmt.Println()
	}
}
