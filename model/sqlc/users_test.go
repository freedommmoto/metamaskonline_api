package model

import (
	"context"
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
