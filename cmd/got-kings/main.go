package main

import (
	"fmt"
	"net/http"

	gotKings "github.com/dennisssdev/game-of-thones-data/rpc/got-kings"

	server "github.com/dennisssdev/game-of-thones-data/internal/got-kings-server"
)

func main() {
	dbHandle := server.DatabaseHandle{nil}
	err := dbHandle.Connect("/got")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	gotServer := &server.Server{&dbHandle}
	twirpHanldler := gotKings.NewGotKingsServer(gotServer, nil)

	http.ListenAndServe(":8080", twirpHanldler)
}
