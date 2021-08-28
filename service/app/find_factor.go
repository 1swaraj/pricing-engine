package app

import (
	"fmt"
	"pricingengine"
)

// Writing a Generic Method that handles finding the factor for
// 1. Driver's Age
// 2. Insurance Group
// 3. License Length
func (a *App) GetFactor(input, max int, factors map[int]float64) (factor float64, err error) {
	fmt.Printf("%d %d %v\n",input,max,factors)
	// If the input value exists in the map
	// i.e. it was present in the csv file
	// then we just check if the factor is not -1
	// as -1 stands for denied
	if val, ok := factors[input]; ok {
		if val == -1 {
			return 0, fmt.Errorf(pricingengine.Declined)
		}
		return val, nil
	}

	// For any inputs greater than the maximum value in the map
	// we return the maximum factor. We also check that the maximum factor
	// is not -1 (as -1 stands for denied).
	if input > max && factors[max] != -1 {
		return factors[max], nil
	}
	return 0, fmt.Errorf(pricingengine.Declined)
}
