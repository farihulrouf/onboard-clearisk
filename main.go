package main

import (
	"fmt"
	"BE_CLEARISK.IO/config"
	"BE_CLEARISK.IO/database"
	"net/http"
	"github.com/gorilla/mux"
	
)

func main() {
	conf := config.GetConfig()
	//db := database.ConnectDB(ctx, conf.Mongo)
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(db)
	r := mux.NewRouter()
	http.ListenAndServe(":8222",r)

	/*
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	*/

	
}
