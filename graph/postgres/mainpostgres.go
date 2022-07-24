package postgres

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func InitDbPool() {
	var host, port, user, password, dbname string
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "vishal123"
	dbname = "user"
	//	sslmode = os.Getenv(constants.POSTGRES_SSLMODE)

	databaseURL := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname

	maxOpenConnection, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_CONN"))
	if err != nil {
		log.Println(err)
		maxOpenConnection = 5
	}
	maxIdleTime, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_TIME"))
	if err != nil {
		log.Println(err)
		maxIdleTime = 5
	}
	maxConnectionLifetime, err := strconv.Atoi(os.Getenv("POSTGRES_MAX_LIFETIME"))
	if err != nil {
		log.Println(err)
		maxConnectionLifetime = 2
	}

	healthcheckperiod, err := strconv.Atoi(os.Getenv("POSTGRES_HEALTHCHECK_PREIOD"))
	if err != nil {
		log.Println(err)
		healthcheckperiod = 2
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Print(err)
		log.Print("CONFIG_ERR")
	}
	config.MaxConns = int32(maxOpenConnection)
	config.MaxConnLifetime = time.Duration(maxConnectionLifetime) * time.Minute
	config.HealthCheckPeriod = time.Duration(healthcheckperiod) * time.Minute
	config.MaxConnIdleTime = time.Duration(maxIdleTime) * time.Minute

	pool, err = pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Print(err)
		log.Print("POSTGRES_NOT_CONNECTED")
	} else {
		log.Println("POSTGRES_CONNECTED")
	}

}

func GetPool() *pgxpool.Pool {
	if pool == nil {
		InitDbPool()
	}

	return pool
}
