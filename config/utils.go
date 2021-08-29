package config

import (
	"encoding/csv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"pricingengine"
	"strconv"
	"strings"
)

// Reading Yaml and Unmarshalling it to config.Config struct
func (conf *Config) ReadFromYaml() error {
	yamlFile, err := ioutil.ReadFile(pricingengine.ConfigurationYaml)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return err
	}
	return nil
}

// Reading the CSV file and creating a map[int]float64
// Map will look something like
// ______________
// | Age | Factor |   ==> For representation purposes only
// ---------------
// |  0  |  1.1   |
// |  1  |  1.2   |
// |  2  |  1.3   |
// ---------------
// Now if we have a CSV that has 0-4 (range) as key we will iterate through
// that range and update the corresponding value
// Age Factor
// 0-2   1.1
// 3     1.2
// ______________
// | Age | Factor |
// ---------------
// |  0  |  1.1   |
// |  1  |  1.1   |
// |  2  |  1.1   |
// |  3  |  1.2   |
// ---------------
// This is done to make the system optimized for processing thousands of requests
// a minute
func (conf *Config) ReadFromCSV(fileName string) (kv map[int]float64, maximum, minimum int, err error) {
	kv = make(map[int]float64)
	lines, err := readCsvAsStringArray(fileName)
	if err != nil {
		return
	}
	// Loop through lines & turn into ranges
	for _, line := range lines[1:] {
		factor, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return kv, maximum, minimum, err
		}

		// If the csv contains a range
		// Eg :
		// Age    Factor
		// 0-2      1.1
		// We create a map for the entire range
		// kv[0] = 1.1
		// kv[1] = 1.1
		// kv[2] = 1.1
		if strings.Contains(line[0], "-") {
			key := strings.Split(line[0], "-")
			min, err := strconv.ParseInt(key[0], 10, 64)
			if err != nil {
				return kv, maximum, minimum, err
			}
			max, err := strconv.ParseInt(key[1], 10, 64)
			if err != nil {
				return kv, maximum, minimum, err
			}
			if min > max {
				min, max = max, min
			}
			for i := min; i <= max; i++ {
				kv[int(i)] = factor
			}
			if int(max) > maximum {
				maximum = int(max)
			}
			if int(min) < minimum {
				minimum = int(min)
			}
			continue
		}

		// If we have just one key then we simply mark that key
		// in the map
		key, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			return kv, maximum, minimum, err
		}
		kv[int(key)] = factor

		// Keeping track of the maximum and minimum values
		// in the csv. This will help us in optimizing our
		// system.
		if int(key) > maximum {
			maximum = int(key)
		}
		if int(key) < minimum {
			minimum = int(key)
		}
	}
	return
}

// readCsvAsStringArray (note r is small because we don't need to expose this function to other packages)
func readCsvAsStringArray(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}