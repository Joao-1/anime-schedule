package main

import (
	"log"
	"os"
	"os/signal"

	. "github.com/Joao-1/animeSchedule/src"
)

func main() {
	bot := Bot{}
	bot.Start()
	bot.Login()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop 

	defer Session.Close()

	log.Println("Gracefully shutting down.")
}