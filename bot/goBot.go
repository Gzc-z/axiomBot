// Package bot
package bot

import (
	"log"
	"os"

	"axiom/bot/config"
	"axiom/bot/handlers"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	GuildID string
}

func NewBot(cfg config.Config) *Bot {
	bot, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatal("something's wrong, can't create discord bot")
		os.Exit(1)
	}

	return &Bot{
		Session: bot,
		GuildID: cfg.GuildID,
	}
}

func (bot *Bot) SessionEvents() {
	ds := bot.Session
	discordHandlers := []any{
		handlers.MessageCreate,
	}
	for _, handler := range discordHandlers {
		ds.AddHandler(handler)
	}
}
