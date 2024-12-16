package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Структура для хранения данных о зарегистрированных пользователях
type ContestBot struct {
	registeredUsers map[int64]bool
	mu              sync.Mutex
}

func NewContestBot() *ContestBot {
	return &ContestBot{
		registeredUsers: make(map[int64]bool),
	}
}

func (cb *ContestBot) RegisterUser(userID int64) bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.registeredUsers[userID] {
		return false // Пользователь уже зарегистрирован
	}

	cb.registeredUsers[userID] = true
	return true // Новая регистрация
}

func (cb *ContestBot) IsUserRegistered(userID int64) bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	return cb.registeredUsers[userID]
}

func main() {
	// Укажите ваш токен и username канала
	bot, err := tgbotapi.NewBotAPI("7758148248:AAGPFJYGWj-puL619sOq76sELDUy1wFKous")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	contestBot := NewContestBot()
	channelLink := "https://t.me/rakhimov393" // Ссылка на ваш канал

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			userID := update.Message.From.ID
			text := update.Message.Text

			log.Printf("Received message from user %s: %s", update.Message.From.UserName, text)

			// Обработка команд
			if strings.HasPrefix(text, "/start") {
				sendMessage(bot, chatID, fmt.Sprintf("Привет! Добро пожаловать на конкурс.\nЧтобы участвовать, зарегистрируйтесь с помощью команды /register.\nТакже подпишитесь на канал: %s", channelLink))
			} else if strings.HasPrefix(text, "/register") {
				if contestBot.RegisterUser(userID) {
					sendMessage(bot, chatID, "Вы успешно зарегистрированы на конкурс!")
				} else {
					sendMessage(bot, chatID, "Вы уже зарегистрированы.")
				}
			} else if strings.HasPrefix(text, "/participate") {
				if contestBot.IsUserRegistered(userID) {
					sendMessage(bot, chatID, "Вы участвуете в конкурсе. Удачи!")
				} else {
					sendMessage(bot, chatID, fmt.Sprintf("Вы не зарегистрированы. Сначала используйте команду /register.\nНе забудьте подписаться на канал: %s", channelLink))
				}
			} else {
				sendMessage(bot, chatID, "Я понимаю только команды /start, /register и /participate.")
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
