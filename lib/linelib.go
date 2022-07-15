package lib

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/metamaskonline?sslmode=disable"
)

func ValidationLineRequest(lineRequest *LineMessage) error {
	if len(lineRequest.Events) < 1 {
		return errors.New("Events nil")
	}

	if lineRequest.Events[0].Source.UserID == "" {
		return errors.New("user id from line request is nil")
	}
	return nil
}

func GetProfile(ChannelToken, userId string) (ProFile, error) {
	var profile ProFile
	url := "https://api.line.me/v2/bot/profile/" + userId

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		statusCodeStr := strconv.Itoa(res.StatusCode)
		return profile, errors.New("call get profile got http status : " + statusCodeStr)
	}

	if err != nil {
		return profile, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return profile, err
	}

	if err := json.Unmarshal(body, &profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func CheckProfileRegister(Queries *db.Queries, profile ProFile) (db.User, bool, error) {
	var err error
	var userNil db.User

	log.Println("CheckProfileRegister ", profile.UserID)
	userFromDB, err := Queries.SelectUserByLineUserID(context.Background(), sql.NullString{String: profile.UserID, Valid: true})

	if err != nil {
		return userNil, false, err
	}

	if userFromDB.IDUser < 1 {
		return userNil, false, nil
	}

	return userFromDB, true, nil
}
