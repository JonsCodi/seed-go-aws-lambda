package packages

import "text/template"

var ConfigPackageTemplate = template.Must(template.New("config.go").
	Parse(`package config

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strings"

	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"{{ .Name }}/internal/environment"
)

const (
	fatalLevel = "fatal"
	errorLevel = "error"
	warnLevel  = "warn"
	infoLevel  = "info"
	debugLevel = "debug"
)

func LoadZerologConfig() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(os.Stderr)
	level := strings.ToLower(os.Getenv("LOG_LEVEL"))

	switch level {
	case fatalLevel:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case errorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case warnLevel:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case debugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case infoLevel:
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
`))
