package model

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectChainByChainCode(t *testing.T) {
	chain, err := testQueries.SelectChainByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, chain)

	assert.Equal(t, chain.IDChain, int32(1))
	assert.Equal(t, chain.ChainCode, "bsc-testnet")
	assert.NotEmpty(t, chain.UrlApi)

	//default chain data table
	//INSERT INTO chain (id_chain, chain_code, chain_name, url_api )
	//VALUES (1, 'bsc-testnet', 'binance smart chain test net', 'https://api-testnet.bscscan.com/api');
}
