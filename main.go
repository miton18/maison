package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/miton18/maison/bus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// init config
	setDefaultConfig()
	viper.SetEnvPrefix("maison")

	viper.AddConfigPath("/etc/maison/")
	viper.AddConfigPath("$HOME/.maison")
	viper.AddConfigPath(".")

	viper.BindEnv()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		stop()
		run()
	})

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Logger
	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	run()

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGINT)
	<-sigs

	stop()
}

// stop must safely close all modules
func stop() {
	log.Info("Maison is stopping...")

	closePlugins()

	// TODO: close stores

	bus.Stop()

	log.Info("Maison is stopped")
}

// run must start dependancies, modules, etc..
func run() {
	log.Info("Maison is starting...")

	bus.Init()

	// TODO: init stores

	initPlugins()

	log.Info("Maison is started")
}
