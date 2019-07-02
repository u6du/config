package config

import (
	"os"

	"github.com/u6du/zerolog"
	"github.com/u6du/zerolog/log"
)

func init() {
	//	log.SetColored(true)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

}
