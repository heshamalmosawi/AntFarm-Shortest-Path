package pkg

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices  []*Vertix
	startRoom string
	endRoom   string
	ants      int
}
type Vertix struct {
	key         string
	Connections []*Vertix
}

func (g *Graph) AddVertix(key string) error {
	if g.Contains(key) {
		err := fmt.Sprintf("Vertix %s already exists", key)
		return errors.New(err)
	}
	g.Vertices = append(g.Vertices, &Vertix{key: key})
	return nil
}

func (g *Graph) AddConnection(from, to string) {
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
	Vto.Connections = append(Vto.Connections, Vfrom)
}

func (g *Graph) getVertix(key string) (*Vertix, error) {
	for _, v := range g.Vertices {
		if v.key == key {
			return v, nil
		}
	}
	return nil, errors.New("no vertix exists")
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
	for _, v := range g.Vertices {
		fmt.Printf("Vertix key:" + v.key)
		fmt.Print("\tIts connections: ")
		for _, connection := range v.Connections {
			fmt.Print(connection.key, " ")
		}
		fmt.Println()
	}
}
