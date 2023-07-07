package datalayer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muesli/cache2go"
)

// MyDB  stores a pointer to database
type MyDB struct {
	db    *sql.DB
	cache *cache2go.CacheTable
}

// CreateDBConnection  Create Connection to database
func CreateDBConnection(connString string) (*MyDB, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	newcache := cache2go.Cache("serviceCache")
	mydb := MyDB{
		db:    db,
		cache: newcache,
	}
	newcache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		err = mydb.UpdateCServices(e.Data().(*CService))
		if err != nil {
			println(err)
		}
		//fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})
	return &mydb, nil
}
