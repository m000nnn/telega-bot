package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const telegramBotToken = "7093264031:AAGhECj-whrLj9rWNgjXKWuxwEP6H0PZ2Ng"

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	updates := tgbotapi.NewUpdate(0)
	updates.Timeout = 60

	messages, err := bot.GetUpdatesChan(updates)
	if err != nil {
		log.Fatal(err)
	}
	permissionToRepeat := false
	for message := range messages {
		var reply string
		if permissionToRepeat {
			reply = message.Message.Text
		}
		log.Printf("[%s] %s", message.Message.From.UserName, message.Message.Text)

		switch message.Message.Command() {
		case "start_repeat":
			reply = "теперь я буду повторять за тобой"
			permissionToRepeat = true
		case "stop_repeat":
			reply = "теперь я не буду повторять за тобой"
			permissionToRepeat = false
		}

		answer := tgbotapi.NewMessage(message.Message.Chat.ID, reply)
		bot.Send(answer)
	}
}
