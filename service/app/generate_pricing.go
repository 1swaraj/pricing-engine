package app

import (
	"context"
	"fmt"
	"pricingengine"
	"sort"
	"time"
)

// GeneratePricing will calculate how much a 'risk' be priced or if they should
// be denied.
func (a *App) GeneratePricing(ctx context.Context, input *pricingengine.GeneratePricingRequest) (*pricingengine.GeneratePricingResponse, error) {
	// Getting the config in a variable
	config := a.Config.GetConfig()
	licenseHeldSince, err := time.Parse(config.DateFormat, input.LicenseHeldSince)
	if err != nil {
		return nil, err
	}
	dateOfBirth, err := time.Parse(config.DateFormat, input.DateOfBirth)
	if err != nil {
		return nil, err
	}

	// The license can't be acquired before the date of birth
	if licenseHeldSince.Sub(dateOfBirth).Seconds() <= 0 {
		return nil, fmt.Errorf(pricingengine.LicenseBeforeBirth)
	}
	// License can't be acquired before current date
	// Because license is already greater than the date of birth
	// so we don't need to check the date of birth
	if licenseHeldSince.Sub(time.Now()).Seconds() > 0 {
		return nil, fmt.Errorf(pricingengine.DateAheadToday)
	}

	// Getting the age of the applicant / user from the Date of Birth
	age, err := a.DateHelper.AgeFromDate(dateOfBirth)
	if err != nil {
		return nil, err
	}

	// Getting the number of years the applicant / user is a license holder
	licenseHeldYears, err := a.DateHelper.AgeFromDate(licenseHeldSince)
	if err != nil {
		return nil, err
	}

	// Checking if the age is not less than the minimum age allowed
	if age < config.MinAge {
		return nil, fmt.Errorf(pricingengine.Declined)
	}

	// Getting the driver's age factor
	driverAgeFactor, err := a.GetFactor(age, config.MaxAge, config.Age)
	if err != nil {
		return nil, err
	}

	// Getting the insurance group factor
	insuranceFactor, err := a.GetFactor(input.InsuranceGroup, config.MaxInsuranceGroup, config.InsuranceGroup)
	if err != nil {
		return nil, err
	}
	// Getting the license length factor
	licenseLength, err := a.GetFactor(licenseHeldYears, config.MaxLicenseLength, config.LicenseLength)
	if err != nil {
		return nil, err
	}

	// Generating the BaseCost that will be given in the User Response
	var baseCost BaseCosts
	// BaseCostConfig is the data provided in the CSV
	baseCostConfig := config.Base
	// Iterating through the baseCostConfig and finding and appending the baseCost for the
	// given conditions - drivers age factor, insurance group factor and license length factor
	for sec, cost := range baseCostConfig {
		i := sort.Search(len(baseCost), func(i int) bool { return baseCost[i].TimeSec >= sec })
		if i < len(baseCost) && baseCost[i].TimeSec == sec {
			baseCost[i]=pricingengine.BaseCost{TimeSec: sec, TimeDays: a.DateHelper.Days(sec), TimeHours: a.DateHelper.Hours(sec), BaseRate: cost * licenseLength * driverAgeFactor * insuranceFactor}
		} else {
			baseCost = append(baseCost, pricingengine.BaseCost{})
			copy(baseCost[i+1:], baseCost[i:])
			baseCost[i]=pricingengine.BaseCost{TimeSec: sec, TimeDays: a.DateHelper.Days(sec), TimeHours: a.DateHelper.Hours(sec), BaseRate: cost * licenseLength * driverAgeFactor * insuranceFactor}
		}
	}
	
	// baseCost is now an array that shows the price in pence for the given factors
	return &pricingengine.GeneratePricingResponse{BaseCosts: baseCost, LicenseLength: licenseLength, DriverAgeFactor: driverAgeFactor, InsuranceGroup: insuranceFactor}, nil
}

// Init Config initializes the configs
// It is noteworthy that we are calling the a.config.GetConfig() method
// this method shall change the state of the App struct and thus the changes
// are reflected globally
func (a *App) InitConfig() (err error) {
	err = a.Config.SetConfig()
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Couldn't read service config")
	}
	return
}

