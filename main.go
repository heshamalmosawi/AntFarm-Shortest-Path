package main

import (
	"lemin/lemin"
	"log"
	"os"
	"strings"
	// lemin "lemin/lemin"
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
	err = lemin.Farm.ValidCoord()
	if err != nil {
		log.Fatal(err)
	}
	// lemin.Farm.PathFinder()

	// example00
	// paths := [][]string{
	// 	{"start", "2", "3", "1", "end"},
	// }

	// example01
	// var paths = [][]string{
	// 	{"start", "t", "E", "a", "m", "end"},
	// 	{"start", "h", "A", "c", "k", "end"},
	// 	{"start", "0", "o", "n", "e", "end"},
	// }

	// example01
	paths := [][]string{
		{"start", "3", "end"},
		{"start", "1", "2", "3", "end"},
	}
	lemin.Farm.PathWalker(paths)
}
