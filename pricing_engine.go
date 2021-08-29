package pricingengine

// GeneratePricingRequest is used for generate pricing requests, it holds the
// inputs that are used to provide pricing for a given user.
type GeneratePricingRequest struct {
	DateOfBirth      string `json:"date_of_birth"`
	InsuranceGroup   int    `json:"insurance_group"`
	LicenseHeldSince string `json:"license_held_since"`
}

// GeneratePricingResponse
type GeneratePricingResponse struct {
	BaseCosts       []BaseCost `json:"base_cost"`
	DriverAgeFactor float64    `json:"driver_age_factor"`
	InsuranceGroup  float64    `json:"insurance_group_factor"`
	LicenseLength  float64    `json:"license_length_factor"`
}

// This is the most important struct in our response
// The base rate is in pence
type BaseCost struct {
	TimeSec   int     `yaml:"time_sec"`
	TimeHours int     `yaml:"time_hours"`
	TimeDays  int     `yaml:"time_days"`
	BaseRate  float64 `yaml:"base_rate"`
}