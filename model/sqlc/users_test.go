package model

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateUserOwnerValidation(t *testing.T) {
	user, err := testQueries.SelectUserID(context.Background(), int32(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, user.IDUser)

	//sent
	userAfterUpdata, errAfterUpdate := testQueries.UpdateUserOwnerValidation(context.Background(), user.IDUser)
	assert.NoError(t, errAfterUpdate)
	assert.NotEmpty(t, userAfterUpdata)

	assert.Equal(t, userAfterUpdata.IDUser, user.IDUser)
	assert.Equal(t, userAfterUpdata.OwnerValidation, true)
}

func TestSelectUserByLineUserID(t *testing.T) {
	lineID := sql.NullString{String: "Ue5308cc32ee5ca607c596e87877715b6", Valid: true}
	user, err := testQueries.SelectUserByLineUserID(context.Background(), lineID)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.IDUser)

	//sent
	assert.Equal(t, user.IDUser, int32(1))
	assert.Equal(t, lineID, user.IDLine)
}
