package main

import (
	"math"
	"slices"
)

type cost struct {
	day   int
	value float64
}

func modifyCost(c *cost, value float64) {
	c.value = value
	return
}

func getCostsByDay(costs []cost) []float64 {
	var totalCostsPerDay []cost

    // Sort the costs
    

	// Find the max day
	for i := 0; i < len(costs); i++ {
		// Find the day in totalCostsPerDay
		totalCostsIndex := slices.IndexFunc(totalCostsPerDay, func(tC cost) bool {
			return tC.day == costs[i].day
		})

		if totalCostsIndex != -1 {
			currCost := totalCostsPerDay[totalCostsIndex]

			modifyCost(&currCost, currCost.value+costs[i].value)
		} else if tot
	}

	// Create the array

	return nil
}
