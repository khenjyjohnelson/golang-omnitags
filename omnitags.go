package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config stores all mappings dynamically
type Config struct {
	Aliases     map[string]string
	Reverse     map[string]string
	VInput      map[string]string
	VPost       map[string]string
	VGet        map[string]string
	Flash1Msg   map[string]string
	Flash       map[string]string
	FlashFunc   map[string]string
	FlashMsg    map[string]string
	VUploadPath map[string]string
	Views       map[string]string
	Titles      map[string]string
	V           map[int]string
	TL          map[string]interface{}
}

// NewConfig initializes a new instance of Config with empty maps
func NewConfig() *Config {
	c := &Config{
		Aliases:     make(map[string]string),
		Reverse:     make(map[string]string),
		VInput:      make(map[string]string),
		VPost:       make(map[string]string),
		VGet:        make(map[string]string),
		Flash1Msg:   make(map[string]string),
		Flash:       make(map[string]string),
		FlashFunc:   make(map[string]string),
		FlashMsg:    make(map[string]string),
		VUploadPath: make(map[string]string),
		Views:       make(map[string]string),
		Titles:      make(map[string]string),
		V:           make(map[int]string),
		TL:          make(map[string]interface{}),
	}

	// Initialize `V` dynamically
	for i := 1; i <= 11; i++ {
		c.V[i] = fmt.Sprintf("contents/section_%d", i)
		if i <= 6 {
			c.Flash1Msg[fmt.Sprintf("flash_%d", i)] = fmt.Sprintf("Flash message %d", i)
		}
		if i <= 5 {
			c.FlashMsg[fmt.Sprintf("error_%d", i)] = fmt.Sprintf("Error message %d", i)
		}
	}

	// Initialize `TL` dynamically
	c.TL["ot"] = nil
	c.TL["a1"] = nil
	groups := map[string]int{"b": 11, "c": 2, "d": 4, "e": 8, "f": 4}
	for group, count := range groups {
		for i := 1; i <= count; i++ {
			c.TL[fmt.Sprintf("%s%d", group, i)] = nil
		}
	}

	return c
}

// LoadData extracts key-value pairs from JSON and initializes mappings
func (c *Config) LoadData(data map[string]interface{}) {
	if values, ok := data["values"].([]interface{}); ok {
		for _, item := range values {
			if obj, isObject := item.(map[string]interface{}); isObject {
				key, keyExists := obj["key"].(string)
				value, valueExists := obj["value"].(string)

				if keyExists && valueExists {
					// Aliases & Reverse Mapping
					c.Aliases[key] = value
					c.Reverse[value+"_realname"] = key

					// Input Fields
					c.VInput[key+"_input"] = "txt_" + value
					c.VInput[key+"_filter1"] = "min_" + value
					c.VInput[key+"_filter2"] = "max_" + value
					c.VInput[key+"_old"] = "old_" + value
					c.VInput[key+"_new"] = "new_" + value
					c.VInput[key+"_confirm"] = "confirm_" + value

					// Post & Get Requests
					c.VPost[key] = "txt_" + value
					c.VPost[key+"_old"] = "old_" + value
					c.VPost[key+"_new"] = "new_" + value
					c.VPost[key+"_confirm"] = "confirm_" + value

					c.VGet[key] = "txt_" + value
					c.VGet[key+"_filter1"] = "min_" + value
					c.VGet[key+"_filter2"] = "max_" + value

					// Flash Messages
					c.Flash1Msg[key] = value + " successfully saved!"
					c.Flash[key] = "pesan_" + value
					c.FlashFunc[key] = "$(\"." + value + "\").modal(\"show\")"
					c.FlashMsg[key] = value + " tidak bisa diupload!"

					// Upload Path
					c.VUploadPath[key] = "./assets/img/" + key + "/"

					// Views
					c.Views[key] = "contents/" + key + "/index"
					c.Views[key+"_daftar"] = "contents/" + key + "/daftar"
					c.Views[key+"_admin"] = "contents/" + key + "/admin"
					c.Views[key+"_laporan"] = "contents/" + key + "/laporan"
					c.Views[key+"_print"] = "contents/" + key + "/print"

					// Titles
					c.Titles[key+"_v1"] = value
					c.Titles[key+"_v2"] = "List of " + value
					c.Titles[key+"_v3"] = value + " Data"
					c.Titles[key+"_v4"] = value + " Report"
					c.Titles[key+"_v5"] = value + " Data"
					c.Titles[key+"_v6"] = value + " Profile"
					c.Titles[key+"_v7"] = value + " Successful!"
				}
			}
		}
	}
}

// GetValue fetches a value dynamically
func (c *Config) GetValue(field string) string {
	if value, exists := c.Aliases[field]; exists {
		return value
	}
	return "Unknown Field"
}