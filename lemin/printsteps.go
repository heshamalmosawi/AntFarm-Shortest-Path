package lemin

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func (g *Graph) walkIt(paths [][]string) {
	queue := g.makeQueue(paths)
	antsWithNum := make(map[string]string, g.ants)
	for i := 0; i < g.ants; i++ {
		antName := "L" + strconv.Itoa(i)
		antsWithNum[antName] = ""
	}
	fmt.Println(antsWithNum)

	antsRemaining := g.ants
	// numberOfPaths := len(paths)
    // fmt.Println("\033[1;32m" , queue , "\033[0m")
	for antsRemaining >= 0 {
		g.stepForward(paths, queue, antsWithNum)
		// g.printsteps(paths)
		antsRemaining--
	}

}

/* This function reduces the number of ants from the queue, this is to be used after ants take a step forward. */
func reduceQueueNum(queue *[]int) {
	for i := range *queue {
		if (*queue)[i] > 0 {
			(*queue)[i]--
		}
	}
	// fmt.Print("\033[1;32m" , *queue , "\033[0m")

}

func (g Graph) stepForward(paths [][]string, queue []int, antsWithNum map[string]string) {

	// for i := range paths {
	// 	for j := range paths[i]{
	// 	fmt.Println(j, paths[i][j])
	// }
	// fmt.Print(len(paths[i]))
	// os.Exit(1)

	// }
	for i := range paths {
		// if paths[i][4] == "k" {
		// 	fmt.Println("YEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
		// }
		node, _ := g.GetVertex(paths[i][len(paths[i])-2])
		if node.antInRoom[0] != "" {
			fmt.Print("\033[1;32m", node.antInRoom[0], "-", paths[i][len(paths[i])-1], " \033[0m")
			node.antInRoom[0] = ""
		}
	}
	for i := range paths{
		for j := len(paths[i])-1 ; j > 0; j--{
			node, err := g.GetVertex(paths[i][j])
			if err != nil{
				log.Fatal("hello world")
			}
			previous, err := g.GetVertex(paths[i][j-1])
			if err != nil{
				log.Fatal("hello world2")
			}
			if	node.antInRoom[0] == ""  && previous.antInRoom[0] != "" {
				node.antInRoom[0] = previous.antInRoom[0]
				previous.antInRoom[0] = ""
				fmt.Print("\033[1;34m", node.antInRoom[0], "-", node.key, " \033[0m")
			}
		}
	}

	
	for i := range paths{
		node, _ := g.GetVertex(paths[i][1])
		if node.antInRoom[0] == "" && queue[i] > 0 {
			temp := findNextFreeWaitingAnt(antsWithNum)
			if temp == ""{
				break
			}
			node.antInRoom[0] = temp
			// thids := fmt.Sprint("L", i)
			// fmt.Print("\033[1;31m",thids,"\033[0m")
			antsWithNum[temp] = node.antInRoom[0] 
			fmt.Print("\033[1;31m", node.antInRoom[0], "-", node.key, " \033[0m")
		} else if queue[i]<0{
			fmt.Println(queue)
			os.Exit(1)
		}
	}
	reduceQueueNum(&queue)
	fmt.Println()
}

func findNextFreeWaitingAnt(antsWithNum map[string]string) string{
	for key, room := range antsWithNum{
		if room == ""{
			return key 
		}
	} 
	log.Fatal("WAAA")
	return ""
}