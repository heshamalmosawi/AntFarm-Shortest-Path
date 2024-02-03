package main

import (
	"errors"
	"fmt"
	"os"
)

type Graph struct {
	ants     int
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
	Vto.Connections = append(Vto.Connections, Vfrom)
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
	fmt.Println("Vertices:")
	for _, v := range g.Vertices {
		fmt.Printf("Vertix key: %d\n", v.key)
	}

	fmt.Println("Connections:")
	for _, v := range g.Vertices {
		for _, c := range v.Connections {
			if v.key < c.key {
				fmt.Printf("%d--%d\n", v.key, c.key)
			}
		}
	}
}

// TODO: test vertices connections
// TODO: print connections
func main() {
	var file string
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: go run . [ARG]")
		return
	} else {
		file = fmt.Sprintf("./examples/%s", args[1])
		fmt.Printf("file directory: %s\n", file)
	}
	reader, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", string(reader))

	// graph := Graph{}
	//
	// for i := 0; i <= 5; i++ {
	// 	err := graph.addVertix(i)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	//
	// graph.addConnection(0, 1)
	// graph.addConnection(2, 1)
	// graph.addConnection(4, 5)
	//
	// for _, v := range graph.Vertices {
	// 	for _, c := range v.Connections {
	// 		fmt.Printf("Vertix %d is connected to Vertix %d\n", v.key, c.key)
	// 	}
	// }
	// graph.print()
}
