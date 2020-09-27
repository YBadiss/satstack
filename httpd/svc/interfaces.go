package svc

import (
	"github.com/onyb/sat-stack/bus"
	"github.com/onyb/sat-stack/types"
)

type TransactionsService interface {
	GetTransaction(hash string, block *types.Block) (*types.Transaction, error)
	GetTransactionHex(hash string) (string, error)
	SendTransaction(tx string) (*string, error)
}

type BlocksService interface {
	GetBlock(ref string) (*types.Block, error)
}

type AddressesService interface {
	GetAddresses(addresses []string, blockHash *string) (types.Addresses, error)
}

type ExplorerService interface {
	GetHealth() error
	GetStatus() (*bus.ExplorerStatus, error)
	GetFees(targets []int64, mode string) map[string]interface{}
}

type ServiceInterface interface {
	BlocksService
	TransactionsService
	AddressesService
	ExplorerService
}
