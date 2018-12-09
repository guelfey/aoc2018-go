package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

type state struct {
	field  *ring.Ring
	player int
	scores []int64
}

func (s *state) placeMarble(marble int) {
	if marble%23 == 0 {
		s.scores[s.player] += int64(marble)
		for i := 0; i < 7; i++ {
			s.field = s.field.Prev()
		}
		remMarble := s.field.Value.(int)
		s.scores[s.player] += int64(remMarble)
		s.field = s.field.Prev()
		s.field.Link(s.field.Next().Next())
		s.field = s.field.Next()
	} else {
		s.field = s.field.Next()
		n := ring.New(1)
		n.Value = marble
		s.field.Link(n)
		s.field = s.field.Next()
	}
	//fmt.Println(s)
	s.nextPlayer()
}

func (s *state) nextPlayer() {
	s.player++
	if s.player == len(s.scores) {
		s.player = 0
	}
}

func (s *state) String() string {
	r := fmt.Sprintf("[%d] ", s.player+1)
	s.field.Do(func(v interface{}) {
		r += strconv.Itoa(v.(int)) + " "
	})
	return r
}

func sim(players, maxMarble int) int64 {
	s := state{
		field:  ring.New(1),
		player: 0,
		scores: make([]int64, players),
	}
	s.field.Value = 0
	for marble := 1; marble <= maxMarble; marble++ {
		s.placeMarble(marble)
	}
	maxScore := int64(0)
	for i := 0; i < players; i++ {
		if s.scores[i] > maxScore {
			maxScore = s.scores[i]
		}
	}
	return maxScore
}

func main() {
	fmt.Println(sim(418, 7076900))
}
