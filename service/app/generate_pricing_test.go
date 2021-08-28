package app

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pricingengine"
	"pricingengine/config"
	"testing"
)

type IConfigMock struct {
	config.IConfig
	mock.Mock
}

func (conf *IConfigMock) SetConfig() error {
	args := conf.Called()
	return args.Error(0)
}

func (conf *IConfigMock) GetConfig() config.Config {
	args := conf.Called()
	return args.Get(0).(config.Config)
}

func TestApp_InitConfig(t *testing.T) {
	confMock := new(IConfigMock)
	confMock.On("SetConfig").Return(nil)
	app := App{Config: confMock}
	assert.Nil(t, app.InitConfig())
}

func TestApp_GeneratePricing(t *testing.T) {
	confMock := new(IConfigMock)
	c := config.Config{DateFormat: "2006-01-02", DataSource: config.DataSource{BaseCSV: "../../config/data/base_rate.csv", AgeCSV: "../../config/data/driver_age.csv", InsuranceGroupCSV: "../../config/data/insurance_group.csv", LicenseLengthCSV: "../../config/data/license.csv"}}
	var err error
	c.Base, c.MaxBase, c.MinBase, err = c.ReadFromCSV(c.DataSource.BaseCSV)
	assert.Nil(t, err)
	c.Age, c.MaxAge, c.MinAge, err = c.ReadFromCSV(c.DataSource.AgeCSV)
	assert.Nil(t, err)
	c.LicenseLength, c.MaxLicenseLength, c.MinLicenseLength, err = c.ReadFromCSV(c.DataSource.LicenseLengthCSV)
	assert.Nil(t, err)
	c.InsuranceGroup, c.MaxInsuranceGroup, c.MinInsuranceGroup, err = c.ReadFromCSV(c.DataSource.InsuranceGroupCSV)
	assert.Nil(t, err)
	confMock.On("GetConfig").Return(c)
	app := App{Config: confMock}
	resp, err := app.GeneratePricing(context.Background(), &pricingengine.GeneratePricingRequest{DateOfBirth: "2002-12-28", InsuranceGroup: 35, LicenseHeldSince: "2016-05-20"})
	assert.Nil(t, err)
	assert.Equal(t, 1.54, resp.DriverAgeFactor)
}
