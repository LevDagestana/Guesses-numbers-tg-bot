package main

import (
	"math/rand"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var chatid int64

const TOKEN = "6885783088:AAGE_TqpQUho--keXj5PJLdqN3h8I27ct2A"

func connect() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic("ERROR")
	}
}

func sendMessage(msg string) {

	msgConfig := tgbotapi.NewMessage(chatid, msg)

	bot.Send(msgConfig)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

var target int
var guesses int

func resset() {
	target = rand.Intn(100)

	guesses = 0
	sendMessage("Я загадал число от 0 до 100")
	sendMessage("У тебя есть 10 попыток угадать число")
}

func main() {

	rand.NewSource(time.Now().UnixNano())

	connect()
	u := tgbotapi.NewUpdate(0)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil && update.Message.Text == "/start" {

			chatid = update.Message.Chat.ID
			resset()

			continue
		}
		if update.Message != nil {
			guesses++
			input, err := strconv.Atoi(update.Message.Text)
			handleErr(err)

			if guesses >= 10 {
				sendMessage("Ты проиграл. Задуманное число - " + strconv.Itoa(target))

				resset()
				continue
			}
			if input > target {
				sendMessage("Меньше")
				sendMessage("У тебя осталось " + strconv.Itoa(10-guesses) + " попыток")
			} else if input < target {
				sendMessage("Больше")
				sendMessage("У тебя осталось " + strconv.Itoa(10-guesses) + " попыток")
			} else {
				sendMessage("Ты угадал")
				sendMessage("Тебе понадобилось " + strconv.Itoa(guesses) + " попыток")
				resset()
				continue
			}

		}
	}
}
