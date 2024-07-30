package main

import (
	"fmt"
	"log"

	db2POOL "github.com/ibmdb/go_ibm_db"
)

const (
	HOSTNAME = "localhost"
	DATABASE = "SAMPLE"
	PORT     = 50000
	UID      = "db2inst1"
	PWD      = "password"
)

func main() {
	// Connect to the DB2 database
	con := fmt.Sprintf("HOSTNAME=%s;DATABASE=%s;PORT=%d;UID=%s;PWD=%s",
		HOSTNAME,
		DATABASE,
		PORT,
		UID,
		PWD)

	// SetConnMaxLifetime will take the value in SECONDS
	pool := db2POOL.Pconnect("PoolSize=100")
	db2 := pool.Open(con, "SetConnMaxLifetime=30")
	if db2 == nil {
		log.Fatalf("failed to open DB2 connection")
	}
	defer db2.Close()
	defer pool.Release()

	// Check if the connection is alive
	if db2.Ping() != nil {
		log.Fatalf("failed to ping DB2 connection")
	}

	fmt.Println("Successfully connected to DB2")
}
