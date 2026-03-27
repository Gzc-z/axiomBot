package main

import (
	"log"
	"os"
	"os/signal"

	"axiom/bot"
	"axiom/bot/config"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discordBot := bot.NewBot(config.Load())
	ds := discordBot.Session

	ds.Identify.Intents |= discordgo.IntentMessageContent
	discordBot.SessionEvents()

	if err := ds.Open(); err != nil {
		panic(err)
	}
	defer ds.Close()

	log.Printf("Logged in as: %v#%v", ds.State.User.Username, ds.State.User.Discriminator)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	os.Exit(0)
}
