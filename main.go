package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var BotToken = "REDACTED" 

func main() {
	bot, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Bot creation error:", err)
		return
	}

	bot.AddHandler(handleMessage)

	err = bot.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord:", err)
		return
	}

	fmt.Println("The bot is up and running!")
	select {}
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "/hello" {
		s.ChannelMessageSend(m.ChannelID, "Salem, Alem! I'm your statistical bot.")
	}
}
