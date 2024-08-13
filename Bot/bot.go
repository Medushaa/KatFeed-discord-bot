package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	//"strings"
	"github.com/bwmarrin/discordgo"
)

var BotToken string
const food int64 = 500 //grams of food that we have
const daily_rate int64 = 200 //grams per day

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// open session
	discord.Open()
	checkNilErr(err)
	defer discord.Close() // close session, after function termination

	// Register slash commands
	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "feed",
		Description: "Record amount of dry food (in grams) given to Momo and Rusty",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "amount",
				Description: "Amount of dry food (in grams)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
		},
	})
	checkNilErr(err)
	
	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "summary",
		Description: "Get the overall total amount of food fed and remaining.",
	})
	checkNilErr(err)

	// add a event handler
	discord.AddHandler(newMessage)
	discord.AddHandler(commandHandler) // Handle slash commands

	// keep bot running untill ctrl + C
	fmt.Println("KatFeed is Onlinee!!!!")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}