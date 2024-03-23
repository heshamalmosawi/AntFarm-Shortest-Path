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
		newPath := RemoveFromPath(paths, node.key)
		if len(newPath) == 0 {
			continue
		}
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

			if (i != len(paths)-1 && counter[i]+len(path) <= len(paths[i+1])+counter[i+1]) || i == len(paths)-1 {
				antsplaced++
				counter[i]++
				break
			}
			// fmt.Println(len(path), counter[i], i)

		}
	}

	fmt.Print("The queue:", counter)
	fmt.Println("\n-------------")
	fmt.Println("total ants: ", g.ants, "== ants placed:", antsplaced)

	// finding how many steps it will take from maximum in counter (farthest ant)
	max := counter[0] + len(paths[0])
	for i, x := range counter {
		if max < x+len(paths[i]) {
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

func RemoveFromPath(paths [][]string, key string) [][]string {
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

type ant struct {
	name         string
	path         []string
	pathIndx     int
	currRoomIndx int
	isWalking    bool
}
type round struct {
	ants         []ant
	currPathIndx int
	rounds       []string // might not need it
	queue        map[string]int
	numOfPaths   int
}

func (g Graph) PathWalker(paths [][]string) {
	var round = round{}
	round.numOfPaths = len(paths)
	round.queue = make(map[string]int, len(paths))
	// initialize the queue
	for _, path := range paths {
		_, ok := round.queue[path[1]]
		if !ok {
			round.queue[path[1]] = 0
		}
	}
	// cosntruct ants
	var counter = 1
	for i := 1; i <= g.ants; i++ {
		var (
			nextPath     []string
			nextPathCap  int
			ants         int
			ant          = ant{}
			currPathIndx = round.currPathIndx
			currPath     = paths[currPathIndx]
			currPathCap  = len(paths[currPathIndx][1 : len(paths[currPathIndx])-1])
		)

		ant.name = fmt.Sprintf("L%d", i)
		ant.currRoomIndx = 1
		ants, _ = round.queue[currPath[1]] // map indexed by first element after startRoom
		currPathCap += ants
		if i == 1 {
			ant.path = paths[round.currPathIndx]
			ant.pathIndx = round.currPathIndx
			round.ants = append(round.ants, ant)
			round.queue[paths[currPathIndx][1]]++
			continue
		}
		// Determine which path the ant will traverse
		// First find: nextPath, ants and nextPath values
		// These values help determine the capacities of current and next paths (to help determine the path of the ant)
		if currPathIndx+1 < round.numOfPaths {
			// if next path exists (currentPathIndx + 1 < length paths)
			// calculate values for next path
			tmp := paths[currPathIndx+1]
			nextPath = tmp[1 : len(tmp)-1]
			ants, _ := round.queue[nextPath[1]]
			nextPathCap = len(nextPath) + ants
			// i doubt that below code is redundant since the case is solved in the else statement for the next if statement
		} //else {
		// else set ant to first path (paths[0])
		// ant.path = paths[0]
		// ant.pathIndx = 0
		// round.currPathIndx = 0
		// round.ants = append(round.ants, ant)
		// continue
		// }

		if currPathCap <= nextPathCap {
			fmt.Printf("@Interation: %d\n    currPathCap: %d\n    nextPathCap: %d\n", counter, currPathCap, nextPathCap)
			counter += 1
			ant.path = paths[round.currPathIndx]
			ant.pathIndx = round.currPathIndx
			round.ants = append(round.ants, ant)
			round.queue[paths[currPathIndx][1]]++
		} else {
			if currPathIndx == len(paths)-1 {
				ant.path = paths[0]
				ant.pathIndx = 0
				round.currPathIndx = 0
				round.ants = append(round.ants, ant)
				round.queue[paths[0][1]]++
			} else {
				ant.path = paths[round.currPathIndx+1]
				ant.pathIndx = round.currPathIndx + 1
				round.currPathIndx += 1
				round.ants = append(round.ants, ant)
				round.queue[paths[currPathIndx+1][1]]++
			}
		}
	}
	for _, ant := range round.ants {
		fmt.Printf("ant:%s\n    path: %v\n    pathIndx: %d\n", ant.name, ant.path, ant.pathIndx)
	}
}
