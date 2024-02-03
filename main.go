package main

import (
	"lemin/pkg"
	"log"
	"os"
	"strings"
)


func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect format. \nUsage: go run . <filename>")
	}
	// testWork()
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	contentArr := strings.Split(string(file), "\n")
	pkg.ProcessData(contentArr)

	// fmt.Println(temp)
}

func testWork(){
	graph := pkg.Graph{}

	for i := 0; i < 5; i++ {
		err := graph.AddVertix(i)
		if err != nil {
			log.Fatal("1", err)
		}
	}

	graph.AddConnection(2,3)

	// err := graph.AddVertix(0)
	// if err != nil {
	// 	log.Fatal("2", err)
	// }

	graph.Print()
}