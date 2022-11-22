package evm

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type evm struct {
	AddressContractMap map[string]*Contract
	AddressBalanceMap  map[string][]byte
}

func (e *evm) AddContract(address string, contract *Contract) {
	if e.AddressContractMap == nil {
		e.AddressContractMap = make(map[string]*Contract)
	}
	e.AddressContractMap[address] = contract
}

var evminstance evm

func GetEVM() *evm {
	return &evminstance
}

type Contract struct {
	Address        string
	Bytecode       []byte
	ProgramCounter int16    // evm smart contract has max. size of 24kb => 24.576b
	Stack          [][]byte // reverse stack => len(Stack)-1 is top
	Memory         []byte
	Storage        [][]byte
}

func NewContract(bytecode []byte) Contract {
	privk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	pubKey := privk.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		panic("public key is not type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return Contract{
		Address:  strings.ToLower(address),
		Bytecode: bytecode,
	}
}

type Tx struct {
	Origin string // origin sender address
	Value  []byte
	Data   []byte
}
