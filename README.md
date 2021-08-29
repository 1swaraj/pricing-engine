# Pricing Engine

#### How to Execute

```
git clone https://github.com/1swaraj/pricing-engine.git
cd pricing-engine
```

If you have make installed on your system then :-
```
make build
make test
make run
```

Else the classic way
```go
cd cmd
go build .
./cmd
```

#### Understanding the CodeBase

##### pricing-engine-go
1. constants.go :-
   1. Stores all the error messages
   2. Stores the default location for the config.yaml file
   
2. pricing_engine.go :-
   1. Stores the structs needed for the request and response

##### config
This directory has code that is primarily concerned with the config management like reading the configs from a yaml file and reading the data from the csv files present in the data (config/data) directory
1. data :-
   1. This directory has config.yaml which stores the configuration details such as the date format and the data sources corresponding to base_rate, driver_age factor etc
2. interface.go :-
   1. Defines the blueprint for the config package
   2. Intentionally create to make creation of test cases possible
3. types.go :- 
   1. Defines the Config struct which stores information about things like :-
      1. Date Format
      2. Data Source - Paths of all the CSV files needed for calulations
      3. Base,Age,LicenseLength,InsuranceGroup - used to store the hash map generated from the csv
      4. MaxBase, MinBase, MaxAge, MinAge etc - used to store the max / min value in the range (generated from the CSV)
4. utils.go :- 
   1. Helps us with utilities like ReadFromCSV and ReadFromYAML
5. config.go :- 
   1. Implements the GetConfig and SetConfig methods defined in the interface

##### service

###### app
1. find_factor.go :-
   1. Implements a generic method that finds the factor corresponding to the input
   2. i.e. GetFactor method will be able to find the age factor as well as the insurance group factor
2. generate_pricing.go :- 
   1. Implements the core logic of the application
3. generate_pricing_test.go :-
   1. Wrote a few test cases (that depict the importance of the architectural choice of defining the IConfig interface)
4. types.go :-
   1. Defines the App struct
   
###### rpc
Implements the rpc call mechanism (go-chi)

##### utils
1. interface.go :-
   1. Defines the IDateHelper interface
   2. This interface will be helpful in generating mocks for testing
   3. Similar to what we did for IConfig interface
2. utils.go :- 
   1. Implements functions that will help us in doing date time related operations.
   2. Eg :- AgeFromDate,Parse etc 
      This also helps in making the code extensible. If tomorrow, time.Time is deprecated for some reason,
      we just need to make the changes in this package and not in the entire codebase.
