package main

import (
	"github.com/theovidal/onyxcord"

	"github.com/BecauseOfProg/bop-le-bricoleur/commands"
)

func main() {
	bot := onyxcord.RegisterBot("Bop Le Bricoleur")

	// bot.RegisterCommand("archive", commands.Archive()) TODO: make permissions work in onyxcord
	bot.RegisterCommand("ping", commands.Ping())
	bot.RegisterCommand("poll", commands.Poll())
	// bot.RegisterCommand("reactionRole", reaction_role.ReactionRole())
	bot.RegisterCommand("weather", commands.Weather())

	/*bot.Client.AddHandler(func(_ *discordgo.Session, message *discordgo.MessageDelete) {
		handlers.ReactionRoleHandlerDelete(&bot, message)
	})
	bot.Client.AddHandler(func(_ *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
		handlers.ReactionRoleAdd(&bot, reaction)
	})
	bot.Client.AddHandler(func(_ *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
		handlers.ReactionRoleRemove(&bot, reaction)
	})*/

	bot.Run(true)
}
