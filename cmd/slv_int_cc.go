package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SLV Integer smart contract
type SlvIntSmartContract struct {
	contractapi.Contract
}

// SLV Integer type
type SlvInt struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewSlvInt creates a new serverless ledger variable of integer type and sets the value
func (s *SlvIntSmartContract) NewSlvInt(ctx contractapi.TransactionContextInterface, name, value string) error {
	existing, err := ctx.GetStub().GetState(name)
	if err != nil {
		return fmt.Errorf("unable to read from ledger - [%s]", err.Error())
	}

	if existing != nil {
		return fmt.Errorf("ledger variable [%s] already exists", name)
	}

	slv := SlvInt{
		Name:  name,
		Value: value,
	}
	slvBytes, _ := json.Marshal(slv)

	return ctx.GetStub().PutState(name, slvBytes)
}

// GetSlvInt finds the variable by name in the ledger and returns the SlvInt structure.
func (s *SlvIntSmartContract) GetSlvInt(ctx contractapi.TransactionContextInterface, name string) (*SlvInt, error) {
	slvBytes, err := ctx.GetStub().GetState(name)
	if err != nil {
		return nil, fmt.Errorf("unable to read from ledger - [%s]", err.Error())
	}

	if slvBytes == nil {
		return nil, fmt.Errorf("ledger variable [%s] does not exist", name)
	}

	slv := new(SlvInt)
	_ = json.Unmarshal(slvBytes, slv)

	return slv, nil
}

// GetSlvIntValue finds the variable by name and returns the value
func (s *SlvIntSmartContract) GetSlvIntValue(ctx contractapi.TransactionContextInterface, name string) (string, error) {
	slvBytes, err := ctx.GetStub().GetState(name)
	if err != nil {
		return "", fmt.Errorf("unable to read from ledger - [%s]", err.Error())
	}

	if slvBytes == nil {
		return "", fmt.Errorf("ledger variable [%s] does not exist", name)
	}

	slv := new(SlvInt)
	_ = json.Unmarshal(slvBytes, slv)

	return slv.Value, nil
}

// SetSlvIntValue sets the value of the variable in the ledger.
func (s *SlvIntSmartContract) SetSlvIntValue(ctx contractapi.TransactionContextInterface, name string, value string) error {
	existing, err := ctx.GetStub().GetState(name)
	if err != nil {
		return fmt.Errorf("unable to read from ledger - [%s]", err.Error())
	}

	if existing == nil {
		return fmt.Errorf("ledger variable [%s] does not exist", name)
	}

	elv := new(SlvInt)
	_ = json.Unmarshal(existing, elv)

	slv := SlvInt{
		Name:  name,
		Value: value,
	}

	slvBytes, _ := json.Marshal(slv)

	return ctx.GetStub().PutState(name, slvBytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SlvIntSmartContract))

	if err != nil {
		fmt.Printf("Error in creating SlvInt chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting SlvInt chaincode: %s", err.Error())
	}
}
