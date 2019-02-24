package document

import (
	"io"

	"github.com/spf13/viper"
)

// Store is used to store and query documents like text files, images
type Store interface {

	// Init setup the store
	Init(c viper.Viper) error

	// Close must gracefully stop the store
	Stop()

	// Open an existring document
	Open(documentID string) (io.ReadCloser, error)

	// Save or update a new document
	Save(documentID string, document io.WriteCloser) error

	// List existing keys
	List() []string
}
