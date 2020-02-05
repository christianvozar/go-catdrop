// Package catdrop utilizes recursion to evaluate the minimum floor to drop
// a cat with survival given a number of cats dropped from a given building
// size.
//
// It is also possible to evaluate the solution with a Binary Search Tree but
// is omitted here.
//
// Example usage: catdrop.Solve(10000, 10)
package catdrop

import "math"

var (
	// TotalCats is the number of cats in given run.
	// A cat that lives from a fall is reusable.
	TotalCats int
	// TotalFloors is the number of floors in a given build in a given run.
	TotalFloors int
	// TotalCatDrops represents the total number of cat drops.
	TotalCatDrops int
)

// Solve takes a number of floors and cats and returns the minimum floor to
// drop a cat with surival.
// On error return 0 per Go standards.
func Solve(floors, cats int) int {
	var (
		totalDropsCatDies  int
		totalDropsCatLives int
		maxDrops           int
		minDrops           int
	)

	// Base scenario: If no floors or only one floor return.
	if floors == 0 || floors == 1 {
		return floors
	}

	// Base scenario: If only one cat we must try from all floors starting at one.
	if cats == 1 {
		return floors
	}

	minDrops = math.MaxInt32

	// Drop cat from floor 1 to given floors
	for i := 1; i <= floors; i++ {
		// Cat dies at i-th floor
		totalDropsCatDies = Solve(cats-1, i-1)

		// Cat lives at i-th floor
		totalDropsCatLives = Solve(floors-i, cats)

		// Determine worst case scenario
		maxDrops = max(totalDropsCatDies, totalDropsCatLives)

		// Determine minimum drops for all floors.
		if minDrops > maxDrops {
			minDrops = maxDrops
		}
	}

	return minDrops
}

// max returns the larger of two given values.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
