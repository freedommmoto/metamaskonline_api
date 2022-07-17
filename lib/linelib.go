package lib

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"github.com/freedommmoto/metamaskonline_api/tool"
	uuid "github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func CheckProfileRegisterFromDB(Queries *db.Queries, profile ProFile) (db.User, bool, error) {
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

func CheckCodeWithIn3Hr(Queries *db.Queries, lineRequest *LineMessage) (bool, int32, error) {
	if len(lineRequest.Events) == 0 {
		return false, int32(0), errors.New("no event for CheckCodeWithIn3Hr")
	}
	lineOwnerValidation, err := Queries.SelectCodeUnConfirmWithIn3Houses(context.Background(), lineRequest.Events[0].Message.Text)
	if err != nil {
		return false, int32(0), err
	}
	return true, lineOwnerValidation.IDUser.Int32, nil
}

func CheckCodeWithUserProfile(Queries *db.Queries, lineRequest *LineMessage, userID int32) (bool, error) {
	userIDarg := sql.NullInt32{
		Int32: int32(userID),
		Valid: true,
	}

	lineOwnerValidation, err := Queries.SelectLastLineOwnerValidation(context.Background(), userIDarg)
	if err != nil {
		return false, err
	}
	log.Println("CheckCodeWithUserProfile CheckCodeWithUserProfile : !!!!", lineRequest.Events[0].Message.Text)
	log.Println("CheckCodeWithUserProfile CheckCodeWithUserProfile : !!!!", lineOwnerValidation.Code)
	if lineOwnerValidation.Code == lineRequest.Events[0].Message.Text {
		return true, nil
	}
	return false, nil
}

func UpdateUserOwnerValidation(Queries *db.Queries, userID int32) error {
	if userID < 1 {
		errors.New("call UpdateUserOwnerValidation with wrong format userid ")
	}
	user, err := Queries.UpdateUserOwnerValidation(context.Background(), userID)
	if err != nil {
		return err
	}
	if user.IDUser != userID {
		errors.New("id user that update OwnerValidation not correct")
	}
	return nil
}

func UpdateLineIdByWhereUserIDParams(Queries *db.Queries, lineid string, userID int32) error {
	arg := db.UpdateLineIdByWhereUserIDParams{
		IDLine: sql.NullString{String: lineid, Valid: true},
		IDUser: userID,
	}
	user, err := Queries.UpdateLineIdByWhereUserID(context.Background(), arg)
	if err != nil {
		log.Println("UpdateLineIdByWhereUserID :", err)
		return err
	}
	if user.IDLine.String != lineid {
		return errors.New("id line after save is not same with line id profile")
	}
	return nil
}

//func MakeNewCodeForNewUser(Queries *db.Queries, userID int32) (string, error) {
//	code := tool.RandomCodeNumber(4)
//	arg := db.InsertLineOwnerValidationParams{
//		IDUser: int32(userID),
//		Code:   code,
//	}
//	user, err := Queries.UpdateUserOwnerValidation(context.Background(), userID)
//	if err != nil {
//		return err
//	}
//	if user.IDUser != userID {
//		errors.New("id user that update OwnerValidation not correct")
//	}
//	err = mainQueries.UpdateLineIdByWhereUserID(context.Background(), profile.UserID, userIDFromCode)
//	if err != nil {
//		log.Println("UpdateLineIdByWhereUserID :", err)
//		return err
//	}
//	return nil
//}

//func MakeNewCodeForNewUser(Queries *db.Queries, userID int32) (string, error) {
//	code := tool.RandomCodeNumber(4)
//	arg := db.InsertLineOwnerValidationParams{
//		IDUser: int32(userID),
//		Code:   code,
//	}
//	lineOwner, err := Queries.InsertLineOwnerValidation(context.Background(), arg)
//	if err != nil {
//		return "", err
//	}
//	length := len([]rune(lineOwner.Code))
//	if length != 4 {
//		err = errors.New("code is not length 4")
//		return "", err
//	}
//	return lineOwner.Code, nil
//}

func CheckCodeFormatIs4Number(lineRequest *LineMessage) bool {
	message := lineRequest.Events[0].Message.Text
	length := len([]rune(message))
	if length != 4 {
		return false
	}
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		if !strings.Contains(tool.CodeNumber, string(runes[i])) {
			return false
		}
	}

	return true
}

func ReplyLineUserWithNormalText(messageStr string, ChannelToken string, lineRequest *LineMessage) error {

	message := Message{
		Type: "text",
		Text: messageStr,
	}
	replyRequest := ReplyRequest{
		ReplyToken: lineRequest.Events[0].ReplyToken,
		Messages: []Message{
			message,
		},
	}

	value, _ := json.Marshal(replyRequest)
	var jsonStr = []byte(value)

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//log.Println("response Status:", resp.Status)
	//log.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println("response Body:", string(body))
	return nil
}

func MakePushOneUserWithLineAPI(messageStr string, user db.User, ChannelToken string) error {
	if !user.IDLine.Valid {
		errors.New("no error line id for send !")
	}
	uuidData := uuid.New()
	uuidDataStr := uuidData.String()

	url := "https://api.line.me/v2/bot/message/push"
	message := Message{
		Type: "text",
		Text: messageStr,
	}
	MessagePush := PushRequest{
		To: user.IDLine.String,
		Messages: []Message{
			message,
		},
	}
	value, _ := json.Marshal(MessagePush)
	var jsonStr = []byte(value)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
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
