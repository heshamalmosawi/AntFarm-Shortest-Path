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

	combinations = removeDuplicates(combinations)
	mapping := make([]int, len(combinations))

	for i := range combinations {
		// simulating the number of steps needed then adding it to the map (weird name choice, i know)
		mapping[i] = g.stepSimulator(combinations[i])
	}

	// finding the path with the least number of steps needed.
	var minIndex int
	minValue := math.MaxInt64
	for index, value := range mapping {
		if value < minValue {
			minIndex = index
			minValue = value
		}
	}
	g.walkIt(combinations[minIndex])

}

/*
This function checks how many steps it would take for the farthest ant to cross and returns that number. the result is two more
than the real number but it doesnt make a difference in the comparisions to subtract 2 from all results.
*/
func (g *Graph) stepSimulator(paths [][]string) int {
	counter := g.makeQueue(paths)

	// finding how many steps it will take from maximum in counter (farthest ant)
	max := counter[0] + len(paths[0])
	for i, x := range counter {
		if max < x+len(paths[i]) {
			max = x
		}
	}
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

/* This function takes in a 2D array of paths, and appends all possible combinations to the result 3d array. */
func GenerateCombinations(paths [][]string, result *[][][]string) {
	if len(paths) == 1 {
		*result = append(*result, paths)
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
	for i := range combsWithoutFirst {
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

/*
This function takes in a 3d array, checks if any of the elements is repeated and returns the 3d array without them.
It is not necessary and probably useless but this function is implemented just in case there was any error made.
*/
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
