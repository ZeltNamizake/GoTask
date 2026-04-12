package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var loc *time.Location

type TimeConfig struct {
	Timezone string `json:"timezone"`
}

const config = "config"

func LoadTimeConfig() TimeConfig {
	tzConfig := filepath.Join(config, "timezone.json")

	data, err := os.ReadFile(tzConfig)
	if err != nil {
		return TimeConfig{Timezone: ""}
	}

	var cfg TimeConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return TimeConfig{Timezone: ""}
	}

	return cfg
}

func EnsureTimezone() {
	cfg := LoadTimeConfig()

	if cfg.Timezone == "" {
		fmt.Println("[FIRST RUN SETUP]")
		PrintTimezoneList()

		for {
			tz := Ask("Choose timezone: ")

			if tz == "" {
				fmt.Println("[ERR] - Timezone cannot be empty")
				continue
			}

			if _, err := time.LoadLocation(tz); err != nil {
				fmt.Println("[ERR] - Invalid timezone")
				continue
			}

			cfg.Timezone = tz

			os.MkdirAll("config", 0755)

			data, _ := json.MarshalIndent(cfg, "", "  ")
			os.WriteFile(filepath.Join(config, "timezone.json"), data, 0644)

			fmt.Println("[INFO] timezone saved:", tz)
			break

		}
	}

	initTime(cfg.Timezone)
}

func initTime(tz string) {
	l, err := time.LoadLocation(tz)
	if err != nil {
		l, _ = time.LoadLocation("UTC")
	}
	loc = l
}

func Now() time.Time {
	if loc == nil {
		EnsureTimezone()
	}
	return time.Now().In(loc)
}

func DateNow() string {
	return Now().Format("2006-01-02")
}

func TimeNow() string {
	return Now().Format("15:04")
}

func Timezone() string {
	if loc == nil {
		EnsureTimezone()
	}
	return loc.String()
}

func PrintTimezoneList() {
	fmt.Println(`
Available Timezones:
- UTC
- Asia/Jakarta   (WIB)
- Asia/Makassar  (WITA)
- Asia/Jayapura  (WIT)
- Asia/Tokyo
- Asia/Singapore
- Europe/London
- America/New_York
`)
}
