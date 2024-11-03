package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var moodEntries []string

func main() {
	// Получаем токен из переменной окружения
	token := "8002123865:AAF8x9X4g_PVSkHKFSfYTsebraBZL-KAKmA" // Замените на свой токен

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60})

	for update := range updates {
		if update.Message == nil { // игнорируем не текстовые сообщения
			continue
		}

		switch {
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в GriBotMoodLev! Используйте /log для записи своего настроения и /list для просмотра записей.")
				bot.Send(msg)
			case "log":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ваше текущее настроение:")
				bot.Send(msg)
				go handleMoodInput(update.Message.Chat.ID, bot)
			case "list":
				listMoods(bot, update.Message.Chat.ID)
			}
		}
	}
}

// handleMoodInput обрабатывает ввод настроения
func handleMoodInput(chatID int64, bot *tgbotapi.BotAPI) {
	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60})

	for update := range updates {
		if update.Message == nil || update.Message.Chat.ID != chatID {
			continue
		}

		moodEntry := update.Message.Text
		moodEntries = append(moodEntries, moodEntry)

		msg := tgbotapi.NewMessage(chatID, "Ваше настроение записано: "+moodEntry)
		bot.Send(msg)
		break // Прекращаем слушать дальнейшие сообщения для этого ввода
	}
}

// listMoods отображает список всех записей настроения
func listMoods(bot *tgbotapi.BotAPI, chatID int64) {
	if len(moodEntries) == 0 {
		msg := tgbotapi.NewMessage(chatID, "У вас нет записей о настроении.")
		bot.Send(msg)
		return
	}

	var sb strings.Builder
	sb.WriteString("Ваши записи о настроении:\n")
	for i, mood := range moodEntries {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, mood))
	}

	msg := tgbotapi.NewMessage(chatID, sb.String())
	bot.Send(msg)
}

