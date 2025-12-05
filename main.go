package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Gary05CX/eat-what-7/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalf("No DISCORD_TOKEN provided in .env")
	}
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Register the interaction handler
	dg.AddHandler(commands.SlashHandler)
	dg.Identify.Intents = discordgo.IntentsGuilds
	dg.Open()
	defer dg.Close()

	// Register both slash commands at startup
	applications := []*discordgo.ApplicationCommand{
		commands.PingCommand,
		commands.EchoCommand,
	}
	for _, cmd := range applications {
		dg.ApplicationCommandCreate(dg.State.User.ID, "", cmd)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
