package gotron

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logger = log.Logger
)

//UseLogger in this library and add a defined logger
func UseLogger(zl zerolog.Logger) {
	logger = zl
}
