package app

import (
	"pricingengine"
	"pricingengine/config"
	"pricingengine/utils"
)

// We have these attributes which will help us
// in creating mocks for unit testing
type App struct {
	Config     config.IConfig
	DateHelper utils.DateHelper
}

type BaseCosts []pricingengine.BaseCost