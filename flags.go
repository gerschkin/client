package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var configPath = kingpin.Flag("config", "Path to config").Default("/etc/gerschkin/client.toml").String()
