package config

// GetConfig function reads the base configuration from the yaml
// and all the csvs
func (c *Config) SetConfig() (err error) {
	// Reading application configs from yaml file
	err = c.ReadFromYaml()
	if err != nil {
		return err
	}
	// Reading the CSV for Base Rate
	c.Base, c.MaxBase, c.MinBase, err = c.ReadFromCSV(c.DataSource.BaseCSV)
	if err != nil {
		return err
	}
	// Reading the CSV for Drivers Age Factors
	c.Age, c.MaxAge, c.MinAge, err = c.ReadFromCSV(c.DataSource.AgeCSV)
	if err != nil {
		return err
	}
	// Reading the CSV for License Length Factors
	c.LicenseLength, c.MaxLicenseLength, c.MinLicenseLength, err = c.ReadFromCSV(c.DataSource.LicenseLengthCSV)
	if err != nil {
		return err
	}
	// Reading the CSV for Insurance Group Factors
	c.InsuranceGroup, c.MaxInsuranceGroup, c.MinInsuranceGroup, err = c.ReadFromCSV(c.DataSource.InsuranceGroupCSV)
	if err != nil {
		return err
	}
	return
}

func (c *Config) GetConfig() (Config) {
	return *c
}
