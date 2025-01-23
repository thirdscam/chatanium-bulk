package actions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Delete struct {
	ChannelID string
	StartAt   string
	From      uint
	Count     uint
	Filter    string
}

func (t *Delete) Run(s *discordgo.Session) error {
	channel, err := s.Channel(t.ChannelID)
	if err != nil {
		return err
	}

	if channel.Type != discordgo.ChannelTypeGuildText {
		return fmt.Errorf("channel type is not text")
	}

	// find the message to start from
	var startID string
	if t.From == 0 {
		m, err := s.ChannelMessages(t.ChannelID, 1, "", "", "")
		if err != nil {
			return err
		}
		startID = m[0].ID
	} else {
		for i := uint(0); i < t.From; {
			limit := uint(100)
			if t.From-i < limit {
				limit = t.From - i
			}

			m, err := s.ChannelMessages(t.ChannelID, int(limit), t.StartAt, "", "")
			if err != nil {
				return err
			}
			startID = m[len(m)-1].ID
			i += limit
		}
	}

	// get the messages
	var rawMessages []*discordgo.Message
	for i := uint(0); i < t.From; {
		limit := uint(100)
		if t.From-i < limit {
			limit = t.From - i
		}

		m, err := s.ChannelMessages(t.ChannelID, int(limit), startID, "", "")
		if err != nil {
			return err
		}
		rawMessages = append(rawMessages, m...)
		startID = m[len(m)-1].ID
		i += limit
	}

	var messageIds []string
	for _, m := range rawMessages {
		messageIds = append(messageIds, m.ID)
	}

	s.ChannelMessagesBulkDelete(t.ChannelID, messageIds)

	return nil
}
