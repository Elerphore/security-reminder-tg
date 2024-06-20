package telegram

import (
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegramm_bot *tgbotapi.BotAPI
)

func InitTelegrammBot(wg *sync.WaitGroup) {
	defer wg.Done()

	telegramm_bot = newTelegrammBot()

	listenToUpdates()
}

func newTelegrammBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}
