package lemin

import (
	"fmt"
	"os"
)

func (g *Graph) optimalPath(paths [][]string) {
	x, _ := g.GetVertex(g.startRoom)
	mapping := make(map[string]int)
	for _, node := range x.Connections { // step 1
		fmt.Println(node.key)
		graph2 := g.Remove(node) // step 2
		newPath := removeFromPath(paths, node.key) 
		numOfSteps := graph2.stepSimulator(newPath) //step 3
		mapping[node.key] = numOfSteps
		// graph2.Print()
		fmt.Print(numOfSteps)
		os.Exit(1)
	}
}

func (g *Graph) stepSimulator(paths [][]string) int {
	fmt.Println("\n------------ Step Sim ------------")
	for _, path := range paths {
		fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
	}

	counter := make([]int, len(paths)) // to keep track of calculations
	antsplaced := 0
	for antsplaced < g.ants {
	 	for i, path := range paths {
			// if len(path) + antsinpath[i] > len(paths[i+1]) + antsinpath[i+1]
			// antsinpath[i+1]++
			// break

			if i != len(paths)-1 && counter[i] + len(path)-2 > len (paths[i+1]){
				antsplaced++
				counter[i]++
			} else if i == len(paths)-1 && counter[i] + len(path)-2 > len (paths[i+1]){
				antsplaced++
				counter[i]++
				break
			}
			fmt.Println(len(path), counter[i], i)

		}
	}
	fmt.Println("-------------")
	fmt.Println(g.ants, "==" , antsplaced)
	// return 	stepForward(paths, counter)
	return 0

}

// func stepForward(paths [][]string, counter []int) int{
// 	x := 0
// 	for i, path := range paths{
// 		if len(path) - 2 != counter[i]{ 
// 			counter[i]++
// 			// break
// 		}
// 		x = i
// 	}
// 	return x
// }

// func removeFromPath(paths [][]string, key string) [][]string{
// 	var newpath [][]string
// 	for _, path := range paths {
// 		var found = false
// 		for _, room := range path {
// 			if room == key {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			newpath = append(newpath, path)
// 		}
// 	}

// 	return newpath
// }


// func (g *Graph) stepSimulator(paths [][]string) int {
// 	// fmt.Println("\n------------ Step Sim ------------")
// 	for _, path := range paths {
// 		fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
// 	}
// 	antsRemainingInStart := g.ants
// 	counter := make([]int, len(paths)) // to keep track of calculations
// 	index := 0
// 	steps := -1

// 	for i := range antsRemainingInStart{
// 		i = i;
// 	}
// 	for antsRemainingInStart != 0 {
// 		steps++
// 		counter = stepForward(antsRemainingInStart, paths, counter)
// 		fmt.Println(len(paths[index])-2, "?:", counter[index])
// 		if len(paths[index])-2 != counter[index]{
// 			counter[index]++
// 			antsRemainingInStart--
// 			continue
// 		}
// 		if len(paths[index])-2 == counter[index]{
// 			fmt.Println("ants at", index , ": ", counter[index])
// 			break			
// 		}
// 	}
// 	return 0
// }