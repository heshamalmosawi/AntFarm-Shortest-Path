package lemin

import (
	"fmt"
	"math"
)

/* This function finds the optimal path to use then proceeds to call the printsteps function to walk the steps and print them. */
func (g *Graph) optimalPath(paths [][]string) {
	// fmt.Println("\n\n", (paths), "\n\n")

	visited := make(map[string]bool) 
	combinations := GenerateCombinations(paths, visited)
	fmt.Println("\n-combination done-\n")
	fmt.Println("length:", len(combinations))
	if len(combinations) == 0 {
		fmt.Println(paths)
	}
	for i := range combinations{
		combinations[i] = disjointPaths(combinations[i])
		fmt.Println("\n\n\n", combinations[i])
	}
	combinations = removeDuplicates(combinations)
	mapping := make([]int, len(combinations))

		for i := range combinations  { // step 1
			fmt.Println("\033[34m", combinations[i], "\033[0m")
			
			mapping[i] = g.stepSimulator(combinations[i]) //step 3
			// graph2.Print()
			// fmt.Println(numOfSteps)
		}
		// fmt.Println(mapping)
		var minIndex int
		minValue := math.MaxInt64
		for index, value := range mapping {
			if value < minValue {
				minIndex = index
				minValue = value
			}
		}
		fmt.Println("::", len(mapping), "::", len(combinations))
		fmt.Println("--", mapping[minIndex], ": ", combinations[minIndex])

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

			if (i != len(paths)-1 && counter[i]+len(path) <= len(paths[i+1])+counter[i+1]) || i == len(paths)-1{
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

/* This function removes any path containing the vertex passed as parameter. */
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

func GenerateCombinations(paths[][]string, visitedRooms map[string]bool) [][][]string{
	fmt.Print("length of paths", len(paths), paths)
	if len(paths) == 1 {
		for i, room := range paths[0]{
			if i == 0 || i == len(paths[0])-1 {
				continue
			}
			if visitedRooms[room] {
				return nil
			}
		}
		return [][][]string{paths}
	}
	firstPath := paths[0]
	// tester := false
	for _, room := range firstPath{
		if visitedRooms[room] {
			return GenerateCombinations(paths[1:], visitedRooms)
		// break
		}
		// visitedRooms[room] = true
	}
	// if tester {
		// for _, room := range firstPath{
		// 	if visitedRooms[room]{
		// 		visitedRooms[room] = false
		// 	}
		// }
		// return GenerateCombinations(paths[1:], visitedRooms)
	// }
	// fmt.Println(paths)
	combsWithoutFirst := GenerateCombinations(paths[1:], visitedRooms)
	if combsWithoutFirst == nil {
		return [][][]string{{firstPath}}
	}
	// fmt.Print("\x1b[34m", len(combsWithoutFirst), "\x1b[0m")

	allCombsWithFirst := [][][]string{}
	for i := range combsWithoutFirst{
		oneCombWithFirst := append([][]string{firstPath}, combsWithoutFirst[i]...)
		allCombsWithFirst = append(allCombsWithFirst, oneCombWithFirst)
	}
	fmt.Print(len(allCombsWithFirst))
	if len(allCombsWithFirst) > 0 {
		return append(combsWithoutFirst, allCombsWithFirst...)
	} else {
		return combsWithoutFirst
	}
}

func removeDuplicates(input [][][]string) [][][]string {
    uniqueMap := make(map[string]bool)
    var result [][][]string

    for _, arr := range input {
        arrString := fmt.Sprintf("%v", arr)
        if !uniqueMap[arrString] {
            uniqueMap[arrString] = true
            result = append(result, arr)
        }
    }

    return result
}