package timeSeries

import (
	"time"

	"github.com/spf13/viper"
)

// Store is used to store times data like, room temperatures
type Store interface {

	// Init setup the store
	Init(c viper.Viper) error

	// Close must gracefully stop the store
	Stop()

	// Put a metric
	Put(metric string, context map[string]string, t time.Time, value interface{}) error

	// Fetch a metric
	Fetch(metric string, context map[string]string, options interface{}) (interface{}, error)
}
