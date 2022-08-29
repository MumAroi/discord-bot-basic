package libraries

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	Prefix string
	BotId  string
	CMDH   *CommandHandler
}

func NewHandler(prefix string, botId string, cmdh *CommandHandler) *Handler {
	handler := new(Handler)
	handler.Prefix = prefix
	handler.BotId = botId
	handler.CMDH = cmdh
	return handler
}

func (h Handler) GetHandlers(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == h.BotId || user.Bot {
		return
	}

	content := message.Content
	if len(content) <= len(h.Prefix) {
		return
	}

	if content[:len(h.Prefix)] != h.Prefix {
		return
	}
	content = content[len(h.Prefix):]
	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := h.CMDH.Get(name)
	if !found {
		return
	}

	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return
	}
	guild, err := discord.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return
	}
	ctx := NewContext(discord, guild, channel, user, message, h.CMDH)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)
}
