package controllers

import (
	"log"
	"time"
	"webhook/configs"
	"webhook/models"
	"webhook/repositories"
	"webhook/responses"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveChat() gin.HandlerFunc {
	bot, bot_err := linebot.New(configs.GetConfig("line.channelSecret"), configs.GetConfig("line.channelAccessToken"))
	if bot_err != nil {
		log.Println("Bot:", bot, " err:", bot_err)
	}
	return func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			log.Fatal(err)
			if err == linebot.ErrInvalidSignature {
				c.JSON(400, gin.H{
					"message": "Invalid Signature Error.",
				})
			} else {
				c.JSON(500, gin.H{
					"message": "Error.",
				})
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					user, err := repositories.GetUser(event.Source.UserID)
					if err != nil {
						if err == mongo.ErrNoDocuments { //user not found,create one
							user = models.User{
								Id:      primitive.NewObjectID(),
								Sources: map[string]string{"line": event.Source.UserID},
							}

							if _, err := repositories.CreateUser(user); err != nil {
								log.Print(err)
								c.JSON(500, gin.H{
									"message": "Create User Error.",
								})
								return
							}
						} else {
							c.JSON(500, gin.H{
								"message": "Get User Error.",
							})
							return
						}
					}

					chat := models.Chat{
						UserId:  user.Id,
						Source:  "line",
						Type:    "recive",
						Message: message.Text,
						Time:    time.Now(),
					}

					if _, err := repositories.SaveChat(chat); err != nil {
						log.Print(err)
						c.JSON(500, gin.H{
							"message": "Save Chat Error.",
						})
						return
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("user ID:"+event.Source.UserID+"\nmsg ID:"+message.ID+":"+"\nGet:"+message.Text)).Do(); err != nil {
						log.Print(err)
						c.JSON(500, gin.H{
							"message": "Reply Chat Error.",
						})
						return
					}

				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Please send Texts.")).Do(); err != nil {
						log.Print(err)
						c.JSON(500, gin.H{
							"message": "Reply Chat Error.",
						})
						return
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
	if bot_err != nil {
		log.Println("Bot:", bot, " err:", bot_err)
	}
	return func(c *gin.Context) {
		json := map[string]string{}
		c.BindJSON(&json)

		if val, ok := json["message"]; ok {
			userID := c.Param("userID")

			user, err := repositories.GetUser(userID)
			if err != nil {
				log.Print(err)
				c.JSON(400, gin.H{
					"message": "User Not Found.",
				})
				return
			}

			chat := models.Chat{
				UserId:  user.Id,
				Source:  "line",
				Type:    "push",
				Message: val,
				Time:    time.Now(),
			}

			if _, saveChatErr := repositories.SaveChat(chat); saveChatErr != nil {
				log.Print(saveChatErr)
				c.JSON(500, gin.H{
					"message": "Save Chat Error.",
				})
				return
			}

			if _, err := bot.PushMessage(userID, linebot.NewTextMessage(val)).Do(); err != nil {
				log.Print(err)
				c.JSON(500, gin.H{
					"message": "push Error.",
				})
				return
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
			if err == mongo.ErrNoDocuments {
				c.JSON(400, gin.H{
					"message": "User Not Found.",
				})
				return
			} else {
				log.Fatal(err)
				c.JSON(500, gin.H{
					"message": "Get User Error.",
				})
				return
			}
		}
		data := repositories.IndexChat(user)

		res := responses.ChatIndexRes{
			Data: data,
		}

		c.JSON(200, res)
	}
}
