package main

import (
	"fmt"
	"io"
	"math"
	"sort"
)

type Graph struct {
	nodes    []rune
	outEdges map[rune][]rune
	inEdges  map[rune][]rune
}

func readGraph() Graph {
	g := Graph{
		outEdges: make(map[rune][]rune),
		inEdges:  make(map[rune][]rune),
	}
	for {
		var from, to rune
		_, err := fmt.Scanf("Step %c must be finished before step %c can begin.\n", &from, &to)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				panic(err)
			} else {
				break
			}
		}
		g.outEdges[from] = append(g.outEdges[from], to)
		g.inEdges[to] = append(g.inEdges[to], from)
	}
	for n := range g.outEdges {
		g.nodes = append(g.nodes, n)
	}
	for n := range g.inEdges {
		if _, ok := g.outEdges[n]; !ok {
			g.nodes = append(g.nodes, n)
		}
	}
	return g
}

func getReady(s map[rune]bool) rune {
	slice := []rune{}
	for k := range s {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	return slice[0]
}

func isReady(t rune, g Graph, order []rune) bool {
	for _, req := range g.inEdges[t] {
		found := false
		for _, done := range order {
			if req == done {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func topSort(g Graph) []rune {
	ready := make(map[rune]bool)
	for _, n := range g.nodes {
		if len(g.inEdges[n]) == 0 {
			ready[n] = true
		}
	}
	order := []rune{}
	for len(ready) != 0 {
		n := getReady(ready)
		delete(ready, n)
		order = append(order, n)

		for _, after := range g.outEdges[n] {
			if isReady(after, g, order) {
				ready[after] = true
			}
		}
	}
	return order
}

// second part
const workers = 5

func nextFreeWorker(workerReady []int) int {
	minTime := workerReady[0]
	minWorker := 0
	for i, t := range workerReady {
		if t < minTime {
			minWorker = i
			minTime = t
		}
	}
	return minWorker
}

func taskTime(task rune) int {
	return int(task-'A') + 61
}

func taskBegin(task rune, g Graph, finTime map[rune]int) int {
	t := 0
	for _, req := range g.inEdges[task] {
		if _, ok := finTime[req]; !ok {
			return -1
		}
		if finTime[req] > t {
			t = finTime[req]
		}
	}
	return t
}

func nextTask(g Graph, finTime map[rune]int) rune {
	earliestTask := rune(0)
	earliestTime := math.MaxInt32
	for _, n := range g.nodes {
		if _, ok := finTime[n]; ok {
			continue
		}
		t := taskBegin(n, g, finTime)
		if t == -1 {
			continue
		}
		if t < earliestTime ||
			(t == earliestTime && n < earliestTask) {
			earliestTask = n
			earliestTime = t
		}
	}
	return earliestTask

}

func totalTime(g Graph) int {
	finTime := make(map[rune]int)
	workerReady := make([]int, workers)
	for task := nextTask(g, finTime); task != 0; task = nextTask(g, finTime) {
		w := nextFreeWorker(workerReady)
		startTime := taskBegin(task, g, finTime)
		if workerReady[w] > startTime {
			startTime = workerReady[w]
		}
		endTime := startTime + taskTime(task)
		fmt.Println("task", string(task), "worker", w, "start", startTime, "end", endTime)
		workerReady[w] = endTime
		finTime[task] = endTime
	}
	endTime := 0
	for _, t := range finTime {
		if t > endTime {
			endTime = t
		}
	}
	return endTime
}

func main() {
	g := readGraph()
	order := topSort(g)
	fmt.Println(string(order))
	fmt.Println(totalTime(g))
}
