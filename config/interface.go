package config

type IConfig interface {
	GetConfig() (Config)
	SetConfig() (error)
	ReadFromCSV(string) (map[int]float64, int, int, error)
}