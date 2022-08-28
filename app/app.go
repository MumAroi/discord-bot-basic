package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MumAroi/discord-bot-basic/config"
	"github.com/MumAroi/discord-bot-basic/models"
	"github.com/MumAroi/discord-bot-basic/routes"
	"github.com/bwmarrin/discordgo"
)

var (
	BOTTOKEN string
	PREFIX   string
)

func init() {
	_config := config.LoadConfig()
	BOTTOKEN = _config.BotToken
	PREFIX = _config.Prefix
}

func Run() {

	dg, err := discordgo.New("Bot " + BOTTOKEN)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	usr, err := dg.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}

	botId := usr.ID

	commands := routes.RegisterCommands()
	commandHandler := models.NewHandler(PREFIX, botId, commands)

	dg.AddHandler(commandHandler.GetHandlers)

	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	dg.Close()
}
