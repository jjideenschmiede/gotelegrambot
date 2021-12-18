//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gotelegrambot.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gotelegrambot

import (
	"encoding/json"
	"net/http"
)

// WebhookBody is to decode the json data
type WebhookBody struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

// SetWebhookReturn is to decode the json data
type SetWebhookReturn struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

// SetWebhook is to create a new webhook url
func SetWebhook(webhookUrl, apiToken string) (SetWebhookReturn, error) {

	// Get base url
	url := baseUrl + apiToken + "/setWebhook"

	// Define client for request
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	// Parse url & add attributes
	parse := request.URL.Query()
	parse.Add("url", webhookUrl)
	request.URL.RawQuery = parse.Encode()

	// Send request
	response, err := client.Do(request)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	// Decode response
	var decode SetWebhookReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	// Return data
	return decode, nil

}
