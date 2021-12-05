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

// CreateMessageReturn is to decode the json data
type CreateMessageReturn struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id                          int    `json:"id"`
			Title                       string `json:"title"`
			Type                        string `json:"type"`
			AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

// CreateMessage is to create a message with a bot in a Telegram chat
func CreateMessage(message, chatId, parseMode, apiToken string) (CreateMessageReturn, error) {

	// Get base url
	url := "https://api.telegram.org/bot" + apiToken + "/sendMessage"

	// Define client for request
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	// Parse url & add attributes
	parse := request.URL.Query()
	parse.Add("chat_id", chatId)
	parse.Add("parse_mode", parseMode)
	parse.Add("text", message)
	request.URL.RawQuery = parse.Encode()

	// Send request
	response, err := client.Do(request)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	// Decode response
	var decode CreateMessageReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	// Return data
	return decode, nil

}
