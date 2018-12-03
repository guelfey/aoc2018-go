package main

import "fmt"
import "io"

type Claim struct {
	id int
	x, y, w, h int
}

func (c *Claim) mark(field [][]int, overlapping map[int]bool) {
	for x := c.x; x < c.x+c.w; x++ {
		for y := c.y; y < c.y+c.h; y++ {
			if field[x][y] == 0 {
				field[x][y] = c.id
			} else {
				overlapping[c.id] = true
				overlapping[field[x][y]] = true
				field[x][y] = -1
			}
		}
	}
}

func readClaims() []Claim {
	claims := make([]Claim, 0)
	for {
		c := Claim{}
		_, err := fmt.Scanf("#%d @ %d,%d: %dx%d\n", &c.id, &c.x, &c.y, &c.w, &c.h)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			} else {
				panic(err)
			}
		}
		claims = append(claims, c)
	}
	return claims
}

const fieldWidth = 1000
const fieldHeight = 1000

func main() {
	claims := readClaims()

	field := make([][]int, fieldWidth)
	for i := 0; i < fieldWidth; i++ {
		field[i] = make([]int, fieldHeight)
	}

	overlapping := make(map[int]bool)
	for _, c := range claims {
		c.mark(field, overlapping)
	}

	doubleClaimed := 0
	for x := 0; x < fieldWidth; x++ {
		for y := 0; y < fieldWidth; y++ {
			if field[x][y] == -1 {
				doubleClaimed++
			}
		}
	}
	fmt.Println(doubleClaimed)

	for i := 1; i <= len(claims); i++ {
		if !overlapping[i] {
			fmt.Println(i)
		}
	}
}
