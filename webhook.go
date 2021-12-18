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
	"fmt"
	"net/url"
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
func SetWebhook(webhookUrl string, r Request) (SetWebhookReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/setWebhook",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	newUrl.Add("url", webhookUrl)

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode response
	var decode SetWebhookReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SetWebhookReturn{}, err
	}

	// Return data
	return decode, nil

}
