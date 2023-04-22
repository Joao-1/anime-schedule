package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)


type Bot struct {}
var Session *discordgo.Session

func (bot *Bot) Start() {
	var err error
	Session, err = discordgo.New(("Bot " + "MTA5MjQ4OTgwMzgwMDUyNjk0OA.GvRj4s.jCPn1qBuRm9mfkPENqal6_OcmmeIvswrNcWVJI"))
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}
}

func (bot *Bot)  Login() {
	Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := Session.Open()
	if err != nil {
		log.Fatalf("Error login in Discord bot account: %v", err)
	}
}

