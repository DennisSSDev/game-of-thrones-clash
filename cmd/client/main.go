package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	gotKings "github.com/dennisssdev/game-of-thones-data/rpc/got-kings"
)

func main() {
	client := gotKings.NewGotKingsProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeKing(context.Background(), &gotKings.CharacterInfo{
		FullName:      "Arya",
		Description:   "Stark",
		ConqueredLand: []string{"King's Landing"},
		Alliances:     []string{"Greyjoy", "NotGreyjoy"},
	})

	if err != nil {
		fmt.Printf("oh no: %v", err)
		os.Exit(1)
	}
	fmt.Printf("I have a nice new hat: %+v", hat)
}
