package postify

import "github.com/sirupsen/logrus"

const (
	Name    = "postify"
	Usage   = "generate and publish data"
	EnvName = "POSTIFY"
)

// define internals.
var (
	Logger *logrus.Logger = logrus.New()
)
