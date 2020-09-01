package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	zenviaUser     = os.Getenv("ZENVIA_USER")
	zenviaPassword = os.Getenv("ZENVIA_PASSWORD")
	validPhone     = os.Getenv("VALID_PHONE")
)

func TestValidSendMessage(t *testing.T) {
	client, err := NewClient(zenviaUser, zenviaPassword)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	input := SendMessageRequest{
		To:  validPhone,
		Msg: "Message1 from SDK!",
	}
	err = client.SendMessage(input)
	assert.NoError(t, err)
}

func TestValidSendMessageWitCustomFrom(t *testing.T) {
	client, err := NewClient(zenviaUser, zenviaPassword)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	input := SendMessageRequest{
		From: "Me",
		To:   validPhone,
		Msg:  "Message2 from SDK!",
	}
	err = client.SendMessage(input)
	assert.NoError(t, err)
}

func TestSendMessageWithInvalidCredentials(t *testing.T) {
	client, err := NewClient("UUUSSSEEERRR", "AAABBBCCC")
	assert.NoError(t, err)
	assert.NotNil(t, client)
	input := SendMessageRequest{
		To:  validPhone,
		Msg: "Message3 from SDK!",
	}
	err = client.SendMessage(input)
	assert.EqualError(t, err, "Invalid credentials")
}

func TestSendMessageWithEmptyUser(t *testing.T) {
	client, err := NewClient("", "AAABBBCCC")
	assert.Nil(t, client)
	assert.EqualError(t, err, "Invalid user")
}

func TestSendMessageWithEmptyPassword(t *testing.T) {
	client, err := NewClient("UUUSSSEEERRR", "")
	assert.Nil(t, client)
	assert.EqualError(t, err, "Invalid password")
}

func TestValidSendMessageWithEmptyTo(t *testing.T) {
	client, err := NewClient(zenviaUser, zenviaPassword)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	input := SendMessageRequest{
		To:  "",
		Msg: "Message4 from SDK!",
	}
	err = client.SendMessage(input)
	assert.EqualError(t, err, "Incorrect or incomplete 'to' mobile number")
}

func TestValidSendMessageWithInvalidTo(t *testing.T) {
	client, err := NewClient(zenviaUser, zenviaPassword)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	input := SendMessageRequest{
		From: "aaaaa",
		To:   "5551123456789",
		Msg:  "Message5 from SDK!",
	}
	err = client.SendMessage(input)
	assert.EqualError(t, err, "Incorrect or incomplete 'to' mobile number")
}

func TestValidSendMessageWithEmptyText(t *testing.T) {
	client, err := NewClient(zenviaUser, zenviaPassword)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	input := SendMessageRequest{
		From: "aaaaa",
		To:   validPhone,
		Msg:  "",
	}
	err = client.SendMessage(input)
	assert.EqualError(t, err, "Message body invalid")
}

func TestValidSendMessageWithBigMessage(t *testing.T) {
	client, err := NewClient(zenviaUser, zenviaPassword)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	message := strings.Repeat("1234567890", 20)

	input := SendMessageRequest{
		From: "aaaaa",
		To:   validPhone,
		Msg:  message,
	}
	err = client.SendMessage(input)
	assert.EqualError(t, err, "Message body invalid")
}
