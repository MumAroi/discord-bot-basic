package helpers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	Context struct {
		Discord     *discordgo.Session
		Guild       *discordgo.Guild
		TextChannel *discordgo.Channel
		User        *discordgo.User
		Message     *discordgo.MessageCreate
		Args        []string
	}
)

func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel,
	user *discordgo.User, message *discordgo.MessageCreate, cmdHandler *CommandHandler,
) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	return ctx
}

func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, content)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}
