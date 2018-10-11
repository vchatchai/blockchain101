package chaincode01

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func getTradeKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	tradeKey, err := stub.CreateCompositeKey("Trade", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return tradeKey, nil
	}
}

func getLCKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	lcKey, err := stub.CreateCompositeKey("LetterOfCredit", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return lcKey, nil
	}
}

func getELKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	elKey, err := stub.CreateCompositeKey("ExportLicense", []string{tradeID})
	if err != nil {
		return "", err
	} else {
		return elKey, nil
	}
}

func getShipmentLocationKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	shipmentLocationKey, err := stub.CreateCompositeKey("Shipment", []string{"Location", tradeID})
	if err != nil {
		return "", err
	} else {
		return shipmentLocationKey, err
	}
}

func getBLKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	blKey, err := stub.CreateCompositeKey("BillOfLading", []string{tradeID})
	if err != nil {
		return "", err
	}
	return blKey, err
}

func getPaymentKey(stub shim.ChaincodeStubInterface, tradeID string) (string, error) {
	paymentKey, err := stub.CreateCompositeKey("Payment", []string{tradeID})
	if err != nil {
		return "", err
	}
	return paymentKey, err
}
