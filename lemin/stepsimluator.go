package lemin

import (
	"fmt"
	"math"
)

/* This function finds the optimal path to use then proceeds to call the printsteps function to walk the steps and print them. */
func (g *Graph) optimalPath(paths [][]string) {
	if len(paths) > 20 {
		paths = paths[:15]
	}
	var combinations = [][][]string{}
	GenerateCombinations(paths, &combinations)
	fmt.Println("\n-combination done-\n")
	if len(combinations) == 0 {
		fmt.Println(paths)
	}
	
	combinations = removeDuplicates(combinations)
	mapping := make([]int, len(combinations))

		for i := range combinations  { // step 1
			// fmt.Println("\033[34m", combinations[i], "\033[0m")
			
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
		// fmt.Println("::", len(mapping), "::", len(combinations))
		fmt.Println("--", mapping[minIndex], ": ", combinations[minIndex])
		g.walkIt(combinations[minIndex])

}

func (g *Graph) stepSimulator(paths [][]string) int {
	// fmt.Println("\n------------ Step Sim ------------")
	// for _, path := range paths {
	// 	fmt.Printf("Path through room: %v ==> Path:%v\n", path[1], path) // Temporary check for paths
	// }

	counter := g.makeQueue(paths)

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

func GenerateCombinations(paths[][]string, result *[][][]string){
	if len(paths) == 1 {
		*result  = append(*result, paths)
		return
	}
	firstPath := paths[0]
	combsWithoutFirst := [][][]string{}
	GenerateCombinations(paths[1:], &combsWithoutFirst)
	if combsWithoutFirst == nil {
		*result = append(*result, [][]string{firstPath})
		return
	}

	allCombsWithFirst := [][][]string{}
	for i := range combsWithoutFirst{
		allCombsWithFirst = append(allCombsWithFirst, append([][]string{firstPath}, combsWithoutFirst[i]...))
		allCombsWithFirst[len(allCombsWithFirst)-1] = disjointPaths(allCombsWithFirst[len(allCombsWithFirst)-1])
	}
	if len(allCombsWithFirst) > 0 {
		*result = append(*result, combsWithoutFirst...)
		*result = append(*result, allCombsWithFirst...)
	} else {
		*result = append(*result, combsWithoutFirst...)
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