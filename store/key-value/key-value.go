package keyValue

import "github.com/spf13/viper"

// Store is used to store and query things like variables, strings, numbers
type Store interface {

	// Init setup the store
	Init(c viper.Viper) error

	// Close must gracefully stop the store
	Stop()

	// Set store a value under a key
	Set(key string, i interface{}) error

	// Get return the value for a given Key
	Get(key string) (interface{}, error)

	// List existing keys
	List() []string
}
