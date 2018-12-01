package main

import (
	"flag"
	"fmt"
	"io"
)

func doFirst() int {
	freq := 0
	for {
		change := 0
		_, err := fmt.Scanln(&change)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		freq += change
	}
	return freq
}

func doSecond() int {
	changes := make([]int, 0)
	for {
		change := 0
		_, err := fmt.Scanln(&change)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		changes = append(changes, change)
	}
	if len(changes) == 0 {
		panic("no changes read")
	}
	freq := 0
	reached := make(map[int]bool)
	i := 0
	for {
		if reached[freq] {
			return freq
		}
		reached[freq] = true
		freq += changes[i]
		i++
		if i == len(changes) {
			i = 0
		}
	}
}

func main() {
	secondPart := flag.Bool("second", false, "solve second part")
	flag.Parse()

	res := 0
	if *secondPart {
		res = doSecond()
	} else {
		res = doFirst()
	}
	fmt.Println(res)
}
