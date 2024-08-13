package bot

import (
	"os"
	"strings"
	"github.com/bwmarrin/discordgo"
)

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// prevent bot responding to its own message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// respond to user message if it contains `!help` or `!bye`
	switch {
	case strings.Contains(message.Content, "?katFeed"):
		embed := &discordgo.MessageEmbed{
			Color:       0xc2dbf9, // Light blue color
			Title:       "Hii! 😺",
			Description: "I am Mr.KatFeed and I am here to help you feed Momo and Rusty. Nyan 🩵",
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "attachment://cat-tongue.jpg", // URL of the image
			},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Mr. KatFeed",
				IconURL: discord.State.User.AvatarURL(""), // Bot's profile picture
			},
		}
		file, err := os.Open("assets/cat-tongue.jpg")
		checkNilErr(err)
		defer file.Close()

		_, err = discord.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
			Embed: embed,
			Files: []*discordgo.File{
				{
					Name:   "cat-tongue.jpg",
					Reader: file,
				},
			},
		})
		checkNilErr(err)
	case strings.Contains(message.Content, "?bye"):
		discord.ChannelMessageSend(message.ChannelID, "Good Bye👋")
		// add more cases if required
	}

}