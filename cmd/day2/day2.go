package main

import "flag"
import "fmt"
import "io"

func doFirst() {
	two := 0
	three := 0
	for {
		s := ""
		_, err := fmt.Scanln(&s)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		counts := make(map[rune]int)
		for _, r := range s {
			counts[r]++
		}

		for _, count := range counts {
			if count == 2 {
				two++
				break
			}
		}
		for _, count := range counts {
			if count == 3 {
				three++
				break
			}
		}
	}
	fmt.Println(two*three)
}

func diffPos(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	p := -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if p != -1 {
				return -1
			}
			p = i
		}
	}
	return p
}

func doSecond() {
	ids := make([]string, 0)
	for {
		id := ""
		_, err := fmt.Scanln(&id)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		ids = append(ids, id)
	}
	for _, id1 := range ids {
		for _, id2 := range ids {
			if p := diffPos(id1, id2); p != -1 {
				fmt.Printf("%s%s\n", id1[:p], id1[p+1:])
				return
			}
		}
	}
}

func main() {
	secondPart := flag.Bool("second", false, "solve second part")
	flag.Parse()

	if *secondPart {
		doSecond()
	} else {
		doFirst()
	}
}
