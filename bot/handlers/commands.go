package handlers

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "axiom-test",
		Description: "example command",
		Type:        discordgo.ChatApplicationCommand,
	},
}
