package backtracking

import "fmt"

func Hamiltoniancycle(graph [][]int) {
	hamiltoniancycle(0, []int{}, graph)
}

func hamiltoniancycle(v int, path []int, graph [][]int) {
	path = append(path, v)

	if len(path) == 2 {
		return
	}

	for i, val := range graph[v] {
		exists := false
		for _, e := range path {
			if i == e {
				exists = true
			}
		}

		if val == 1 && !exists {
			newPath := []int{}
			copy(newPath, path)
			fmt.Println("the vertex is ", i)
			fmt.Println(path)
			hamiltoniancycle(i, newPath, graph)
		}
	}

	// fmt.Println(path)
}

func DFS(graph [][]int) {
	path := []int{}
	dfs(0, graph, path)
}

func dfs(v int, graph [][]int, path []int) {
	path = append(path, v)

	recursed := false
	for i, val := range graph[v] {
		e := false
		for _, ex := range path {
			if ex == i {
				e = true
			}
		}
		if val == 1 && !e {
			npath := make([]int, len(path))
			copy(npath, path)
			recursed = true
			dfs(i, graph, npath)
		}
	}

	if !recursed {
		fmt.Println(path)
	}
}
