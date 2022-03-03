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

func StartDatabaseConnection(env environment.Environment) (*gorm.DB, error) {
	log.Info().Msg("Loading Database connection....")

	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName("{{ .Product}}-{{ .Name}}-psql"))
	sqlDb, err := sqltrace.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			env.DBUser,
			env.DBPassword,
			env.DBHost,
			env.DBPort,
			env.DBName,
			env.DBSslMode,
		),
	)
	if err != nil {
		return nil, err
	}

	db, err := gormtrace.Open(postgres.New(postgres.Config{Conn: sqlDb}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Info().Msg("Loading Database connection....Done")

	return db, nil
}

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
