package main

import "math/rand"


func Rnd(min int, max int) int {
	return rand.Intn(max - min + 1) + min
}

func ComputeComplexity(score3Shot bool) int {
	if score3Shot {
		return Rnd(Rnd(78, 88), 100)
	}

	return Rnd(Rnd(60, 70), 100)
}