package config

type Config struct {
	DateFormat        string     `yaml:"date_format"` // Getting the date format
	DataSource        DataSource `yaml:"data_source"`// Getting the sources of Data
	// Base
	Base              map[int]float64
	MaxBase           int
	MinBase           int
	// Drivers Age
	Age               map[int]float64
	MaxAge            int
	MinAge            int
	// License Length
	LicenseLength     map[int]float64
	MaxLicenseLength  int
	MinLicenseLength  int
	// Insurance Group
	InsuranceGroup    map[int]float64
	MaxInsuranceGroup int
	MinInsuranceGroup int
}

// We store the paths corresponding the the individual csv
type DataSource struct {
	BaseCSV           string `yaml:"base_csv"`
	AgeCSV            string `yaml:"age_factor_csv"`
	LicenseLengthCSV  string `yaml:"license_length_csv"`
	InsuranceGroupCSV string `yaml:"insurance_group_csv"`
}