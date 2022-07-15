package controller

import (
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

	user, regis, err := lib.CheckProfileRegister(mainQueries, profile)
	if err != nil {
		log.Println("checkProfileRegister error :", err)
		return err
	}

	//regis = false // uncomment this line for test case 1
	//case 1 user not in database yet mean user not register with UI yet.
	if !regis {
		messageStr := tool.GetLineText(1)
		err = lib.ReplyLineUserWithNormalText(messageStr, ChannelToken, Line)
		if err != nil {
			log.Println("ReplyLineUserWithNormalText :"+messageStr, err)
			return err
		}
		return nil
	}

	//case 2 user register on UI but not send validation code to line group yet.
	if !user.OwnerValidation {
		codeCorrect, err := lib.CheckCodeWithUserProfile(mainQueries, Line, user.IDUser)

		if codeCorrect {
			//case 2.1 user send code correct format but not correct code ex : user send 2123 but in database is 8472
			err = lib.UpdateUserOwnerValidation(mainQueries, user.IDUser)
			if err != nil {
				log.Println("UpdateUserOwnerValidation :", err)
				return err
			}

			messageStr := tool.GetLineText(2)
			err = lib.ReplyLineUserWithNormalText(messageStr, ChannelToken, Line)
			if err != nil {
				log.Println("ReplyLineUserWithNormalText :"+messageStr, err)
				return err
			}
		} else {
			//case 2.2 user send code not correct format maybe send text for ask how system work maybe send a sicker
			isCorrectFormatCode := lib.CheckCodeFormatIs4Number(Line)
			if isCorrectFormatCode {
				messageStr := tool.GetLineText(3)
				err = lib.ReplyLineUserWithNormalText(messageStr, ChannelToken, Line)
				if err != nil {
					log.Println("ReplyLineUserWithNormalText :"+messageStr, err)
					return err
				}

				_, err := lib.MakeNewCodeForNewUser(mainQueries, user.IDUser)
				if err != nil {
					log.Println("error makeNewCodeForNewUser :", err)
					return err
				}
			} else {
				messageStr := tool.GetLineText(4)
				err = lib.ReplyLineUserWithNormalText(messageStr, ChannelToken, Line)
				if err != nil {
					log.Println("ReplyLineUserWithNormalText :"+messageStr, err)
					return err
				}
			}
		}

	} else {
		//case 3 user already register user no need to do any chat with line group only wait for alert from metamask action
		messageStr := tool.GetLineText(5)
		err = lib.ReplyLineUserWithNormalText(messageStr, ChannelToken, Line)
		if err != nil {
			log.Println("ReplyLineUserWithNormalText :"+messageStr, err)
			return err
		}
	}

	return nil
}
