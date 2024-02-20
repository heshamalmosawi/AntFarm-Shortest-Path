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

/* This function finds the disjoint paths, sorry for it being ineffecient and confusing 😂 */
func disjointPaths(paths [][]string) [][]string {
	var firstPath [][]string = [][]string{paths[0]} // saving for back up, you'll see why at the end of the function.
	roomsUsed := make(map[string]string, 0) // Map to store the rooms used so far, for effeciency
	for path1index, path := range paths{
		for _, room := range path{
			if room == Farm.startRoom || room == Farm.endRoom { // ignore start and end rooms
				continue
			}

			// The searching part of this function. If room already found or just found now, delete it from the entire path and break
			for path2index, path2 := range paths {
				if path1index == path2index {
					continue
				}
				for _, room2 := range path2{
					if room2 == Farm.startRoom || room2 == Farm.endRoom { // ignore start and end rooms
						continue
					}
					if _, ok := roomsUsed[room2]; ok || room == room2 {
						fmt.Println(room, " ", room2, "\tmap:", roomsUsed)
						roomsUsed[room2] = ""
						if path2index <= len(paths)-1{
							paths = append(paths[:path2index], paths[path2index+1:]...)
						} else { // this is for not going out of bounds with the +1
							paths = paths[:path2index]
						}
						break
					} 
				}
			}
			// adding room to the map 
			roomsUsed[room] = ""
		}
	}
	fmt.Print(roomsUsed)
	// if the length of paths is 0, then just return the first path, this is done when all paths intersect, or just in case of anything going wrong.
	if len(paths) == 0 {
		return firstPath
	}
	return paths
}