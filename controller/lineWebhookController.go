package controller

import (
	"context"
	"database/sql"
	lib "github.com/freedommmoto/metamaskonline_api/lib"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"github.com/freedommmoto/metamaskonline_api/tool"
	echo "github.com/labstack/echo/v4"
	"log"
)

var ChannelToken string

func ReplyMessageLine(c echo.Context, mainQueries *db.Queries, config tool.ConfigObject) error {
	ChannelToken = config.LINEToken
	Line := new(lib.LineMessage)
	if err := c.Bind(Line); err != nil {
		log.Println("err Bind LineMessage", err)
		return err
	}

	err := lib.ValidationLineRequest(Line)
	if err != nil {
		log.Println("validationLineRequest error :", err)
		return err
	}

	//check with line api make sure is not random lineid in request
	profile, err := lib.GetProfile(ChannelToken, Line.Events[0].Source.UserID)
	if err != nil || profile.UserID == "" {
		log.Println("getProfile error :", err)
		return err
	}

	user, _, err := lib.CheckProfileRegisterFromDB(mainQueries, profile)
	if err != nil && err != sql.ErrNoRows {
		log.Println("checkProfileRegister error :", err)
		return err
	}
	isCorrectFormatCode := lib.CheckCodeFormatIs4Number(Line)

	if user.OwnerValidation && user.IDUser > 0 {
		//case 3 user active no need to do anything
		sendTextToLine(4, ChannelToken, Line)
		return nil
	}

	if isCorrectFormatCode {
		//case 2 check code
		isCodeCorrect, userIDFromCode, idLineOwnerValidation, _ := lib.CheckCodeWithIn3Hr(mainQueries, Line)
		if isCodeCorrect && userIDFromCode > 0 {

			//update user owner validation
			err = lib.UpdateUserOwnerValidation(mainQueries, userIDFromCode)
			if err != nil {
				log.Println("UpdateUserOwnerValidation :", err)
				return err
			}
			err = lib.UpdateLineIdByWhereUserIDParams(mainQueries, profile.UserID, userIDFromCode)
			if err != nil {
				log.Println("UpdateLineIdByWhereUserIDParams :", err)
				return err
			}

			//error code that have been used map user with code and line id
			_, errDeleteCode := mainQueries.DeleteLineOwnerValidation(context.Background(), idLineOwnerValidation)
			if errDeleteCode != nil {
				log.Println("DeleteLineOwnerValidation :", errDeleteCode)
				return errDeleteCode
			}

			//done update active user then update
			sendTextToLine(3, ChannelToken, Line)
			return nil

		} else {
			sendTextToLine(2, ChannelToken, Line)
			return nil
		}
	} else {
		//case 1 ask user to register first
		sendTextToLine(1, ChannelToken, Line)
		return nil
	}

	return nil
}

func sendTextToLine(caseID int, channelToken string, line *lib.LineMessage) {
	messageStr := tool.GetLineText(caseID)
	err := lib.ReplyLineUserWithNormalText(messageStr, channelToken, line)
	if err != nil {
		log.Println("ReplyLineUserWithNormalText :"+messageStr, err)
	}
}
