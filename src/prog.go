package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type IntMatrix [][]int
type BoolMatrix [][]bool

type StrategyMatrix []func([]bool, []bool, int) bool

func simulate(sm StrategyMatrix, steps int) {
	strategy := map[int]string{
		0: "random",
		1: "titForTat",
		2: "alwaysTrue",
		3: "alwaysFalse",
		4: "forgivingTitForTat",
		5: "strategyChecker",
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(sm); i++ {
		for j := 0; j < len(sm); j++ {
			// if i-j == 1 {
			// 	continue
			// }
			points1 := 0
			points2 := 0
			choices1 := make([]bool, steps)
			choices2 := make([]bool, steps)

			for s := 0; s < steps; s++ {
				out1 := sm[i](choices1, choices2, s)
				out2 := sm[j](choices2, choices1, s)
				choices1[s] = out1
				choices2[s] = out2

				oldpoints1 := points1
				oldpoints2 := points2

				switch {
				case out1 == false && out2 == false:
					points1 += 1
					points2 += 1
				case out1 == false && out2 == true:
					points1 += 3
					points2 += 0
				case out1 == true && out2 == false:
					points1 += 0
					points2 += 3
				case out1 == true && out2 == true:
					points1 += 2
					points2 += 2
				}

				fmt.Println(strategy[i] + "," + strategy[j] + "," + strconv.Itoa(s) + "," + strconv.FormatBool(out1) + "," + strconv.FormatBool(out2) + "," + strconv.Itoa(points1-oldpoints1) + "," + strconv.Itoa(points2-oldpoints2) + "," + strconv.Itoa(points1) + "," + strconv.Itoa(points2))

			}
		}
	}
}

func main() {

	fmt.Println("firstAlg,secondAlg,step,firstOut,secondOut,firstGain,secondGain,firstPoints,secondPoints")

	alwaysTrue := func(m []bool, n []bool, s int) bool {
		return true
	}
	alwaysFalse := func(m []bool, o []bool, s int) bool {
		return false
	}
	titForTat := func(m []bool, o []bool, s int) bool {
		if s == 0 {
			return true
		} else {
			return o[s-1]
		}
	}
	// maliciousTitForTat := func(m []bool, o []bool, s int) bool {
	// 	if s == 0 {
	// 		return false
	// 	} else {
	// 		return o[s-1]
	// 	}
	// }
	random := func(m []bool, o []bool, s int) bool {
		return rand.Intn(2) == 0
	}
	// startWithTrue := func(m []bool, o []bool, s int) bool {
	// 	for i := 0; i < s; i++ {
	// 		if o[i] == false {
	// 			return false
	// 		}
	// 	}
	// 	return true
	// }
	// startWithFalse := func(m []bool, o []bool, s int) bool {
	// 	for i := 0; i < s; i++ {
	// 		if o[i] == true {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }
	forgivingTitForTat := func(m []bool, o []bool, s int) bool {
		if s == 0 {
			return true
		}
		if s == 1 {
			if o[0] == false {
				return false
			}
			return true
		}
		return o[s-1] || o[s-2]
	}
	// unforgivingTitForTat := func(m []bool, o []bool, s int) bool {
	// 	if s == 0 {
	// 		return true
	// 	}
	// 	if s == 1 {
	// 		if o[0] == false {
	// 			return false
	// 		}
	// 		return true
	// 	}
	// 	return o[s-1] && o[s-2]
	// }
	// slowTitForTat := func(m []bool, o []bool, s int) bool {
	// 	if s == 0 {
	// 		return true
	// 	}
	// 	if s == 1 {
	// 		return true
	// 	}
	// 	return o[s-1] || o[s-2]
	// }
	strategyChecker := func(m []bool, o []bool, s int) bool {
		if s == 0 {
			return false
		}
		if s == 1 {
			return false
		}
		if s == 2 {
			return true
		}
		if s == 3 {
			return false
		}
		if s == 4 {
			return true
		}
		points := 0
		for i := s - 5; i < s; i++ {
			if o[s-i] {
				points++
			}
		}
		if points == 0 {
			return false
		}
		if points == 1 {
			return o[s-1] || o[s-2]
		}
		if points == 2 {
			return o[s-1] || o[s-2]
		}
		if points == 3 {
			return !o[s-1]
		}
		if points == 4 {
			return !o[s-1]
		}
		if points == 5 {
			return false
		}
		return true
	}

	sm := StrategyMatrix{
		random,
		titForTat,
		alwaysTrue,
		alwaysFalse,
		// maliciousTitForTat,
		// startWithTrue,
		// startWithFalse,
		forgivingTitForTat,
		// unforgivingTitForTat,
		// slowTitForTat,
		strategyChecker,
	}

	simulate(sm, 100)

}
