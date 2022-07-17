package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/freedommmoto/metamaskonline_api/tool"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func makeLineOwnerValidation(t *testing.T) (string, int32) {
	code := tool.RandomCodeNumber(4)
	lineValidation, err := testQueries.InsertLineOwnerValidation(context.Background(), code)
	assert.NoError(t, err)
	assert.NotEmpty(t, lineValidation)

	assert.Equal(t, lineValidation.Code, code)
	return code, lineValidation.IDLineOwnerValidation
}

func makeUpdateLineOwnerValidation(t *testing.T) LineOwnerValidation {
	_, idLineOwnerValidation := makeLineOwnerValidation(t)
	userID := 1
	arg := UpdateUserIDtoLineOwnerValidationParams{
		IDUser: sql.NullInt32{
			Int32: int32(userID),
			Valid: true,
		},
		IDLineOwnerValidation: int32(idLineOwnerValidation),
	}
	updateLineValidation, err := testQueries.UpdateUserIDtoLineOwnerValidation(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, updateLineValidation)

	assert.Equal(t, updateLineValidation.IDUser, sql.NullInt32{
		Int32: int32(userID),
		Valid: true,
	})
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
	fmt.Println(code)
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now())

	lineValidation, err := testQueries.SelectCodeUnConfirmWithIn3Houses(context.Background(), code)
	assert.NoError(t, err)
	assert.NotEmpty(t, lineValidation)
	assert.Equal(t, code, lineValidation.Code)
}
