package lasagna

import "fmt"

// TODO: define the 'OvenTime' constant

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	// panic("RemainingOvenTime not implemented")
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	// panic("PreparationTime not implemented")
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// panic("ElapsedTime not implemented")
	prepTime := PreparationTime(numberOfLayers)
	return prepTime + actualMinutesInOven
}

func main() {
	// Task 1: Define the expected oven time
	fmt.Printf("OvenTime: %d\n", OvenTime)

	// Task 2: Calculate the remaining oven time
	actualMinutesInOven := 30
	remainingTime := RemainingOvenTime(actualMinutesInOven)
	fmt.Printf("RemainingOvenTime: %d\n", remainingTime)

	// Task 3: Calculate the preparation time
	numberOfLayers := 2
	preparationTime := PreparationTime(numberOfLayers)
	fmt.Printf("PreparationTime: %d\n", preparationTime)

	// Task 4: Calculate the elapsed working time
	elapsedTime := ElapsedTime(3, 20)
	fmt.Printf("ElapsedTime: %d\n", elapsedTime)

}
