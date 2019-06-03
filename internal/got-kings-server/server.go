package server

import (
	"context"

	gotKings "github.com/dennisssdev/game-of-thones-data/rpc/got-kings"
)

// Server is the entry point for creating the handler for user game of thrones requests
type Server struct {
	DBHandle *DatabaseHandle
}

type KingRequest struct {
	name  string `db:"name"`
	house string `db:"house"`
}

// GetCharacter returns information on a given king
func (s *Server) GetCharacter(ctx context.Context, charReq *gotKings.CharacterReq) (charInfo *gotKings.CharacterInfo, err error) {
	return nil, nil
}

// ClashKings attempts to make 2 kings attack each other which will determine what happens to them
func (s *Server) ClashKings(ctx context.Context, scenario *gotKings.Scenario) (result *gotKings.GOTResult, err error) {
	return nil, nil
}

// MakeKing is the entry to creating a new Westeros King
func (s *Server) MakeKing(ctx context.Context, charInfo *gotKings.CharacterInfo) (success *gotKings.BoolResult, err error) {
	// @todo : this should be inside the DBHandle and prevent the outside from calling this directly from the DB
	success = &gotKings.BoolResult{Result: false}
	//kingReq := KingRequest{name: charInfo.FullName, house: charInfo.Description}
	s.DBHandle.DB.NamedExec("INSERT INTO king (full_name, house) VALUES (:name, :house)", &KingRequest{name: "NotFriend", house: "Lannister"})
	//_, err = s.DBHandle.DB.Exec(`INSERT INTO king (full_name, house) VALUES (:name, :house)`, &KingRequest{name: "Friend", house: "Lannister"}) // just a test
	if err != nil {
		success.Result = false
	} else {
		success.Result = true
	}
	return success, err
}

// GetHistory attempts to retreive the kings action history
func (s *Server) GetHistory(ctx context.Context, format *gotKings.BoolResult) (result *gotKings.ScenarioResult, err error) {
	return nil, nil
}
