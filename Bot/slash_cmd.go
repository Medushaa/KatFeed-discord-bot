package bot

import (
	"fmt"
	// "log"
	// "os"
	// "os/signal"
	//"strings"
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "feed":
		amount := i.ApplicationCommandData().Options[0].IntValue()
		// Save the feed amount to a file
		recordFeed(amount)
		todays_feed := dailyFeed() //get today's feed amount

		// Respond to the user
		embed := &discordgo.MessageEmbed{
			Color:       0xc2dbf9, // Light blue color
			Title:       "Feed Recorded 🍿",
			Description: fmt.Sprintf("%dg of dry food was placed in Momo and Rusty's bowl. \nHappy nomnoming!! 😋\nDon't forget to add water!!! 🌊\n \nToday's total feed = %dg", amount, todays_feed),
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
		checkNilErr(err)

	case "summary":
		total := calculateTotalFeed()
		left := food - total
		days_left := int64(math.Ceil(float64(left) / float64(daily_rate)))
		// Respond with an embed
		embed := &discordgo.MessageEmbed{
			Color:       0x4df7ad, // Light blue color
			Title:       "Total Feed Summary 🗒️",
			Description: fmt.Sprintf("The kitties ate %dg of dry food in total. \nSo, we have remaining %dg of food left.\nThis could last us %d more day.", total, left, days_left),
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
		checkNilErr(err)
	}
}

func recordFeed(amount int64) {
	file, err := os.OpenFile("feed_data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkNilErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		time.Now().Format("2006-01-02"), // Date in YYYY-MM-DD format
		time.Now().Format("15:04:05"),   // Time in HH:MM:SS format
		strconv.FormatInt(amount, 10),
	}

	err = writer.Write(record)
	checkNilErr(err)
}

func dailyFeed() int64 {
	var total int64 = 0

	file, err := os.Open("feed_data.csv")
	checkNilErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkNilErr(err)

	today := time.Now().Format("2006-01-02") // Get today's date in YYYY-MM-DD format

	for _, record := range records {
		date := record[0]
		amount, err := strconv.ParseInt(record[2], 10, 64)
		checkNilErr(err)

		if date == today {
			total += amount
		}
	}
	return total
}

func calculateTotalFeed() int64 {
	var total int64 = 0

	file, err := os.Open("feed_data.csv")
	checkNilErr(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkNilErr(err)

	for _, record := range records {
		amount, err := strconv.ParseInt(record[2], 10, 64)
		checkNilErr(err)

		total += amount
	}

	return total
}
