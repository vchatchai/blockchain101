package chaincode01

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type TradeWorkflowChaincode struct {
	testMode bool
}

func (t *TradeWorkflowChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("TradeWorkflow Init")

	return shim.Success(nil)
}

func (t *TradeWorkflowChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("TradeWorkflow Invoke")
	return shim.Success(nil)
}
