# Pricing Engine

#### How to Execute

```
git clone https://github.com/1swaraj/pricing-engine.git
```

If you have make installed on your system then :-
```makefile
make run
```

Else the classic way
```go
cd cmd
go build .
./cmd
```

#### Understanding the CodeBase

###### config
This directory has code that is primarily concerned with the config management like reading the configs from a yaml file and reading the data from the csv files present in the data (config/data) directory
1. data :-
   1. This directory has config.yaml which stores the configuration details such as the date format and the data sources corresponding to base_rate, driver_age factor etc
2. interface.go :-
   1. Has the