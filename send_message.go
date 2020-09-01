package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	apiMessagesEndpoint = "https://api-rest.zenvia.com/services/send-sms"
)

type sendMessageRequest struct {
	SendSmsRequest SendMessageRequest `json:"sendSmsRequest"`
}

type sendMessageResponse struct {
	SendSmsResponse sendMessageResponsePayload `json:"sendSmsResponse"`
}

type sendMessageResponsePayload struct {
	// TODO: verify this fields...
	StatusCode        string `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	DetailCode        string `json:"detailCode"`
	DetailDescription string `json:"detailDescription"`
}

type SendMessageRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

// SendMessage sent a message (SMS).
func (c *Client) SendMessage(input SendMessageRequest) error {
	if input.To == "" {
		return errors.New("Incorrect or incomplete 'to' mobile number")
	}

	if input.Msg == "" {
		return errors.New("Message body invalid")
	}

	request := &sendMessageRequest{input}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("when marshal JSON: %s", err.Error())
	}

	req, err := http.NewRequest("POST", apiMessagesEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("when creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.authorization)

	client := &http.Client{}
	apiResponse, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("when sending request: %s", err.Error())
	}

	defer apiResponse.Body.Close()

	if apiResponse.StatusCode != 200 {
		if apiResponse.StatusCode == 401 {
			return errors.New("Invalid credentials")
		}
		return errors.New(apiResponse.Status)
	}

	sendMessageResponse := &sendMessageResponse{}
	if err := json.NewDecoder(apiResponse.Body).Decode(sendMessageResponse); err != nil {
		return err
	}

	if sendMessageResponse.SendSmsResponse.StatusCode != "00" {
		return errors.New(sendMessageResponse.SendSmsResponse.DetailDescription)
	}

	return nil
}
