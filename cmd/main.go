package main

import (
	"github.com/joho/godotenv"
	"github.com/p12s/using-orm-with-db/internal/config"
	"github.com/p12s/using-orm-with-db/internal/repository"
	"github.com/p12s/using-orm-with-db/internal/server"
	"github.com/p12s/using-orm-with-db/internal/service"
	"github.com/p12s/using-orm-with-db/internal/transport/rest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	OPTIONAL_LOCAL_ENV_FILE = ".env" // in "GitHub Actions" should set variables from: internal/config/fixtures/.env.ok.example
	LOG_LEVEL               = zap.InfoLevel
)

func main() {
	logger := getLogger(LOG_LEVEL)
	defer logger.Sync() // nolint:errcheck
	log := logger.Sugar()

	loadLocalEnvIfExists(log)

	// check and define environment variables
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("environment assignment fail: %s\n", err.Error())
	}

	// repository layer (work with database)
	db, err := repository.NewPostgresDB(cfg.Db)
	if err != nil {
		log.Fatalf("db initialize fail: %s\n", err.Error())
	}
	repos := repository.NewRepository(db)
	// service layer (business/domain logic)
	services := service.NewService(repos)
	// handler (in the current server it works by the http protocol)
	handlers := rest.NewHandler(services)

	// http-server
	srv := server.NewHttpServer(cfg.Backend, *handlers)
	srv.Run(log)
	err = srv.WaitShutdown()
	if err != nil {
		log.Errorf("server shutdown fail: %s\n", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occurred on db connection close: %s", err.Error())
	}
	log.Info("db connect closed")
}

// loading .env (just for local development)
func loadLocalEnvIfExists(log *zap.SugaredLogger) {
	info, err := os.Stat(OPTIONAL_LOCAL_ENV_FILE)
	if os.IsNotExist(err) {
		return
	}
	if !info.IsDir() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("error reading env variables from file: %s\n", err.Error())
		}
	}
}

// getLogger - getting logger
func getLogger(level zapcore.Level) *zap.Logger {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(level)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	return zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
}
