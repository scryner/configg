package configg

import (
	"os"
	"reflect"
	"fmt"
	"bytes"
	"encoding/json"
	"strconv"
)

type Config struct {
	unparsed string
	parsed map[string]interface{}
}

func newConfig() *Config {
	c := new(Config)
	c.parsed = make(map[string]interface{})

	return c
}

func LoadConfigFile(filename string) (config *Config, err error) {
	config = newConfig()

	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	if err != nil {
		return
	}

	config.unparsed = b.String()
	err = json.Unmarshal(b.Bytes(), &config.parsed)
	if err != nil {
		return
	}

	return
}

func LoadConfigString(s string) (config *Config, err error) {
	config = newConfig()

	config.unparsed = s
	err = json.Unmarshal([]byte(s), &config.parsed)
	if err != nil {
		return
	}

	return
}

func (c *Config) getValue(k string, v interface{}) error {
	ret, ok := c.parsed[k]
	if !ok {
		return fmt.Errorf("not found: %s", k)
	}

	val := reflect.ValueOf(v)
	val_type := val.Type()

	if val_type.Kind() != reflect.Ptr {
		return fmt.Errorf("destination is must be a pre-allocated pointer")
	}

	val2 := reflect.Indirect(val)
	if !val2.CanInterface() || !val2.CanSet() {
		return fmt.Errorf("destination is not available to set")
	}

	v2 := val2.Interface()

	switch v2.(type) {
	case int:
		switch ret2 := ret.(type) {
		case int:
			val2.Set(reflect.ValueOf(ret2))

		case float64:
			ret3 := int(ret2)
			val2.Set(reflect.ValueOf(ret3))

		case string:
			if ret3, err := strconv.Atoi(ret2); err != nil {
				return fmt.Errorf("type not matched: %s would be int", k)
			} else {
				val2.Set(reflect.ValueOf(ret3))
			}

		default:
			return fmt.Errorf("type not matched: %s would be int", k)
		}

	case float64:
		switch ret2 := ret.(type) {
		case float64:
			val2.Set(reflect.ValueOf(ret2))

		case string:
			if ret3, err := strconv.ParseFloat(ret2, 64); err != nil {
				return fmt.Errorf("type not matched: %s would be float", k)
			} else {
				val2.Set(reflect.ValueOf(ret3))
			}

		default:
			return fmt.Errorf("type not matched: %s would be float", k)
		}

	case string:
		if _, ok := ret.(string); ok {
			val2.Set(reflect.ValueOf(ret))
		} else {
			return fmt.Errorf("type not matched: %s would be string", k)
		}

	case []interface{}:
		if _, ok := ret.([]interface{}); ok {
			val2.Set(reflect.ValueOf(ret))
		} else {
			return fmt.Errorf("type not matched: %s would be array", k)
		}
	}

	return nil
}

func (c *Config) GetString(key string) (val string, err error) {
	err = c.getValue(key, &val)
	return
}

func (c *Config) GetInt(key string) (val int, err error) {
	err = c.getValue(key, &val)
	return
}

func (c *Config) GetFloat(key string) (val float64, err error) {
	err = c.getValue(key, &val)
	return
}

func (c *Config) GetArray(key string) (val []interface{}, err error) {
	err = c.getValue(key, &val)
	return
}