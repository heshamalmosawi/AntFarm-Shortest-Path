package lemin

import (
	"fmt"
	"strconv"
)

/* This method is the method that calls for all walking-printing functions. */
func (g *Graph) walkIt(paths [][]string) {
	queue := g.makeQueue(paths)
	// Creating antsWithNum to have all ants numbered/labeled.
	antsWithNum := make(map[string]string, g.ants)
	for i := 1; i <= g.ants; i++ {
		antName := "L" + strconv.Itoa(i)
		antsWithNum[antName] = ""
	}

	// the walking loop
	antsRemaining := g.ants
	for antsRemaining > 0 {
		antsRemaining = g.stepForward(paths, queue, antsWithNum, antsRemaining)

	}
	fmt.Println()
}

/*
This method simulates the ants stepping a step forward to the next room and prints each step.
It starts from the end and iterates backward to the start so it can free up room to other ants.
*/
func (g Graph) stepForward(paths [][]string, queue []int, antsWithNum map[string]string, antsRemaining int) int {

	needforNewLine := false // for convenience

	// First loop: for ants reaching the end room
	for i := range paths {
		node, _ := g.GetVertex(paths[i][len(paths[i])-2])
		if node.antInRoom[0] != "" { // if second to last room is not empty
			fmt.Print(node.antInRoom[0], "-", paths[i][len(paths[i])-1], " ")
			antsWithNum[node.antInRoom[0]] = "end"
			node.antInRoom[0] = "" // freeing up the room
			needforNewLine = true
		}
	}

	// Second loop: for ants in the middle of the path
	for i := range paths {
		for j := len(paths[i]) - 1; j > 0; j-- {
			node, _ := g.GetVertex(paths[i][j])
			previous, _ := g.GetVertex(paths[i][j-1])
			if node.antInRoom[0] == "" && previous.antInRoom[0] != "" { // if room is empty but previous is not
				node.antInRoom[0] = previous.antInRoom[0]
				previous.antInRoom[0] = "" // freeing up the old room
				fmt.Print(node.antInRoom[0], "-", node.key, " ")
			}
		}
	}

	// Third loop: for ants in queue and waiting to move out
	for i := range paths {
		node, _ := g.GetVertex(paths[i][1])
		if node.antInRoom[0] == "" && queue[i] > 0 {
			temp := findNextFreeWaitingAnt(antsWithNum)
			if temp == "" {
				break
			}
			node.antInRoom[0] = temp
			antsWithNum[temp] = node.antInRoom[0]
			fmt.Print(node.antInRoom[0], "-", node.key, " ")
			antsRemaining--
			queue[i]--
			if paths[i][1] == g.endRoom { // if the first room is also the end room, then we need to free it up immediately
				antsWithNum[node.antInRoom[0]] = "end"
				node.antInRoom[0] = "" // freeing up the room
			}
			needforNewLine = true
		}
	}
	// reduceQueueNum(&queue)
	if needforNewLine {
		fmt.Println()
	}
	return antsRemaining
}

/* This function finds an ant that is waiting for its chance to step in. If none found it will return an empty string. */
func findNextFreeWaitingAnt(antsWithNum map[string]string) string {

	for i := 1; i <= len(antsWithNum); i++ {
		lnum := fmt.Sprintf("L%d", i)
		// fmt.Println(lnum)
		if antsWithNum[lnum] == "" {
			return lnum
		}
	}

	return ""
}
