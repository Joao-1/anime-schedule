package bot

import (
	"log"

	commands "github.com/Joao-1/animeSchedule/src/commands"
	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

type Bot struct {
	commands map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate)
}

func (bot *Bot) Start() {
	var err error
	Session, err = discordgo.New(("Bot " + "MTA5MjQ4OTgwMzgwMDUyNjk0OA.GM_b6c.HXqAl7E9spqZf36u1jhsxc4sWCAYYd8fh5SiCs"))
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	bot.LoadCommands()
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

func (bot *Bot) LoadCommands() {
	// this only a test. I will change this later, but it is looks good
	bot.commands = make(map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate))

	bot.commands["hello"] = commands.Hello.Execute

	// bot.commands["hello"]()

	for _, command := range commands.All {
		log.Println(command)
	}
}
