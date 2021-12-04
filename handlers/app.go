package handlers

import (
	"errors"
	"fiber_api_eth/contract"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type App struct {
	client           *ethclient.Client
	contractInstance *contract.Contract
}

func newEthClient(netUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(netUrl)
	if err != nil {
		return nil, errors.New("unable to create eth client")
	}
	return client, nil
}

func NewApp(netUrl string, ContractAddress string) (*App, error) {
	var err error
	a := App{}

	a.client, err = newEthClient(netUrl)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress("0x4f7f1380239450AAD5af611DB3c3c1bb51049c29")
	log.Println("Address: ", address.Hex())

	a.contractInstance, err = contract.NewContract(address, a.client)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
