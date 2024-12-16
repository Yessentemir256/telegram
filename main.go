package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Specify your token
	bot, err := tgbotapi.NewBotAPI("7758148248:AAGPFJYGWj-puL619sOq76sELDUy1wFKous")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			text := update.Message.Text

			// Log received message
			log.Printf("Received message from user %s: %s", update.Message.From.UserName, text)

			if strings.HasPrefix(text, "/start") {
				sendMessage(bot, chatID, "Welcome! You have started the bot.")
			} else if strings.HasPrefix(text, "/help") {
				sendMessage(bot, chatID, "Here are the commands you can use:\n/start - Start the bot\n/help - Show this help message")
			} else {
				sendMessage(bot, chatID, "I only understand commands for now. Try /start or /help.")
			}
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
