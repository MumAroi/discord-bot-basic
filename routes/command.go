package routes

import (
	"github.com/MumAroi/discord-bot-basic/handlers"
	"github.com/MumAroi/discord-bot-basic/models"
)

func RegisterCommands() *models.CommandHandler {
	commands := models.NewCommandHandler()
	commands.Register("pingpong", handlers.PingPong, "Ping Pong!")
	return commands
}
