package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "deposit"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	DepositKeyPrefix = []byte{0x00}
)

func DepositKey(address sdk.AccAddress) []byte {
	return append(DepositKeyPrefix, address.Bytes()...)
}
