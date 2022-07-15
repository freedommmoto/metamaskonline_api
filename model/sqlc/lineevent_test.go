package model

import (
	"context"
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertLineEvent(t *testing.T) {
	user, err := testQueries.SelectUserID(context.Background(), int32(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, user.IDLine.String)

	arg := InsertLineEventParams{
		IDLineUser:       user.IDLine.String,
		IDUse:            user.IDUser,
		RequestLogEvent:  sql.NullString{String: tool.GetMockOneLineRquest(), Valid: true},
		ResponseLogEvent: sql.NullString{String: "{}", Valid: true},
		Error:            false,
		ErrorText:        sql.NullString{String: "", Valid: false},
	}
	//sent
	lineEvent, err := testQueries.InsertLineEvent(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, lineEvent)

	assert.Equal(t, lineEvent.IDLineUser, user.IDLine.String)
	assert.Equal(t, lineEvent.IDUse, user.IDUser)
	assert.Equal(t, lineEvent.Error, false)
	assert.NotEmpty(t, lineEvent.RequestLogEvent)

}
