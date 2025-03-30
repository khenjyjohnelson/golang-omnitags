package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config stores key-value mappings
type Config struct {
	Aliases map[string]string
}

// LoadData extracts key-value pairs from "values" array
func (c *Config) LoadData(data map[string]interface{}) {
	c.Aliases = make(map[string]string)

	if values, ok := data["values"].([]interface{}); ok {
		for _, item := range values {
			if obj, isObject := item.(map[string]interface{}); isObject {
				key, keyExists := obj["key"].(string)
				value, valueExists := obj["value"].(string)

				if keyExists && valueExists {
					c.Aliases[key] = value
				}
			}
		}
	}
}

// GetValue fetches a value from the loaded data
func (c *Config) GetValue(field string) string {
	if value, exists := c.Aliases[field]; exists {
		return value
	}
	return "Unknown Field"
}

func main() {
	// Read JSON file
	file, err := os.ReadFile("app.postman_environment.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse JSON into a map
	var jsonData map[string]interface{}
	if err := json.Unmarshal(file, &jsonData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Load JSON data into Config struct
	config := &Config{}
	config.LoadData(jsonData)

	// Example: Fetching values dynamically
	fmt.Println("Database Name:", config.GetValue("database"))           // me_work
	fmt.Println("Table A1 Name:", config.GetValue("tabel_a1"))           // ot_website
	fmt.Println("Table A1 Alias:", config.GetValue("tabel_a1_alias"))     // Website Settings
	fmt.Println("Table A1 Field1:", config.GetValue("tabel_a1_field1"))   // id
	fmt.Println("Table A1 Field1 Alias:", config.GetValue("tabel_a1_field1_alias")) // Website ID
}
