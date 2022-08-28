package routes

import (
	"github.com/MumAroi/discord-bot-basic/handlers"
	"github.com/MumAroi/discord-bot-basic/helpers"
)

func RegisterCommands() {
	commands := helpers.NewCommandHandler()
	commands.Register("help", handlers.PingPong, "Ping Pong!")
}
