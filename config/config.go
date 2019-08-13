package config

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
	"os"
	"time"
)

type ConfigList struct {
	ApiKey      string
	ApiSecret   string
	LogFile     string
	Productcode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	Dbname        string
	SQLDriver     string
	Port          int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Fail to read file : %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		"1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
	}

	Config = ConfigList{
		ApiKey:        cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret:     cfg.Section("bitflyer").Key("secret_key").String(),
		LogFile:       cfg.Section("gotrading").Key("log_file").String(),
		Productcode:   cfg.Section("gotrading").Key("product_code").String(),
		Durations:     durations,
		TradeDuration: durations[cfg.Section("gotrading").Key("trading_duration").String()],
		Dbname:        cfg.Section("db").Key("name").String(),
		SQLDriver:     cfg.Section("db").Key("driver").String(),
		Port:          cfg.Section("web").Key("port").MustInt(),
	}

}
