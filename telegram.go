package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func sendMessage(token string, chatID int64, message string) tgbotapi.Message {
	log.Println("Sending message to", chatID)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	msg := tgbotapi.NewMessage(chatID, message)
	if config.Telegram.ParseMode != "" {
		msg.ParseMode = config.Telegram.ParseMode
	}
	result, err := bot.Send(msg)
	if err != nil {
		log.Println("Error sending message to", chatID, "with message", message)
		log.Fatal("Error:", err.Error())
	}

	log.Println("Message successfully sent to", result.Chat.UserName)

	return result
}
