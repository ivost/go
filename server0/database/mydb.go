package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"server/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

var db *sql.DB
var insertStmt *sql.Stmt

func init() {
	log.SetOutput(os.Stdout)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	log.Printf("Open db %s", dbname)
	db, _ = sql.Open("postgres", psqlInfo)
	if err = db.Ping(); err != nil {
		log.Printf("database error: %v", err)
	}
	sql := `CREATE TABLE IF NOT EXISTS poi(
			id serial primary key,
			name varchar(50) not null,
			address1 varchar(50) null,
			address2 varchar(50) null,
			zip varchar(5) null,
			zipsuffix varchar(5) null,
			phone varchar(15) null,
			latitude float not null,
			longitude  float not null,
			radius float not null
			)`
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("error %v", err)
	}

	insertStmt, err = db.Prepare(`INSERT INTO poi
		(name, address1, address2, zip, zipsuffix, phone, latitude, longitude, radius)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		log.Printf("error %v", err)
	}
	//defer db.Close()
}

func AddPoi(poi * model.POI) (err error) {
	log.Printf("AddPoi %v", *poi)
	_, err = insertStmt.Exec(poi.Name, poi.Address1, poi.Address2, poi.Zip, poi.ZipSuffix, poi.Phone,
		poi.Lat, poi.Lng, poi.Radius)
	if err != nil {
		log.Printf("error %v", err)
	}
	return err
}

func Close() {
	db.Close()
}
