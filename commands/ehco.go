package commands

import "github.com/bwmarrin/discordgo"

var EchoCommand = &discordgo.ApplicationCommand{
	Name:        "echo",
	Description: "Replies back with your input",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "text",
			Description: "Text to echo",
			Required:    true,
		},
	},
}

func echoHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	txt := i.ApplicationCommandData().Options[0].StringValue()
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: txt},
	})
}
