package main

import "fmt"

// Your savings function here
func savings(grossPay int, taxRate float64, expenses int) int {
	tax := float64(grossPay) * taxRate
	afterTaxPay := float64(grossPay) - tax
	centavosRemaining := afterTaxPay - float64(expenses)
	return int(centavosRemaining)
}

// Material waste
func materialWaste(totalMaterial int, materialUnits string, numJobs int, jobConsumption int) string {
	totalConsumed := numJobs * jobConsumption
	remainingMaterial := totalMaterial - totalConsumed
	return fmt.Sprintf("%d%s", remainingMaterial, materialUnits)
}

// Interest
func interest(principal int, rate float64, periods int) int {
	interest := float64(principal) * rate * float64(periods)
	finalInvestmentValue := float64(principal) + interest
	return int(finalInvestmentValue)
}

func main() {
	// Example usage
	fmt.Println(savings(5000, 0.1, 500))
	fmt.Println(materialWaste(1000, "units", 10, 5))
	fmt.Println(interest(1000, .05, 2))
}
