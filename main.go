package main

import (
	"lemin/pkg"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect format. \nUsage: go run . <filename>")
	}
	// testWork() // Activate this and deactivate the rest of main if you want to test a hard coded example
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	contentArr := strings.Split(string(file), "\n")
	pkg.ProcessData(contentArr)
}

func testWork() {
	graph := pkg.Graph{}

	for i := 0; i < 5; i++ {
		err := graph.AddVertix(strconv.Itoa(i))
		if err != nil {
			log.Fatal("1", err)
		}
	}

	graph.AddConnection("3", "4")

	// err := graph.AddVertix(0)
	// if err != nil {
	// 	log.Fatal("2", err)
	// }

	graph.Print()
}
