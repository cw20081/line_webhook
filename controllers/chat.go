package controllers

import (
	"log"
	"webhook/configs"
	"webhook/models"
	"webhook/repositories"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func SaveChat() gin.HandlerFunc {
	bot, bot_err := linebot.New(configs.GetConfig("line.channelSecret"), configs.GetConfig("line.channelAccessToken"))
	log.Println("Bot:", bot, " err:", bot_err)
	return func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			log.Fatal(err)
			if err == linebot.ErrInvalidSignature {
				c.String(400, "Error.")
			} else {
				c.String(500, "Error.")
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("user ID:"+event.Source.UserID+"\nmsg ID:"+message.ID+":"+"\nGet:"+message.Text)).Do(); err != nil {
						log.Print(err)
					}

					user, err := repositories.GetUser(event.Source.UserID)
					if err != nil {
						log.Print(err)
					}

					chat := models.Chat{
						UserId:  user.Id,
						Source:  "line",
						Type:    "recive",
						Message: message.Text,
					}

					_, saveChatErr := repositories.SaveChat(chat)
					if saveChatErr != nil {
						log.Print(saveChatErr)
					}

				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Please send Text.")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}

		c.JSON(200, gin.H{
			"message": "ok",
		})
	}

}

func SendChat() gin.HandlerFunc {
	bot, bot_err := linebot.New(configs.GetConfig("line.channelSecret"), configs.GetConfig("line.channelAccessToken"))
	log.Println("Bot:", bot, " err:", bot_err)
	return func(c *gin.Context) {
		json := map[string]string{}
		c.BindJSON(&json)

		if val, ok := json["message"]; ok {
			userID := c.Param("userID")
			if _, err := bot.PushMessage(userID, linebot.NewTextMessage(val)).Do(); err != nil {
				log.Print(err)
				c.JSON(500, gin.H{
					"message": "push Error.",
				})
			}
			user, err := repositories.GetUser(userID)
			if err != nil {
				log.Print(err)
			}

			chat := models.Chat{
				UserId:  user.Id,
				Source:  "line",
				Type:    "push",
				Message: val,
			}

			_, saveChatErr := repositories.SaveChat(chat)
			if saveChatErr != nil {
				log.Print(saveChatErr)
			}

			c.JSON(200, gin.H{
				"message": "ok",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "input Error.",
			})
		}
	}
}

func IndexChat() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		user, err := repositories.GetUser(userID)
		if err != nil {
			log.Print(err)
		}
		data := repositories.IndexChat(user)

		type Res struct {
			Data []models.Chat
		}

		res := Res{
			Data: data,
		}
		c.JSON(200, res)
	}
}
