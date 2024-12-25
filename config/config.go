package config

import (
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v2"
)

type ConfigS struct {
	Address                    string
	UpdateUIEveryXMilliseconds int
	UIPaddingIndentAmount      int
}

var DefaultConfig = ConfigS{
	Address:                    "192.168.1.41:9999",
	UpdateUIEveryXMilliseconds: 1000,
	UIPaddingIndentAmount:      2,
}

var Config ConfigS

func Init() {
	configPath := configdir.LocalConfig("ontake", "creality-print-cli")
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}

	configFile := filepath.Join(configPath, "config.yml")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		fh, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		encoder := yaml.NewEncoder(fh)
		encoder.Encode(&DefaultConfig)
		Config = DefaultConfig
	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		decoder := yaml.NewDecoder(fh)
		decoder.Decode(&Config)
	}
}
