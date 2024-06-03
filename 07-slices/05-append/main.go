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
}

func getCostsByDay(costs []cost) []float64 {
	var totalCostsPerDay []cost
	maxDay := 0

	// Find the max day
	for i := 0; i < len(costs); i++ {
		// Find the day in totalCostsPerDay
		totalCostsIndex := slices.IndexFunc(totalCostsPerDay, func(totalCost cost) bool {
			return totalCost.day == costs[i].day
		})

		if totalCostsIndex != -1 {
			currCost := totalCostsPerDay[totalCostsIndex]
			newCost := currCost.value + costs[i].value

			modifyCost(&totalCostsPerDay[totalCostsIndex], newCost)
		} else {
			totalCostsPerDay = append(totalCostsPerDay, costs[i])
		}

		maxDay = int(math.Max(float64(costs[i].day), float64(maxDay)))
	}

	//* daily costs
	var dailyCosts []float64
	for i := 0; i <= maxDay; i++ {
		// Find the costs for the current day
		currentDayCostIdx := slices.IndexFunc(totalCostsPerDay, func(totalCost cost) bool {
			return totalCost.day == i
		})

		if currentDayCostIdx != -1 {
			dailyCosts = append(dailyCosts, totalCostsPerDay[currentDayCostIdx].value)
		} else {
			dailyCosts = append(dailyCosts, 0.0)
		}
	}

	return dailyCosts
}

func easyGetCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}
