package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/khenjyjohnelson/golang-omnitags/omnitags"
)

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

	// Initialize Config and load JSON data
	config := NewConfig()
	config.LoadData(jsonData)

	// Example Outputs
	fmt.Println("Database Name:", config.GetValue("database"))
	fmt.Println("Table A1 Name:", config.GetValue("tabel_a1"))
	fmt.Println("Upload Path for A1:", config.VUploadPath["tabel_a1"])
	fmt.Println("Flash Message for A1:", config.Flash1Msg["tabel_a1"])
}
