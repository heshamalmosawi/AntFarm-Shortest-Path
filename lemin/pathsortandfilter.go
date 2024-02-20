package lemin

import (
	"fmt"
)

/* This function performs the Quick Sort Algorithm on the two-dimensional array consisting of paths,
and sorts them in ascending order based on path length. */
func QuickSort(paths [][]string) [][]string {
	// stopping condition for the end of the recursion, if the length is 0 or 1, it will return itself.
	if len(paths) <= 1 { 
		return paths
	}

	// calling the partition function that would put the pivot in its place and return its index for partitioning/dividing the recursion input
	pivot := partition(paths) 
	// partitioning/dividing the recursion input arrays
	lower := QuickSort(paths[:pivot])
	higher := QuickSort(paths[pivot+1:])
	// var answer [][]string
	// answer = append(answer, lower...)
	// answer = append(answer, paths[pivot])
	// answer = append(answer, higher...)
	// return answer
	return append(append(lower, paths[pivot]), higher...) // combining all in one line for shorter code
}

/* This function selects a pivot element, places the elements with lower length than the pivot on the left side of the pivot and 
the elements with bigger length on the right side.*/
func partition(paths [][]string) int {
	pivot := len(paths[0]) // choosing first element as the 'pivot'
	i, j := 1, len(paths)-1 // saving the index of the first and last element to use as "pointers" (not really pointers)

	for i <= j {  // will continue until the first and second point (i, j) cross paths in the array, where it will be i>j
		for i <= j && len(paths[i]) <= pivot { 
			// Iterate over all paths that their length are smaller or equal to the pivot, then it will stop when pivot > length of path
			i++
		}
		for i <= j && len(paths[j]) >= pivot {
			// Iterate over all paths with length larger or equal to the pivot, then stop when pivot < path length.
			j--
		}
		if i <= j { // if i & j cross paths, swap their contents 
			paths[i], paths[j] = paths[j], paths[i]
		}
	}
	// swap the pivot to its place (currently index j) 
	paths[0], paths[j] = paths[j], paths[0]
	return j // return index j
}

/* This function  */
func disjointPaths(paths [][]string) [][]string {
	var firstPath [][]string = [][]string{paths[0]}
	roomsUsed := make(map[string]string, 0) //
	for _, path := range paths{
		for _, room := range path{
			if room == Farm.startRoom || room == Farm.endRoom { // ignore start and end rooms
				continue
			}
			for path2index, path2 := range paths {
				foundIntersection := false
				for _, room2 := range path2{
					if _, ok := roomsUsed[room2]; ok || room == room2 {
						// fmt.Print(pathIndex, "\n", room, " ",room2, "\n")
						foundIntersection = true
						break
					} 
				}
				if foundIntersection {
					// fmt.Print("\n\nIteration:" , iteration, "\nPath2:", path2, "\nroomsUsed:", roomsUsed, "\n")
					if path2index < len(paths)-1{
						paths = append(paths[:path2index], paths[path2index+1:]...)
					} else {
						paths = paths[:path2index]
						// disjointPaths = append(disjointPaths, path2)
					}
				}
			roomsUsed[room] = ""
			}
		}
	}
	fmt.Print(roomsUsed)
	if len(paths) == 0 {
		return firstPath
	}
	return paths
}