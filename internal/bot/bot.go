package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)


type Bot struct {
	Session *discordgo.Session
}

// NewBot создаёт новый экземпляр бота с заданным токеном.
func NewBot(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &Bot{Session: session}, nil
}

// Start запускает сессию бота и начинает прослушивание событий.
func (b *Bot) Start() error {
	err := b.Session.Open()
	if err != nil {
		return err
	}
	log.Println("Bot is up and running!")
	return nil
}


func (b *Bot) Stop() {
	b.Session.Close()
}
