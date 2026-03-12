package config_file_reader

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	Port   int    `json:"port"`
	DBHost string `json:"db_host"`
}

// pointer --> &0xc0000a6018 || nil
func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("file:", file)                                 // [123 10 32 ... 125 10]
	fmt.Println("reflect.TypeOf(file):", reflect.TypeOf(file)) // []uint8

	var cfg Config
	// json.Unmarshal parse data([]byte chứa JSON) --> mapping & fill vào Struct &cfg (pointer)
	// Unmarshal need modidy Struct --> must use pointer
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func Perform() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("cfg.Port:", cfg.Port)
	fmt.Println("cfg.DBHost:", cfg.DBHost)
}
