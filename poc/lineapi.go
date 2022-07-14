package main

import (
	"bytes"
	"encoding/json"
	uuid "github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type LineMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

//type LineMessagewebhook struct {
//	ReplyToken string `json:"replyToken"`
//	Type       string `json:"type""`
//	Timestamp  int64  `json:"timestamp""`
//	Source     struct {
//		Type   string `json:"type"`
//		UserID string `json:"text"`
//	} `json:"Source"`
//	Message struct {
//		ID   string `json:"id"`
//		Type string `json:"type"`
//		Text string `json:"text"`
//	} `json:"message"`
//}

type ReplyRequest struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
}

type PushRequest struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ProFile struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

var ChannelToken = "enter_you_line_api_token_here"

//patara
var testUserPatataID = "Ue5308cc32ee5ca607c596e87877715b6"

//🍄~ωiʑanee~🍄
var testUserWizaneID = "U0d8f460619ae601aed723b4cab9c856b"

func main() {
	e := echo.New()

	e.POST("/webhook", func(c echo.Context) error {
		err := replyMessageLine(c)
		if err != nil {
			log.Println("replyMessageLine error :", err)
			return c.String(http.StatusInternalServerError, "error")
		}
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/testpush", func(c echo.Context) error {
		err := SendPushMessageLine("test text")
		if err != nil {
			log.Println("SendPushMessageLine error :", err)
			return c.String(http.StatusInternalServerError, "error")
		}
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8888"))
}

func SendPushMessageLine(text string) error {

	uuidData := uuid.New()
	uuidDataStr := uuidData.String()

	url := "https://api.line.me/v2/bot/message/push"
	message := Message{
		Type: "text",
		Text: text,
	}
	MessagePush := PushRequest{
		To: testUserPatataID,
		Messages: []Message{
			message,
		},
	}
	value, _ := json.Marshal(MessagePush)
	var jsonStr = []byte(value)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("error after call NewRequest", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)
	req.Header.Add("X-Line-Retry-Key", uuidDataStr)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error after ReadAll", err)
		return err
	}
	bodyString := string(bodyBytes)
	log.Println("SendPushMessageLine retrun data : ", bodyString)

	return nil
}

func replyMessageLine(c echo.Context) error {

	log.Println("webhook")
	Line := new(LineMessage)
	if err := c.Bind(Line); err != nil {
		log.Println("err Bnnd LineMessage", err)
	}
	log.Println(Line)
	fullName := getProfile(Line.Events[0].Source.UserID)
	message := Message{
		Type: "text",
		Text: " you message is " + Line.Events[0].Message.Text + " welcome , " + fullName,
	}

	replyRequest := ReplyRequest{
		ReplyToken: Line.Events[0].ReplyToken,
		Messages: []Message{
			message,
		},
	}

	value, _ := json.Marshal(replyRequest)

	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("response Body:", string(body))

	return nil
}

func getProfile(userId string) string {

	url := "https://api.line.me/v2/bot/profile/" + userId

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var profile ProFile
	if err := json.Unmarshal(body, &profile); err != nil {
		log.Println("error Unmarshal ", err)
	}
	log.Println(profile.DisplayName)
	return profile.DisplayName

}
