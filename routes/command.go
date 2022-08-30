package routes

import (
	"github.com/MumAroi/discord-bot-basic/handlers"
	"github.com/MumAroi/discord-bot-basic/libraries"
)

func RegisterCommands() *libraries.CommandHandler {
	commands := libraries.NewCommandHandler()
	commands.Register("pingpong", handlers.PingPong, "Ping Pong!")
	commands.Register("widget", handlers.WidgetDemo, "Widget!")
	return commands
}
