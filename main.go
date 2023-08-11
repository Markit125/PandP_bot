package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	errs "pandp/src/errors"
	rf "pandp/src/readFiles"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegramBotToken string

	wisdoms        []string
	originals      []string
	countOfWisdoms int

	lastIndexOfWisdom int = 0
)

func initArgs() {

	var fileWithWisdomName string
	var fileWithOriginalsName string
	var absent *errs.AbsentFileError

	telegramBotToken, _ = os.LookupEnv("TOKEN")
	log.Print("token in initArgs: ", telegramBotToken)

	if telegramBotToken == "" {
		log.Print("telegram-bot token is required")
		os.Exit(1)
	}

	fileWithWisdomName, _ = os.LookupEnv("WISDOM")
	err := rf.ReadFileInSliceFromFile(&wisdoms, fileWithWisdomName)
	if errors.As(err, &absent) {
		log.Print("File with wisdom is required")
		os.Exit(2)
	}

	if err != nil {
		log.Printf("%v", err.Error())
		os.Exit(3)
	}

	fileWithOriginalsName, _ = os.LookupEnv("ORIGINALS")
	err = rf.ReadFileInSliceFromFile(&originals, fileWithOriginalsName)
	if errors.As(err, &absent) {
		log.Print("File with originals is required")
		os.Exit(4)
	}

	if err != nil {
		log.Print("%w", err.Error())
		os.Exit(5)
	}

	countOfWisdoms = len(wisdoms)
	if len(originals) != countOfWisdoms {
		log.Print("Wisdom file and Originals file do not match")
		os.Exit(6)
	}

	if countOfWisdoms < 1 {
		log.Printf("Wisdom file is out of wisdoms")
		os.Exit(7)
	}
}

func main() {

	initArgs()
	log.Print(telegramBotToken)

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		reply := "Can only use command /wisdom and /original"

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		rand.Seed(time.Now().UnixNano())

		switch update.Message.Command() {
		case "start":
			reply = "PandP /wisdom /original"
		case "wisdom":
			lastIndexOfWisdom = rand.Intn(countOfWisdoms)
			reply = wisdoms[lastIndexOfWisdom]
		case "original":
			reply = originals[lastIndexOfWisdom]
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		bot.Send(msg)
	}
}
