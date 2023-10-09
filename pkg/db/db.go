package db

import (
	"database/sql"
	"iudx_domain_specific_apis/pkg/configs"
	"iudx_domain_specific_apis/pkg/logger"
	"log"
	"os"

	"github.com/go-gorp/gorp/v3"
	_ "github.com/lib/pq" //import postgres
)

var db *gorp.DbMap

func Init() {

	var err error
	db, err = ConnectDB(configs.GetDBConnStr())
	if err != nil {
		log.Fatal(err)
	}

}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	logger.Info.Println("Successfully connected to database")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests

	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}

func Close() error {
	logger.Info.Println("Closing DB connection")
	return db.Db.Close()
}
