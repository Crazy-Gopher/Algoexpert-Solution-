package dijkstras_algorithm

import "math"

//[[1, 7]],
//[[2, 6], [3, 20], [4, 3]],
//[[3, 14]],
//[[4, 2]],
//[],
//[],

func dijkstrasAlgorithm(start int, edges [][][]int) []int {
	shortestPaths := make([]int, len(edges))
	for i := range edges {
		shortestPaths[i] = math.MaxInt32
	}
	shortestPaths[start] = 0
	dijkstrasHelper(shortestPaths, edges)

	return shortestPaths
}

func dijkstrasHelper(shortestPaths []int, edges [][][]int) {
	visited := make(map[int]struct{})

	for len(visited) < len(edges) {
		currNode, currShortest := findNodeWithMinWeight(shortestPaths, visited)

		if currShortest == math.MaxInt32 {
			for i, shortestPath := range shortestPaths {
				if shortestPath == math.MaxInt32 {
					currNode = i
					shortestPaths[currNode] = -1
				}
			}
			break
		}


		edge := edges[currNode]

		for _, pair := range edge {
			dist := pair[0]
			distance := pair[1]
			shortestPaths[dist] = min(shortestPaths[dist], distance+currShortest)
		}

		visited[currNode] = struct{}{}
	}
}

func findNodeWithMinWeight(shortestPaths []int, visited map[int]struct{}) (int, int) {
	shortest := math.MaxInt32
	var currNode int
	for i, shortestPath := range shortestPaths {
		if _, found := visited[i]; found {
			continue
		}

		if shortestPath < shortest {
			shortest = shortestPath
			currNode = i
		}
	}

	return currNode, shortest
}


func min(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num < curr {
			curr = num
		}
	}
	return curr
}
