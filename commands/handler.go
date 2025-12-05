package commands

import "github.com/bwmarrin/discordgo"

func SlashHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "ping":
			pingHandler(s, i)
		case "echo":
			echoHandler(s, i)
		}
	}
}
