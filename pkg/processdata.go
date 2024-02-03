package pkg

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ProcessData(data []string){
	var start, numAnts int
	var i = 0
	graph := Graph{}

	for (i<len(data)){
		if data[i][0] == '#' && data[i][1] != '#'{
			i++
			continue
		}
		if data[i] == "##start"{
			for j := 0; j < i; j++{
				num, err := strconv.Atoi(data[j])
				if err == nil && numAnts != 0 {
					log.Fatal("ERROR: Invalid data format. Number of ants should be mentioned once before the start.")
				} else if err == nil {
					numAnts = num
				}
			}
			if numAnts == 0{
				log.Fatal("ERROR: Ants not found or too few! Please write an appropriate number of ants.")
			}
			for j := i ; i< len(data); i++{
				temp, err := strconv.Atoi(strings.Split(data[i], " ")[0])
				if err == nil {
					start = temp
					break
				} else {
					if data[j] == "##end"{
						log.Fatal("ERROR: Invalid data format. Missing rooms")
					} else if data[j][0] == '#' && data[j][1] != '#' {
						continue
					}
				}
			}
		} else if data[i] == "##end"{
						// getEnd, check no more rooms at the end
			break
		} else if data[i][0] == '#' && data[i][1] != '#'{
			//error
		} else {
			key, err := strconv.Atoi(strings.Split(data[i], " ")[0])
			if err != nil{
				log.Fatal("Invalid Error!", err)
			}
			graph.AddVertix(key)
		}
		
		i++
	}
	fmt.Println("Number of ants:", numAnts,"\tStart room:", start)
}