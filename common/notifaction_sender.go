package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type DiscordNotificationObject struct {
	Username  string `json:"username"`
	Content   string `json:"content"`
	AvatarURL string `json:"avatar_url"`
	// Embeds max is 10
	Embeds []DiscordNotificationEmbed `json:"embeds"`
}

type DiscordNotificationEmbed struct {
	Title       string                     `json:"title"`
	Author      string                     `json:"author"`
	URL         string                     `json:"url"`
	Description string                     `json:"description"`
	Color       int                        `json:"color"`
	Fields      []DiscordNotificationField `json:"fields"`
}

type DiscordNotificationField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	// if true then sets field objects in same line
	Inline bool `json:"inline"`
}

func SendDiscordWebHookMsg(notify DiscordNotificationObject) error {

    // You should make the webhook url private and not hard coding in your code.    
    // i.e. store in database or setup environment
	webhookURL := "https://discord.com/api/webhooks/example"

	jsonStringByte, err := json.Marshal(notify)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonStringByte))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return errors.New("Send request failed")
	}
	return nil
}
