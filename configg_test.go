package configg

import (
	"fmt"
	"testing"
)

func testInt(c *Config, key string, expected int) error {
	intVal, err := c.GetInt(key)
	if err != nil {
		return fmt.Errorf("getting int failed: %v", err)
	}

	if intVal != expected {
		return fmt.Errorf("int value is not matched to %d", expected)
	}

	return nil
}

func testFloat(c *Config, key string, expected float64) error {
	floatVar, err := c.GetFloat64(key)
	if err != nil {
		return fmt.Errorf("getting float failed: %v", err)
	}

	if floatVar != expected {
		return fmt.Errorf("float value is not matched to %f", expected)
	}

	return nil
}

func testString(c *Config, key string, expected string) error {
	stringVar, err := c.GetString(key)
	if err != nil {
		return fmt.Errorf("getting string failed: %v", err)
	}

	if stringVar != expected {
		return fmt.Errorf("string value is not matched to \"%s\"", expected)
	}

	return nil
}

func testBool(c *Config, key string, expected bool) error {
	boolVar, err := c.GetBool(key)
	if err != nil {
		return fmt.Errorf("getting bool failed: %v", err)
	}

	if boolVar != expected {
		return fmt.Errorf("bool value is not matched to \"%v\"", expected)
	}

	return nil
}

func testArray(c *Config, key string, expected []interface{}) error {
	arrayVar, err := c.GetArray(key)
	if err != nil {
		return fmt.Errorf("getting array failed: %v", err)
	}

	matched := 0

	for i, val := range arrayVar {
		if val == expected[i] {
			matched += 1
		}
	}

	if matched != len(arrayVar) {
		return fmt.Errorf("array value is not matched to %v", expected)
	}

	return nil
}

func TestConfig(t *testing.T) {
	c, err := LoadConfigString(`{"intParam":1,"floatParam":1.2,"stringParam":"hello","boolParam":"true","arrayParam":["a", "b", "c"]}`)
	if err != nil {
		t.Errorf("parse error")
		return
	}

	// test int
	if err := testInt(c, "intParam", 1); err != nil {
		t.Errorf("%v", err)
		return
	}

	// test float64
	if err := testFloat(c, "floatParam", 1.2); err != nil {
		t.Errorf("%v", err)
		return
	}

	// test bool
	if err := testBool(c, "boolParam", true); err != nil {
		t.Errorf("%v", err)
		return
	}

	// test string
	if err := testString(c, "stringParam", "hello"); err != nil {
		t.Errorf("%v", err)
		return
	}

	// test array
	if err := testArray(c, "arrayParam", []interface{}{"a", "b", "c"}); err != nil {
		t.Errorf("%v", err)
		return
	}
}

func TestConfigFlexibility(t *testing.T) {
	c, err := LoadConfigString(`{"i1":1,"i2":2.0,"i3":"3","f1":1,"f2":2.1,"f3":"3.1"}`)
	if err != nil {
		t.Errorf("parse error")
		return
	}

	// test int
	if err := testInt(c, "i1", 1); err != nil {
		t.Errorf("%v", err)
		return
	}

	if err := testInt(c, "i2", 2); err != nil {
		t.Errorf("%v", err)
		return
	}

	if err := testInt(c, "i3", 3); err != nil {
		t.Errorf("%v", err)
		return
	}

	// test float
	if err := testFloat(c, "f1", 1.0); err != nil {
		t.Errorf("%v", err)
		return
	}

	if err := testFloat(c, "f2", 2.1); err != nil {
		t.Errorf("%v", err)
		return
	}

	if err := testFloat(c, "f3", 3.1); err != nil {
		t.Errorf("%v", err)
		return
	}
}
