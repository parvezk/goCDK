package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

// The handler our lambda will route to when we need to register our new users

func (api *ApiHandler) RegisterUser(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("invalid request, the fields cannot be empty")
	}

	// We need to check if a user with the same username already exists in our DB
	doesUserExist, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("there was an error registering the user %w", err)
	}

	if doesUserExist {
		return fmt.Errorf("a user with that username already exists")
	}

	// we know that this user does not exist
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error inserting the user %w", err)
	}

	return nil
}
