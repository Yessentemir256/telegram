package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Укажите ваш токен
	bot, err := tgbotapi.NewBotAPI("7758148248:AAGPFJYGWj-puL619sOq76sELDUy1wFKous")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // Включить режим отладки

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настройка обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обработка обновлений
	for update := range updates {
		if update.Message != nil { // Если это сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Ответ пользователю
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, "+update.Message.From.FirstName+"!")
			bot.Send(msg)
		}
	}
}
