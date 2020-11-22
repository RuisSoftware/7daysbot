package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Verify message is for this channel
	if m.ChannelID != config.Discord.Channel {
		return
	}

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// get message content
	content := m.Content

	// Check if this is a command and if so remove the prefix
	if strings.HasPrefix(content, config.Discord.Prefix) {
		content = strings.Replace(content, config.Discord.Prefix, "", -1)
	}

	// If the message is "ping" reply with "Pong!"
	if content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	// If the message is "pong" reply with "Ping!"
	if content == "info" {
		commands := config.Discord.Prefix + "info"
		commands += ", " + config.Discord.Prefix + "time"
		commands += ", " + config.Discord.Prefix + "players"
		_, _ = s.ChannelMessageSend(m.ChannelID, "Valid Commands: "+commands)
		return
	}

	if content == "time" {
		sendTelnet("gt")
		return
	}

	if content == "players" {
		sendTelnet("lp")
		return
	}

	// Nothing matched, treat as normal chat
	chatString := "say \"[" + m.Author.Username + "] " + content + "\""
	sendTelnet(chatString)
}

func sendDiscordMessage(msg string) {
	_, err := discord.ChannelMessageSend(config.Discord.Channel, msg)

	if err != nil {
		fmt.Println("Error sending message: ")
	}
}
