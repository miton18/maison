package core

import (
	"github.com/sirupsen/logrus"
)

// Maison is used by plugins, it's the merge of all functionalities
type Maison struct {
	Logger logrus.Entry
}
