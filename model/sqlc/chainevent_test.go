package model

import (
	"context"
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/tool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertChainEvent(t *testing.T) {
	Wallet, err := testQueries.SelectWalletByID(context.Background(), int32(1))
	assert.NoError(t, err)

	testWallet := "test" + tool.RandomWallet()
	value := "2340000000000000"

	//sent
	chainEvent := addEvent(t, value, Wallet.MetamaskWalletID, testWallet)
	assert.Equal(t, chainEvent.ActionType, sql.NullString{String: "tranfer", Valid: true})
	assert.Equal(t, chainEvent.FromMetamaskWalletID, Wallet.MetamaskWalletID)
	assert.Equal(t, chainEvent.ToMetamaskWalletID, testWallet)
	assert.NotEmpty(t, chainEvent.LogEvent)

	//received
	chainEventReceived := addEvent(t, value, testWallet, Wallet.MetamaskWalletID)
	assert.Equal(t, chainEventReceived.ActionType, sql.NullString{String: "tranfer", Valid: true})
	assert.Equal(t, chainEventReceived.FromMetamaskWalletID, testWallet)
	assert.Equal(t, chainEventReceived.ToMetamaskWalletID, Wallet.MetamaskWalletID)
	assert.NotEmpty(t, chainEventReceived.LogEvent)
}

func addEvent(t *testing.T, value string, FromMetamaskWalletID string, ToMetamaskWalletID string) ChainEvent {
	arg := InsertChainEventParams{
		WalletID:             int32(1),
		ActionType:           sql.NullString{String: "tranfer", Valid: true},
		Value:                sql.NullString{String: value, Valid: true},
		FromMetamaskWalletID: FromMetamaskWalletID,
		ToMetamaskWalletID:   ToMetamaskWalletID,
		LogEvent:             sql.NullString{String: tool.GetMockOneBlockData(), Valid: true},
	}

	chainEvent, err := testQueries.InsertChainEvent(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, chainEvent)
	return chainEvent
}
