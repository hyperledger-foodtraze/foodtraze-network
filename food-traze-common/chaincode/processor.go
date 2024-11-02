package chaincode

import (
	"encoding/json"
	"fmt"

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
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"ReceiverInformation.ReceiverId\":\"%s\",\"Status\":\"%s\",\"IsAccepted\":\"%s\"}}", types, userId, status, accept)
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
