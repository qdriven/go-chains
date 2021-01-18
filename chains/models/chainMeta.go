package models

import "go-chains/global"

type ChainMetaData struct {
	global.GVA_MODEL
	ChainName string
	ChainId   string
	ChainType string  //mainnet/testnet
	RPCUrl string
}

type EvmTransactionRecord struct {
	global.GVA_MODEL
	FromChainId string
	ToChainId   string
	AssetHash string  //mainnet/testnet
	Amount string
	ToAddress string
}

