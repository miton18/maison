package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	version          = "x.x.x"
	gitHash          = "xxxxx"
	workingDirectory = "/"
)

func init() {
	var err error
	workingDirectory, err = os.Getwd()
	if err != nil {
		log.Warnf("Could not detect working directory: %+v", err)
	}
}

func setDefaultConfig() {

	// Core
	viper.SetDefault("core.version", version)
	viper.SetDefault("core.git.hash", gitHash)
	viper.SetDefault("core.working-directory", workingDirectory)

	// Plugins
	viper.SetDefault("plugins-directory", filepath.Join(workingDirectory, "plugins"))
	viper.SetDefault("plugins", map[string]interface{}{})

	// API
	viper.SetDefault("api.listen.address", "127.0.0.1")
	viper.SetDefault("api.listen.port", "8080")
}
