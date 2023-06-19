package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/boj/redistore"
	"github.com/ftsog/auth/database"
	"github.com/ftsog/auth/handlers"
	"github.com/ftsog/auth/models"
	"github.com/ftsog/auth/routers"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

type pgEnVariable struct {
	pgHost      string
	pgPort      string
	pgUser      string
	pgDBName    string
	pgPassword  string
	rdHost      string
	rdPort      string
	rdPassword  string
	rdSecretKey string
}

func LoadEnVariable() *pgEnVariable {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return &pgEnVariable{
		pgHost:      os.Getenv("PGHOST"),
		pgPort:      os.Getenv("PGPORT"),
		pgUser:      os.Getenv("PGUSER"),
		pgDBName:    os.Getenv("PGDBNAME"),
		pgPassword:  os.Getenv("PGPASSWORD"),
		rdHost:      os.Getenv("RDHOST"),
		rdPort:      os.Getenv("RDPORT"),
		rdPassword:  os.Getenv("RDPASSSWORD"),
		rdSecretKey: os.Getenv("RDSECRETKEY"),
	}
}

func Connections() (*sql.DB, *redistore.RediStore) {
	enVariable := LoadEnVariable()

	dbConn := database.PgDBConnection(enVariable.pgHost, enVariable.pgPort, enVariable.pgUser, enVariable.pgPassword, enVariable.pgDBName, "")
	rdConn := database.RediStore(10, "tcp", enVariable.rdHost, enVariable.rdPort, enVariable.rdPassword, []byte(enVariable.rdSecretKey))

	return dbConn, rdConn
}

func NewModel() *models.Model {
	dbConn, rdConn := Connections()

	model := &models.Model{
		PgDBConn:  dbConn,
		RediStore: rdConn,
	}

	return model
}

func NewLogger() *handlers.Logger {
	logFile, err := os.OpenFile("auth.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Panic(err)
	}
	logger := &handlers.Logger{
		Info:    log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Llongfile),
		Warning: log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile),
		Debug:   log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile),
		Error:   log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile),
		Fatal:   log.New(logFile, "FATAL", log.Ldate|log.Ltime|log.Llongfile),
	}

	return logger
}

func NewHandler() *handlers.Handler {
	logger := NewLogger()
	model := NewModel()

	handler := &handlers.Handler{
		Model: model,
		Log:   logger,
	}

	return handler
}

func NewRouter() *routers.Router {
	mux := chi.NewRouter()
	handler := NewHandler()

	router := &routers.Router{
		Mux:     mux,
		Handler: handler,
	}

	router.Route()

	return router
}
