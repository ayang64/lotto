package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

type SortedInt []int

func (si SortedInt) Len() int {
	return len(si)
}

func (si SortedInt) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}

func (si SortedInt) Less(i, j int) bool {
	return si[i] < si[j]

}

type LottoMachine struct {
	Min, Max int
	Ball     []int
}

func (lm *LottoMachine) Init() {
	for i := lm.Min; i <= lm.Max; i++ {
		lm.Ball = append(lm.Ball, i)
	}
}

func (lm *LottoMachine) PickNumbers(n int) (rc []int, err error) {
	for i := 0; i < n; i++ {
		v, err := lm.Pick()

		if err != nil {
			return rc, err
		}

		rc = append(rc, v)
	}
	return rc, err
}

func (lm *LottoMachine) Pick() (int, error) {
	if len(lm.Ball) == 0 {
		return 0, fmt.Errorf("Ball machine is empty.")
	}

	idx := rand.Intn(len(lm.Ball))

	rc := lm.Ball[idx]
	lm.Ball = append(lm.Ball[:idx], lm.Ball[idx+1:]...)
	return rc, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	draws := flag.Int("draws", 1, "Number of powerball sets to pick.")
	flag.Parse()

	for i := 0; i < *draws; i++ {

		PowerBall := LottoMachine{Min: 1, Max: 69}
		PowerPlay := LottoMachine{Min: 1, Max: 29}

		PowerBall.Init()
		PowerPlay.Init()

		pb, err := PowerBall.PickNumbers(5)
		if err != nil {
			log.Fatalf("Could not pick all power ball numbers: %s", err)
		}

		pp, err := PowerPlay.PickNumbers(1)
		if err != nil {
			log.Fatalf("Could not pick power play number: %s", err)
		}

		sn := SortedInt(pb)

		sort.Sort(sn)
		fmt.Printf("%v %v\n", sn, pp)
	}
}
