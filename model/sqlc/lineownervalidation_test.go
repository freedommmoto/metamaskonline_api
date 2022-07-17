package model

import (
	"context"
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeLineOwnerValidation(t *testing.T) (string, int32) {
	code := tool.RandomCodeNumber(4)
	IDUser := int32(1)
	arg := InsertLineOwnerValidationParams{
		Code:   code,
		IDUser: IDUser,
	}
	lineValidation, err := testQueries.InsertLineOwnerValidation(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, lineValidation)

	assert.Equal(t, lineValidation.Code, code)
	assert.Equal(t, lineValidation.IDUser, IDUser)
	return code, lineValidation.IDLineOwnerValidation
}

func makeUpdateLineOwnerValidation(t *testing.T) LineOwnerValidation {
	_, idLineOwnerValidation := makeLineOwnerValidation(t)
	userID := int32(1)
	arg := UpdateUserIDtoLineOwnerValidationParams{
		IDUser:                userID,
		IDLineOwnerValidation: int32(idLineOwnerValidation),
	}
	updateLineValidation, err := testQueries.UpdateUserIDtoLineOwnerValidation(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, updateLineValidation)

	assert.Equal(t, updateLineValidation.IDUser, userID)
	assert.Equal(t, updateLineValidation.IDLineOwnerValidation, idLineOwnerValidation)
	return updateLineValidation
}

func TestInsertLineOwnerValidation(t *testing.T) {
	makeLineOwnerValidation(t)
}

func TestUpdateLineOwnerValidation(t *testing.T) {
	makeUpdateLineOwnerValidation(t)
}

func TestSelectCodeUnConfirmWithIn3Houses(t *testing.T) {
	code, _ := makeLineOwnerValidation(t)
	//fmt.Println(code)
	//fmt.Println(time.Now().Zone())
	//fmt.Println(time.Now())

	lineValidation, err := testQueries.SelectCodeUnConfirmWithIn3Houses(context.Background(), code)
	assert.NoError(t, err)
	assert.NotEmpty(t, lineValidation)
	assert.Equal(t, code, lineValidation.Code)
}

func TestDeleteLineOwnerValidation(t *testing.T) {
	_, idLineOwnerValidation := makeLineOwnerValidation(t)
	LineOwnerValidation, err := testQueries.DeleteLineOwnerValidation(context.Background(), int32(idLineOwnerValidation))
	assert.NoError(t, err)

	LineOwnerValidation, errSelect := testQueries.SelectLineOwnerValidation(context.Background(), int32(idLineOwnerValidation))
	assert.EqualError(t, errSelect, sql.ErrNoRows.Error())
	assert.Empty(t, LineOwnerValidation)
}
