package event

import (
	"log"

	. "github.com/Joao-1/animeSchedule/src"
	"github.com/bwmarrin/discordgo"
)

func Ready() {
	Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
}