package main

import (
	"encoding/base64"
	"errors"
)

// Client keeps information about a connection.
type Client struct {
	authorization string
}

// NewClient creates a new client connection.
func NewClient(user, password string) (*Client, error) {
	if user == "" {
		return nil, errors.New("Invalid user")
	}

	if password == "" {
		return nil, errors.New("Invalid password")
	}

	accessToken := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	authorization := "Basic " + accessToken
	return &Client{
		authorization: authorization,
	}, nil
}
