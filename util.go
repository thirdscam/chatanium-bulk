package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func StringToUint(s string) (uint, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return 0, fmt.Errorf("invalid number format: %v", err)
	}

	if result < 0 {
		return 0, fmt.Errorf("number must be positive")
	}

	return uint(result), nil
}
