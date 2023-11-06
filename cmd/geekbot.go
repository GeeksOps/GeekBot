/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
)

var (
	// Telegram bot token
	TeleToken = os.Getenv("TELE_TOKEN")
	// Add second variable
	seconds = 10
)

// geekbotCmd represents the geekbot command
var geekbotCmd = &cobra.Command{
	Use:     "geekbot",
	Aliases: []string{"start"},
	Short:   "KBot is a bot for Telegram",
	Long: `KBot is a fully functional bot for Telegram.

Created for fun and learning purposes. The main goal is to learn Go and Telegram Bot API.
It's a work in progress, so expect some bugs and missing features.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("geekbot %s started", appVersion)
		geekbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: time.Duration(seconds) * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check your TELE_TOKEN env variable, %s", err)
		}

		geekbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send((fmt.Sprintf("Hello! I'm geekbot %s", appVersion)))
			default:
				err = m.Send(("Sorry, I don't understand"))

			}

			return err
		})

		geekbot.Handle("/help", func(m telebot.Context) error {
			payload := m.Message().Payload

			log.Print(payload, m.Text())
			err = m.Send("How can I help you?")

			return err
		})

		geekbot.Handle("/echo", func(m telebot.Context) error {
			payload := m.Message().Payload

			log.Print(payload, m.Text())
			err = m.Send(payload)

			return err
		})

		geekbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(geekbotCmd)
}
