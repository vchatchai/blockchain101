package chaincode01

import (
	"crypto/x509"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, error) {
	var mspid string
	var err error
	var cert *x509.Certificate

	mspid, err = cid.GetMSPID(stub)
	if err != nil {
		fmt.Printf("Error getting MSP indentity: %s\n", err.Error())
		return "", "", err
	}

	return mspid, cert.Issuer.CommonName, nil
}

//For now, just hardcode an ACL
//We will suppot attribute checks in an upgrade

func authenticateExportingEntityOrg(mspID string, certCN string) bool {
	return (mspID == "ExportingEntityOrgMSP") && (certCN == "ca.exportinggentityorg.trade.com")
}

func authenticateExporterOrg(mspID string, certCN string) bool {
	return (mspID == "ExporterOrgMSP") && (certCN == "ca.exporterorg.trade.com")
}

func authenticateImporterOrg(mspID string, certCN string) bool {
	return (mspID == "ImporterOrgMSP") && (certCN == "ca.importerorg.trade.com")
}

func authenticateCarrierOrg(mspID string, certCN string) bool {
	return (mspID == "CarrierOrgMSP") && (certCN == "ca.carrierorg.trade.com")
}

func authenticateRegulatorOrg(mspID string, certCN string) bool {
	return (mspID == "RegulatorOrgMSP") && (certCN == "ca.regulatororg.trade.com")
}
