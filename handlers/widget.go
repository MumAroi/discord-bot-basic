package handlers

import (
	"time"

	"github.com/MumAroi/discord-bot-basic/libraries"
	"github.com/MumAroi/discord-bot-basic/libraries/widgets"
	"github.com/bwmarrin/discordgo"
)

func WidgetDemo(ctx libraries.Context) {
	p := widgets.NewPaginator(ctx.Discord, ctx.TextChannel.ID)

	// Add embed pages to paginator
	p.Add(&discordgo.MessageEmbed{Description: "Page one"},
		&discordgo.MessageEmbed{Description: "Page two"},
		&discordgo.MessageEmbed{Description: "Page three"})

	// Sets the footers of all added pages to their page numbers.
	p.SetPageFooters()

	// When the paginator is done listening set the colour to yellow
	p.ColourWhenDone = 0xffff

	// Stop listening for reaction events after five minutes
	p.Widget.Timeout = time.Minute * 5

	// Add a custom handler for the gun reaction.
	p.Widget.Handle("🔫", func(w *widgets.Widget, r *discordgo.MessageReaction) {
		ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, "Bang!")
	})

	p.Spawn()
}
