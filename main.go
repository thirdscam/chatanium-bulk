package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/thirdscam/chatanium-bulk/actions"
	"github.com/thirdscam/chatanium/src/Backends/Discord/Interface/Slash"
	"github.com/thirdscam/chatanium/src/Util/Log"
)

var MANIFEST_VERSION = 1

var (
	NAME       = "Bulk"
	BACKEND    = "discord"
	VERSION    = "0.0.1"
	AUTHOR     = "ANTEGRAL"
	REPOSITORY = "github:thirdscam/chatanium-bulk"
)

var DEFINE_SLASHCMD = Slash.Commands{
	{
		Name:        "bulk",
		Description: "Bulk actions to messages",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Select the action",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "delete",
						Value: "from",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "from",
				Description: "Starting point to apply the action (counting from the last message)",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "count",
				Description: "Number of messages to apply the action to from a starting point (from)",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "filter",
				Description: "Select the filter (see README for more info)",
				Required:    false,
			},
		},
	}: Bulk,
}

func Start() {
}

func Bulk(s *discordgo.Session, i *discordgo.InteractionCreate) {
	action := i.ApplicationCommandData().Options[0].StringValue()
	fromStr := i.ApplicationCommandData().Options[1].StringValue()
	countStr := i.ApplicationCommandData().Options[2].StringValue()
	filter := i.ApplicationCommandData().Options[3].StringValue()

	var from uint
	var err error
	if fromStr == "" {
		from = 0
	} else {
		from, err = StringToUint(fromStr)
		if err != nil {
			InteractionResponse(s, i, "**Invalid range!**")
			return
		}
	}

	count, err := StringToUint(countStr)
	if err != nil {
		InteractionResponse(s, i, "**Invalid range!**")
		return
	}

	Log.Verbose.Printf("[Bulk] ACTION: %s, RANGE: %d (from: %d), filter: %s", action, from, count, filter)

	switch action {
	case "delete":
		runner := actions.Delete{
			ChannelID: i.ChannelID,
			StartAt:   i.Member.GuildID,
			From:      from,
			Count:     count,
			Filter:    filter,
		}
		runner.Run(s)
		break
	default:
		InteractionResponse(s, i, "**Invalid action!**")
		break
	}
}
