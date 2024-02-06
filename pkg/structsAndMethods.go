package pkg

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices []*Vertix
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
	// fmt.Print("from is", from)
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
	// fmt.Println(Vto)
	Vfrom.Connections = append(Vfrom.Connections, Vto)
	Vto.Connections = append(Vto.Connections, Vfrom)
	fmt.Println("f ", Vfrom.Connections[0].key)
}

func (g *Graph) getVertix(key string) (*Vertix, error) {
	for _, v := range g.Vertices {
		if v.key == key {
			return v, nil
		}
	}
	// fmt.Print(key)
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
