package config_file_reader

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ConfigS struct {
	TS string  `name:"testString"`
	TB bool    `name:"testBool"`
	TF float64 `name:"testFloat"`
	TI int     `name:"testInt"`
}

func LoadConfigAdvanced(path string) (ConfigS, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return ConfigS{}, err
	}

	// parse file → mapping
	mapping := map[string]string{}
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		key := parts[0]
		val := strings.Split(parts[1], ";")[0]

		mapping[key] = val
	}
	fmt.Println("mapping: ", mapping) // map[testBool:true testFloat:3.5 testInt:3 testString:hoang toni]

	cfg := ConfigS{}
	v := reflect.ValueOf(&cfg).Elem() // { false 0 0} original values
	t := v.Type()                     // config_file_reader.ConfigS

	// map struct fields
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("name")
		val := mapping[key]
		fmt.Println("key: ", key) // testString (ket get from Struct)
		fmt.Println("val: ", val) // hoang toni (value get from file)

		f := v.Field(i)
		fmt.Println("f: ", f) // init value of Struct

		switch f.Kind() {
		case reflect.String:
			f.SetString(val)
		case reflect.Bool:
			b, _ := strconv.ParseBool(val)
			f.SetBool(b)
		case reflect.Float64:
			n, _ := strconv.ParseFloat(val, 64)
			f.SetFloat(n)
		case reflect.Int:
			n, _ := strconv.Atoi(val)
			f.SetInt(int64(n))
		}
	}

	return cfg, nil
}

func PerformAdvanced() {
	cfg, err := LoadConfigAdvanced("config.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println("cfg", cfg)
}

// config.conf --> read file --> parse line(key/value/type)
// --> reflection scan struct tags --> match tag name --> Set value vào field
