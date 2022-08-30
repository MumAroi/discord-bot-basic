package handlers

import (
	"fmt"

	"github.com/MumAroi/discord-bot-basic/libraries"
)

func PingPong(ctx libraries.Context) {

	_, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, "Pong!")
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return
	}

}
