package model

import (
	"context"
	"github.com/freedommmoto/metamaskonline_api/tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertLineOwnerValidation(t *testing.T) {
	user, err := testQueries.SelectUserID(context.Background(), int32(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, user.IDLine.String)

	code := tool.RandomCodeNumber(4)
	arg := InsertLineOwnerValidationParams{
		IDUser: user.IDUser,
		Code:   code,
	}

	lineValidation, err := testQueries.InsertLineOwnerValidation(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, lineValidation)

	assert.Equal(t, lineValidation.IDUser, user.IDUser)
	assert.Equal(t, lineValidation.Code, code)
}
