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
)

// UpdatesReturn is to decode the json data
type UpdatesReturn struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateId     int `json:"update_id"`
		MyChatMember struct {
			Chat struct {
				Id                          int    `json:"id"`
				Title                       string `json:"title"`
				Type                        string `json:"type"`
				AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
			} `json:"chat"`
			From struct {
				Id        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Date          int `json:"date"`
			OldChatMember struct {
				User struct {
					Id        int64  `json:"id"`
					IsBot     bool   `json:"is_bot"`
					FirstName string `json:"first_name"`
					Username  string `json:"username"`
				} `json:"user"`
				Status string `json:"status"`
			} `json:"old_chat_member"`
			NewChatMember struct {
				User struct {
					Id        int64  `json:"id"`
					IsBot     bool   `json:"is_bot"`
					FirstName string `json:"first_name"`
					Username  string `json:"username"`
				} `json:"user"`
				Status string `json:"status"`
			} `json:"new_chat_member"`
		} `json:"my_chat_member,omitempty"`
		Message struct {
			MessageId int `json:"message_id"`
			From      struct {
				Id        int    `json:"id"`
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
			Date                int `json:"date"`
			LeftChatParticipant struct {
				Id        int64  `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"left_chat_participant,omitempty"`
			LeftChatMember struct {
				Id        int64  `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"left_chat_member,omitempty"`
			NewChatParticipant struct {
				Id        int64  `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"new_chat_participant,omitempty"`
			NewChatMember struct {
				Id        int64  `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"new_chat_member,omitempty"`
			NewChatMembers []struct {
				Id        int64  `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"new_chat_members,omitempty"`
		} `json:"message,omitempty"`
	} `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

// Updates is to get the updates from telegram
func Updates(r Request) (UpdatesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/getUpdates",
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdatesReturn{}, err
	}

	// Decode response
	var decode UpdatesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return UpdatesReturn{}, err
	}

	// Return data
	return decode, nil

}
