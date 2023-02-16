package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", hello)
	e.POST("/callback", line)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// conn, err := pgx.Connect(context.Background(), "postgres://root:password@postgres:5432/randomcooking")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	return c.String(http.StatusOK, "Hello, World!")

}

func line(c echo.Context) error {
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
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				replyMessage := fmt.Sprintf(
					"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	return c.String(http.StatusOK, "Hello, World!")
}
