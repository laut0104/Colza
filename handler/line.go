package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Line(c echo.Context) error {
	bot, err := linebot.New(
		// 	// os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		// 	// os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		"9c3e88fee360283fe43c9a4e6436a4ea",
		"QiLW1yvx5y2FevQEfcZB3NUOAqgo10ceSgS0KQmiz4oGTKYyvhwa8TMwtnHFcjl52y2OtzBu6Tw2a2+vZAcdCXta1VgSZM/qrwYcf7NNfHT4/PBpJ5R102QDdsLct98YqOslWOn4siMIHfy2+oBNqwdB04t89/1O/w1cDnyilFU=",
	)
	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Response().WriteHeader(400)
			return c.String(400, "Hello, World!")
		} else {
			c.Response().WriteHeader(500)
			return c.String(500, "Hello, World!")
		}
	}
	for _, event := range events {
		fmt.Printf("%v", event.Source)
		switch event.Type {
		case linebot.EventTypeFollow:
			message := "友達登録ありがとう！\n「問題」と送ってね"
			fmt.Printf("%v", event)
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
				log.Print(err)
				errmsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errmsg)).Do(); err != nil {
					log.Print(err)
				}
			}

		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "問題":
					reply := "お客さんは裸眼・メガネ・コンタクトのどれ？"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				case "裸眼", "メガネ", "コンタクト":
					reply := "年齢は40歳以下ですか？\n40歳以下の場合は「はい」、違う場合は「いいえ」と答えてください"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}

				case "はい", "いいえ":
					reply := "用途はなんですか？\n「運転」・「日常」・「パソコン」で答えてください"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				// case "いいえ":
				// 	reply := "用途はなんですか？\n「運転」・「日常」・「パソコン」で答えてください"
				// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
				// 		log.Print(err)
				// 	}
				case "運転":
					reply := "お客さんにお勧めするレンズは「遠用」です\n40代以上の場合「遠近」の可能性があります！"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				case "日常":
					reply := "お客さんにお勧めするレンズは「遠用」です\n40代以上の場合「中近」の可能性があります！"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				case "パソコン":
					reply := "お客さんにお勧めするレンズは「近用」です\n40代以上の場合「近近」の可能性があります！"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				default:
					reply := "正しく入力してください！"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
	return nil
}
