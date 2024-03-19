package lemin

import (
	"fmt"
)

func (g *Graph) optimalPath(paths [][]string) {
	x, _ := g.GetVertex(g.startRoom)
	mapping := make(map[string]int)
	for _, node := range x.Connections { // step 1
		// fmt.Println(node.key)
		graph2 := g.Remove(node) // step 2
		newPath := removeFromPath(paths, node.key) 
		// newpath = path without selected node
		numOfSteps := graph2.stepSimulator(newPath) //step 3
		mapping[node.key] = numOfSteps
		// graph2.Print()
		// fmt.Println(numOfSteps)
	}
	fmt.Println(mapping)
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
			// if len(path) + antsinpath[i] > len(paths[i+1]) + antsinpath[i+1]{}
			// antsinpath[i+1]++
			// break

			if (i != len(paths)-1 && counter[i] + len(path) <= len (paths[i+1]) + counter[i+1]) || i == len(paths)-1{
				antsplaced++
				counter[i]++
				break
			} 
			// fmt.Println(len(path), counter[i], i)

		}
	}

	fmt.Print("The queue:" , counter)
	fmt.Println("\n-------------")
	fmt.Println("total ants: " , g.ants, "== ants placed:" , antsplaced)
	
	// finding how many steps it will take from maximum in counter (farthest ant)
	max := counter[0] + len(paths[0])
	for i, x := range counter{
		if max < x + len(paths[i]) {
			max = x
		}
	}
	// return 	stepForward(paths, counter)
	return max

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

func removeFromPath(paths [][]string, key string) [][]string{
	var newpath [][]string
	for _, path := range paths {
		var found = false
		for _, room := range path {
			if room == key {
				found = true
				break
			}
		}
		if !found {
			newpath = append(newpath, path)
		}
	}

	return newpath
}
