package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectFollowWalletByIDUser(t *testing.T) {
	user, err := testQueries.SelectUserID(context.Background(), int32(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, user.IDUser)

	wallet, err := testQueries.SelectFollowWalletByIDUser(context.Background(), user.IDUser)
	assert.NoError(t, err)
	assert.NotEmpty(t, wallet)

	assert.Equal(t, wallet.IDUser, user.IDUser)
	assert.NotEmpty(t, wallet.IDChain)
	assert.NotEmpty(t, wallet.CreatedAt)
}

func TestSelectWalletByMetamaskWalletID(t *testing.T) {
	wallet, err := testQueries.SelectWalletByMetamaskWalletID(context.Background(), "0x891B68D6B21c64d56dB262D066B38Ea76B6468f6")
	assert.NoError(t, err)
	assert.NotEmpty(t, wallet.MetamaskWalletID)
	assert.NotEmpty(t, wallet.IDChain)
	assert.NotEmpty(t, wallet.CreatedAt)
}

func TestUpdateLastBlockNumber(t *testing.T) {
	wallet, err := testQueries.SelectWalletByID(context.Background(), int32(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, wallet.WalletID)

	arg := UpdateLastBlockNumberParams{
		LastBlockNumber: wallet.LastBlockNumber + 1,
		WalletID:        wallet.WalletID,
	}

	walletAfterUpdate, errAfterUpdate := testQueries.UpdateLastBlockNumber(context.Background(), arg)
	assert.NoError(t, errAfterUpdate)
	assert.NotEmpty(t, walletAfterUpdate)
	assert.Equal(t, walletAfterUpdate.LastBlockNumber, wallet.LastBlockNumber+1)

	arg2 := UpdateLastBlockNumberParams{
		LastBlockNumber: walletAfterUpdate.LastBlockNumber - 1,
		WalletID:        wallet.WalletID,
	}

	walletAfterUpdate2, errAfterUpdate2 := testQueries.UpdateLastBlockNumber(context.Background(), arg2)
	assert.NoError(t, errAfterUpdate2)
	assert.NotEmpty(t, walletAfterUpdate2)
	assert.Equal(t, walletAfterUpdate2.LastBlockNumber, wallet.LastBlockNumber)
}
