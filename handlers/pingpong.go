package handlers

import (
	"fmt"

	"github.com/MumAroi/discord-bot-basic/libraries"
)

func PingPong(ctx libraries.Context) {

	// if len(ctx.Args) == 0 {
	// 	ctx.Reply("Add command usage")
	// 	return
	// }

	// if ctx.Message.Author.ID == ctx.Discord.State.User.ID {
	// 	return
	// }
	// if ctx.Message.Content != "ping" {
	// 	return
	// }

	_, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, "Pong!")
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return
	}

}
