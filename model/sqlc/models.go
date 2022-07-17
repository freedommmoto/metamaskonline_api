// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package model

import (
	"database/sql"
	"time"
)

type Chain struct {
	IDChain   int32          `json:"id_chain"`
	ChainCode string         `json:"chain_code"`
	ChainName string         `json:"chain_name"`
	UrlApi    sql.NullString `json:"url_api"`
	CreatedAt time.Time      `json:"created_at"`
	Modified  time.Time      `json:"modified"`
	Deleted   sql.NullTime   `json:"deleted"`
}

type ChainEvent struct {
	IDChainEvent         int64          `json:"id_chain_event"`
	WalletID             int32          `json:"wallet_id"`
	ActionType           sql.NullString `json:"action_type"`
	FromMetamaskWalletID string         `json:"from_metamask_wallet_id"`
	ToMetamaskWalletID   string         `json:"to_metamask_wallet_id"`
	Value                sql.NullString `json:"value"`
	LogEvent             sql.NullString `json:"log_event"`
	CreatedAt            time.Time      `json:"created_at"`
}

type Event struct {
	IDEvent      int64     `json:"id_event"`
	IDLineEvent  int32     `json:"id_line_event"`
	IDChainEvent int32     `json:"id_chain_event"`
	CreatedAt    time.Time `json:"created_at"`
}

type LineEvent struct {
	IDLineEvent      int64          `json:"id_line_event"`
	IDLineUser       string         `json:"id_line_user"`
	IDUse            int32          `json:"id_use"`
	RequestLogEvent  sql.NullString `json:"request_log_event"`
	ResponseLogEvent sql.NullString `json:"response_log_event"`
	Error            bool           `json:"error"`
	ErrorText        sql.NullString `json:"error_text"`
	CreatedAt        time.Time      `json:"created_at"`
}

type LineOwnerValidation struct {
	IDLineOwnerValidation int32         `json:"id_line_owner_validation"`
	Code                  string        `json:"code"`
	IDUser                sql.NullInt32 `json:"id_user"`
	CreatedAt             time.Time     `json:"created_at"`
}

type User struct {
	IDUser          int32          `json:"id_user"`
	Username        string         `json:"username"`
	Password        string         `json:"password"`
	IDLine          sql.NullString `json:"id_line"`
	OwnerValidation bool           `json:"owner_validation"`
	CreatedAt       time.Time      `json:"created_at"`
	Modified        time.Time      `json:"modified"`
	Deleted         sql.NullTime   `json:"deleted"`
}

type Wallet struct {
	WalletID         int32          `json:"wallet_id"`
	MetamaskWalletID string         `json:"metamask_wallet_id"`
	FollowWallet     bool           `json:"follow_wallet"`
	IDUser           int32          `json:"id_user"`
	IDChain          sql.NullInt32  `json:"id_chain"`
	LastBlockNumber  int32          `json:"last_block_number"`
	CreatedAt        time.Time      `json:"created_at"`
	Modified         time.Time      `json:"modified"`
	Deleted          sql.NullTime   `json:"deleted"`
	WalletName       sql.NullString `json:"wallet_name"`
}
