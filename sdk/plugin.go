package main

import (
	"github.com/miton18/maison/core"
	"github.com/spf13/viper"
)

// Name of the plugin
const Name = "My test plugin"

// Init setup plugin
func Init(config *viper.Viper) error {
	return nil
}

// Run plugin job
func Run(maison *core.Maison) {

}

// Close gracefully
func Close() {

}
