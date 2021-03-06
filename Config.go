package main

import (
	"os"
	"path"

	"github.com/naoina/toml"
)

type Config struct {
	Scheme string `toml:"scheme"`
	Endpoint string `toml:"endpoint"` }

var DefaultConfig = Config {
	Scheme: "https",
	Endpoint: "ralee.org/api" }

func FindConfig() (string) {
	home, _ := os.UserHomeDir()
	locations := [3]string{
		"config.toml",
		path.Join(home, ".ral", "config.toml"),
		path.Join(home, ".ralrc") }

	for _, l := range locations {
		if FileExists(l) { return l }
	}
	return ""
}

func FileExists(fpath string) (bool) {
	info, err := os.Stat(fpath)
	if os.IsNotExist(err) {
		return false }
	return !info.IsDir()
}

func ReadSystemConfig() (config *Config) {
	if tryFile := FindConfig(); tryFile != "" {
		config, _ := ReadConfig(tryFile)
		return config
	}

	return &DefaultConfig
}

func ReadConfig(fpath string) (config *Config, err error) {
	config = new(Config)
	f, err := os.Open(fpath)
	if err != nil { return }
	defer f.Close()

	err = toml.NewDecoder(f).Decode(config)
	if err != nil { return }

	return
}
