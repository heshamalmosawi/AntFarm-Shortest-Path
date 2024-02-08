package pkg

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ProcessData(d []string) {
	var (
		graph = Graph{}
		err   error
	)
	for i, line := range d {
		// find number of ants on line no 1
		if i == 0 {
			graph.ants, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		// if comment (single #), continue...
		if len(line) > 2 && line[0] == '#' && line[1] != '#' {
			continue
		}

		// find start and finish
		if len(line) > 2 && line[0] == '#' && line[1] == '#' {
			if line[2:] == "start" {
				// check if ##end immediately follows start, if so -> return err
				if d[i+1] == "##end" {
					fmt.Println("error start and end same room") // error start and end are the same room
				} else {
					// add start room
					graph.startRoom = strings.Split(d[i+1], " ")[0]
					i += 2
					continue
				}

			} else if line[2:] == "end" {
				// add end room
				graph.endRoom = strings.Split(d[i+1], " ")[0]
				i += 2
				continue
			}
		}

		// split lines to determine which lines are room and which ones are connections
		// if len(splitline) > 1, line has room name, x coord, y coord
		// else if len(splitline) = 1 and not empty, line is connection
		splitLine := strings.Split(line, " ")
		if len(splitLine) > 1 { // room
			graph.AddVertix(splitLine[0])
		} else if splitLine[0] != "" { // connection
			//split on '-' to get from and to connections
			splitCon := strings.Split(strings.Join(splitLine, ""), "-")
			from, to := splitCon[0], splitCon[1]
			//create vertix connection
			graph.AddConnection(from, to)
		}
	}
	graph.Print()
}

// func ProcessData(data []string) {
// 	var start, numAnts int
// 	var i = 0
// 	// graph := Graph{}
//
// 	for i < len(data) {
// 		// ignore comments
// 		if data[i][0] == '#' && data[i][1] != '#' {
// 			i++
// 			continue
// 		}
// 		if data[i] == "##start" {
// 			if data[i+1][:2] == "##end" {
// 				log.Fatal("ERROR: Invalid data format. Start and end points are the same room.")
// 			}
// 			for j := 0; j < i; j++ {
// 				// Number of ants is always a number, so if no error, we found a number
// 				num, err := strconv.Atoi(data[j])
// 				if err == nil && numAnts != 0 {
// 					log.Fatal("ERROR: Invalid data format. Number of ants should be mentioned once before the start.")
// 				} else if err == nil {
// 					numAnts = num
// 				}
// 			}
// 			if numAnts == 0 {
// 				log.Fatal("ERROR: Ants not found or too few! Please write an appropriate number of ants.")
// 			}
//
// 			/* ----------------------------------------------------------------------------------- */
// 			for j := i; i < len(data); i++ {
// 				temp, err := strconv.Atoi(strings.Split(data[i], " ")[0])
// 				if err == nil {
// 					start = temp
// 					break
// 				} else {
// 					if data[j] == "##end" {
// 						log.Fatal("ERROR: Invalid data format. Missing rooms")
// 					} else if data[j][0] == '#' && data[j][1] != '#' {
// 						continue
// 					}
// 				}
// 			}
// 		} else if data[i] == "##end" {
// 			// getEnd, check no more rooms at the end
// 			break
// 		} else if data[i][0] == '#' && data[i][1] != '#' {
// 			//error
// 		} else {
// 			// key, err := strconv.Atoi(strings.Split(data[i], " ")[0])
// 			// if err != nil{
// 			// 	log.Fatal("Invalid Error!", err)
// 			// }
// 			// graph.AddVertix(key)
// 		}
//
// 		i++
// 	}
// 	fmt.Println("Number of ants:", numAnts, "\tStart room:", start)
// 	// graph.Print()
// }
