package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

// *SmartContract provides functions for managing an Asset
// type *SmartContract struct {
// 	contractapi.Contract
// }

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetAllRecevingKdes(ctx contractapi.TransactionContextInterface, types string, userId string, status string, accept string) ([]map[string]interface{}, error) {

	var data []map[string]interface{}
	// if types == "ShippingKdesEvent" {
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"ReceiverInformation.ReceiverId\":\"%s\",\"Status\":\"%s\",\"IsAccepted\":\"%s\"}}", "ShippingKdes", userId, status, accept)
	// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"ReceiverInformation.ReceiverId\":\"%s\",\"Status\":\"%s\"}}", "ShippingKdes","3", "Transfered")
	// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var assets []map[string]interface{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset map[string]interface{}
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		data = append(data, asset)
	}

	// data["Product"] = assets
	// }
	return data, nil
	// return &response, nil
}



// Update a asset is accepted in the world state with given id.
func (s *SmartContract) UpdateRecevingKdesAccept(ctx contractapi.TransactionContextInterface, id string, accept string) (bool, error) {
	// Use the logger to print a message
	// var logger *logrus.Logger
	// logger.Info("Init function called")
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the farm %s does not exist", id)
	}
	var jsonData ShippingingKdes

	// Unmarshal the byte array into the empty interface
	err = json.Unmarshal(assetJSON, &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return false, fmt.Errorf("failed to unmarshal data: %s", err.Error())
	}
	jsonData.IsAccepted = accept
	// data["Product"] = assets
	// }
	assetJSON2, err4 := json.Marshal(jsonData)
	if err4 != nil {
		return false, fmt.Errorf("the asset json %s already exists", id)
	}
	ctx.GetStub().PutState(id, assetJSON2)

	return true, nil
	// return &response, nil
}

// setAssetStateBasedEndorsement adds an endorsement policy to an asset so that the passed orgs need to agree upon transfer
func setAssetStateBasedEndorsement2(ctx contractapi.TransactionContextInterface, assetID string, orgsToEndorse []string) error {
	endorsementPolicy, err := statebased.NewStateEP(nil)
	if err != nil {
		return fmt.Errorf("failed to NewStateEP to endorsement policy: %v", err)
	}
	err = endorsementPolicy.AddOrgs(statebased.RoleTypeMember, orgsToEndorse...)
	if err != nil {
		return fmt.Errorf("failed to add org to endorsement policy: %v", err)
	}
	policy, err := endorsementPolicy.Policy()
	if err != nil {
		return fmt.Errorf("failed to create endorsement policy bytes from org: %v", err)
	}
	err = ctx.GetStub().SetStateValidationParameter(assetID, policy)
	if err != nil {
		return fmt.Errorf("failed to set validation parameter on asset: %v", err)
	}

	return nil
}
