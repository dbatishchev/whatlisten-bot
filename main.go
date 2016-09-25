package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"io"
	"net/http"
	"fmt"
	"strings"
)

func downloadFromUrl(url string) (filename string, err error) {
	tokens := strings.Split(url, "/")
	filename = tokens[len(tokens)-1]

	fmt.Println("Downloading", url, "to", filename)

	// todo: create tmp dir
	// todo: add constants
	// todo: remove file
	// todo: add tests
	// todo: add error handling

	output, err := os.Create(filename)
	if err != nil {
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return
	}

	return
}

func main() {
	bot, err := tgbotapi.NewBotAPI("283253166:AAFOJjcto2Fa7NIU5sBEu3VEkxFc8ovuW94")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		artist := GetData(update.Message.Text)

		preview := artist.GetPreview()

		if preview != "" {
			downloadedFilename, err := downloadFromUrl(preview)

			if err != nil {
				println("Error in file downloading")
			}

			img := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, downloadedFilename)
			img.Caption = artist.Name
			img.ReplyToMessageID = update.Message.MessageID

			_, err = bot.Send(img)

			if err != nil {
				println(err.Error())
			}
		}
	}
}