package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

var singleInstance *sql.DB

func GetInstance() *sql.DB {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			db, err := sql.Open("sqlite3", "./wc-predictor.db")
			if err != nil {
				log.Fatal(err)
			}
			singleInstance = db
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
