package main

import (
	lemin "lemin/lemin"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect format. \nUsage: go run . <filename>")
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	contentArr := strings.Split(string(file), "\n")
	lemin.ProcessData(contentArr)
	lemin.Farm.ValidCoord()
	lemin.Farm.PathFinder()
}
