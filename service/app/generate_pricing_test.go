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
	//
	resp, err := app.GeneratePricing(context.Background(), &pricingengine.GeneratePricingRequest{DateOfBirth: "2002-12-28", InsuranceGroup: 35, LicenseHeldSince: "2016-05-20"})
	assert.Nil(t, err)
	assert.Equal(t, 1.54, resp.DriverAgeFactor)
	assert.Equal(t, 1.12, resp.InsuranceGroup)
	assert.Equal(t, 1.025, resp.LicenseLength)
	expectedBaseCosts := []pricingengine.BaseCost{{
			TimeSec:   1800,
			TimeHours: 0,
			TimeDays:  0,
			BaseRate:  482.64216000000005,
		},
		{
			TimeSec:   3600,
			TimeHours: 1,
			TimeDays:  0,
			BaseRate:  871.58456,
		},
		{
			TimeSec:   7200,
			TimeHours: 2,
			TimeDays:  0,
			BaseRate:  1334.7796,
		},
		{
			TimeSec:   10800,
			TimeHours: 3,
			TimeDays:  0,
			BaseRate:  1764.38416,
		},
		{
			TimeSec:   21600,
			TimeHours: 6,
			TimeDays:  0,
			BaseRate:  2195.75664,
		},
		{
			TimeSec:   43200,
			TimeHours: 12,
			TimeDays:  0,
			BaseRate:  3594.18136,
		},
		{
			TimeSec:   86400,
			TimeHours: 24,
			TimeDays:  1,
			BaseRate:  3908.87112,
		},
		{
			TimeSec:   172800,
			TimeHours: 48,
			TimeDays:  2,
			BaseRate:  5743.972080000001,
		},
		{
			TimeSec:   259200,
			TimeHours: 72,
			TimeDays:  3,
			BaseRate:  7812.438480000001,
		},
		{
			TimeSec:   345600,
			TimeHours: 96,
			TimeDays:  4,
			BaseRate:  9200.25568,
		}}
	assert.Equal(t, expectedBaseCosts, resp.BaseCosts)

	resp, err = app.GeneratePricing(context.Background(), &pricingengine.GeneratePricingRequest{DateOfBirth: "2032-12-28", InsuranceGroup: 35, LicenseHeldSince: "2016-05-20"})
	assert.Equal(t, pricingengine.LicenseBeforeBirth,err.Error())

	resp, err = app.GeneratePricing(context.Background(), &pricingengine.GeneratePricingRequest{DateOfBirth: "2032-12-28", InsuranceGroup: 35, LicenseHeldSince: "2046-05-20"})
	assert.Equal(t, pricingengine.DateAheadToday,err.Error())
}
