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

	profile, err := lib.GetProfile(ChannelToken, Line.Events[0].Source.UserID)
	if err != nil {
		log.Println("getProfile error :", err)
		return err
	}

	user, regis, err := lib.CheckProfileRegister(mainQueries, profile)
	if err != nil {
		log.Println("checkProfileRegister error :", err)
		return err
	}

	log.Println("ืทำถึงตรงนี้ user is  :", user)
	return nil

	if !regis {
		//reply you need to rigsiter first before used this channal
		return nil
	}

	log.Println("end event here ", user)
	return nil

	if !user.OwnerValidation {

		//if text code is same on validation

		//reply you need to validation code you can see code in system dashboad
		return nil
	}

	//check is validation
	//check is not validation

	log.Println("profile", profile)
	return nil
}
