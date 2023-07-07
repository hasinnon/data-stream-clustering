package main

import (
	"fmt"
	"log"
	"my/ar/399/datastream/controller/restapi"
	"my/ar/399/datastream/datalayer"
)

func main() {
	db, err := datalayer.CreateDBConnection("root@/clustream_sru?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println("connectioon to database failed")
		log.Fatalln(err)
	}
	// defer UpdateAll(db)

	restapi.RunAPI("localhost:8383", *db)

}
