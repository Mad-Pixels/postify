package postify

import "github.com/sirupsen/logrus"

const (
	Name  = "content"
	Usage = "generate and publish data"
)

// define internals.
var (
	Logger *logrus.Logger = logrus.New()
)
