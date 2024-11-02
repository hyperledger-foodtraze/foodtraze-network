package chaincode

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type BlockchainInfo struct {
	TransactionID string    `json:"TransactionID"`
	ClientId      string    `json:"ClientId"`
	BlockNumber   int64     `json:"BlockNumber"`
	ChannelId     string    `json:"ChannelId"`
	Timestamp     time.Time `json:"Timestamp"`
	MspId         string    `json:"MspId"`
}

type Header struct {
	EventWhy      string `json:"eventWhy"`
	EventWhen     string `json:"eventWhen"`
	EventWhere    string `json:"eventWhere"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	UnixTimeStamp string `json:"UnixTimeStamp"`
	DeviceInfo    string `json:"deviceInfo"`
}

type TraceEventRes struct {
	EventID string `json:"eventId"`
	TraceId string `json:"traceId"`
	Data    interface{}
}

type TraceEventErrorRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    map[string]interface{} `json:"record"`
	TxId      string                 `json:"txId"`
	Timestamp time.Time              `json:"timestamp"`
	IsDelete  bool                   `json:"isDelete"`
}

type FoodTazeRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

// InitLedger adds a base set of assets to the ledger
// func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
// 	asset2 := Asset2{
// 		Quantity:  9,
// 		FruitType: "deciduous",
// 		FruitName: "Apple",
// 	}

// 	assets := []Asset{
// 		{ID: "asset1", Color: "bluee", Size: 5, Owner: "Tomoko", AppraisedValue: 300, Data: asset2},
// 	}

// 	for _, asset := range assets {
// 		assetJSON, err := json.Marshal(asset)
// 		if err != nil {
// 			// return ,err
// 		}

// 		err = ctx.GetStub().PutState(asset.ID, assetJSON)
// 		if err != nil {
// 			// return fmt.Errorf("failed to put to world state. %v", err)
// 		}
// 		// Changes the endorsement policy to the new owner org
// 		endorsingOrgs := []string{"Org1MSP"}
// 		err = setAssetStateBasedEndorsement(ctx, asset.ID, endorsingOrgs)
// 		if err != nil {
// 			return fmt.Errorf("failed setting state based endorsement for new owner: %v", err)
// 		}
// 	}

//		return nil
//	}

// data
func (s *SmartContract) EndorseChange(ctx contractapi.TransactionContextInterface, data string) (interface{}, error) {
	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org4MSP"}
	err1 := setAssetStateBasedEndorsement(ctx, data, endorsingOrgs)
	if err1 != nil {
		return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	}

	response := FoodTazeRes{
		Status:  200,
		Message: "Crop Event Created Successfully.",
		Data:    "",
	}
	return response, nil
}

func (s *SmartContract) FoodTrazeCreateNew(ctx contractapi.TransactionContextInterface, status string, data1 string, data2 string, data3 string) (interface{}, error) {
	data := make(map[string]interface{})

	var event map[string]interface{}
	if err1 := json.Unmarshal([]byte(data2), &event); err1 != nil {
		// fmt.Println("Error parsing JSON1:", err1)
		return nil, fmt.Errorf("the ipfs image data error %v", err1)
	}
	data["Crop"] = event
	data["FTLCID"] = data1
	var event2 []map[string]interface{}
	if err1 := json.Unmarshal([]byte(data3), &event2); err1 != nil {
		// fmt.Println("Error parsing JSON1:", err1)
		return nil, fmt.Errorf("the ipfs image data2 error %v", err1)
	}
	data["ArrData"] = event2
	var headerContent Header
	if err1 := json.Unmarshal([]byte(status), &headerContent); err1 != nil {
		// fmt.Println("Error parsing JSON1:", err1)
		return nil, fmt.Errorf("the participant error %v", err1)
	}
	data["Header"] = headerContent
	assetJSON, err4 := json.Marshal(data)
	if err4 != nil {
		return nil, fmt.Errorf("the asset json %s already exists", data["FTLCID"])
	}

	result := ctx.GetStub().PutState(data1, assetJSON)
	response := FoodTazeRes{
		Status:  200,
		Message: "Asset Created Successfully.",
		Data:    result,
	}
	return response, nil
}

func (s *SmartContract) CreateTraze(ctx contractapi.TransactionContextInterface, data string) (interface{}, error) {
	var event map[string]interface{}
	if err1 := json.Unmarshal([]byte(data), &event); err1 != nil {
		// fmt.Println("Error parsing JSON1:", err1)
		return nil, fmt.Errorf("the ipfs image data error %v", err1)
	}

	channelId := ctx.GetStub().GetChannelID()
	transactionId := ctx.GetStub().GetTxID()
	// timestamps, _ := ctx.GetStub().GetTxTimestamp()
	timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
	// Retrieve the block number from the transaction timestamp
	// blockNumber := timestamps.GetSeconds() / 10
	clientId, _ := ctx.GetClientIdentity().GetID()
	var blockChainInfo BlockchainInfo
	// blockChainInfo.BlockNumber = clientId
	blockChainInfo.TransactionID = transactionId
	blockChainInfo.ClientId = clientId
	blockChainInfo.ChannelId = channelId
	blockChainInfo.Timestamp = timestamp
	event["BlockchainInfo"] = &blockChainInfo
	assetJSON, err4 := json.Marshal(event)
	if err4 != nil {
		return nil, fmt.Errorf("the asset json %s already exists", event["FTLCID"].(string))
	}

	result := ctx.GetStub().PutState(event["FTLCID"].(string), assetJSON)
	return result, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadTrazeById(ctx contractapi.TransactionContextInterface, id string) (map[string]interface{}, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}
	var asset map[string]interface{}
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	return asset, nil

}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetAllTraze(ctx contractapi.TransactionContextInterface, filter string) ([]map[string]interface{}, error) {

	var data []map[string]interface{}
	// var asset CropDetails
	// err = json.Unmarshal(assetJSON, &asset)
	// if err != nil {
	// 	return nil, fmt.Errorf("the unmarshall error %s", err)
	// }

	// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "Crop", userId)
	// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
	resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
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
	return data, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) DeleteTrazeById(ctx contractapi.TransactionContextInterface, id string) error {
	// exists, err2 := s.AssetExists(ctx, id)
	// if err2 != nil {
	// 	return fmt.Errorf("the traze data %s exist error", err2)
	// }
	// if !exists {
	// 	return fmt.Errorf("the traze data %s not exists", id)
	// }
	// Retrieve the state to check if it exists
	data, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to get state for id %s: %v", id, err)
	}
	if data == nil {
		return fmt.Errorf("state with id %s does not exist", id)
	}

	// return ctx.GetStub().DelState(id)
	err1 := ctx.GetStub().DelState(id)
	if err1 != nil {
		return fmt.Errorf("failed to delete from state: %v", err1)
	}
	return nil
}

// GetAssetHistory returns the chain of custody for an asset since issuance.
func (s *SmartContract) GetAllTrazeHistoryById(ctx contractapi.TransactionContextInterface, assetID string) ([]HistoryQueryResult, error) {
	// log.Printf("GetAssetHistory: ID %v", assetID)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var records []map[string]interface{}
	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset map[string]interface{}
		// if len(response.Value) > 0 {
		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, err
		}
		// } else {
		// 	asset = TraceEvent{
		// 		BatchNumber: assetID,
		// 	}
		// }
		timestamp, err := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}

func (s *SmartContract) ChangeEndorsePolicy(ctx contractapi.TransactionContextInterface, data1 string, data2 string) (bool, error) {

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{data2}
	err1 := setAssetStateBasedEndorsement(ctx, data1, endorsingOrgs)
	if err1 != nil {
		return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	}
	return true, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) FoodTrazeability(ctx contractapi.TransactionContextInterface, filter string) ([]map[string]interface{}, error) {

	var data []map[string]interface{}
	// if status == "CropEvent" {
	// var asset CropDetails
	// err = json.Unmarshal(assetJSON, &asset)
	// if err != nil {
	// 	return nil, fmt.Errorf("the unmarshall error %s", err)
	// }

	// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "Crop", userId)
	// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
	resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
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

	// }
	return data, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetFarmByPagination(ctx contractapi.TransactionContextInterface, limit, offset string) ([]map[string]interface{}, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\"}}", "Farm")
	var Limit int
	if limit != "" {
		i, err := strconv.ParseInt(limit, 10, 32)
		if err != nil {
			panic(err)
		}
		Limit = int(i)
	} else {
		Limit = 0
	}
	var Offset int
	if offset != "" {
		i, err := strconv.ParseInt(offset, 10, 32)
		if err != nil {
			panic(err)
		}
		Offset = int(i)
	} else {
		Offset = 0
	}

	resultsIterator, _, err := ctx.GetStub().GetQueryResultWithPagination(queryString, int32(Limit), "")
	if err != nil {
		return nil, err
	}
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"Farmer.ContactInformation.Email\":\"%s\"}}", "Farm", email)
	// resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	// if err != nil {
	// 	return nil, err
	// }
	defer resultsIterator.Close()

	var farms []map[string]interface{}
	// Apply offset
	for i := 0; i < Offset; i++ {
		if resultsIterator.HasNext() {
			_, err := resultsIterator.Next()
			if err != nil {
				return nil, fmt.Errorf("failed to iterate over results: %v", err)
			}
		}
	}
	for i := 0; i < Limit && resultsIterator.HasNext(); i++ {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate farm: %v", err)
		}
		// fmt.log("queryResponse.Value", queryResponse.Value)
		var farm map[string]interface{}
		err = json.Unmarshal(queryResponse.Value, &farm)
		if err != nil {
			return nil, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farms = append(farms, farm)
	}
	return farms, nil
}

// GetAllFarms returns all assets found in world state
// func (s *SmartContract) CheckFarmEmail(ctx contractapi.TransactionContextInterface, email string) (bool, error) {

// 	// range query with empty string for startKey and endKey does an
// 	// open-ended query of all assets in the chaincode namespace.
// 	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"Farmer.ContactInformation.Email\":\"%s\"}}", "Farm", email)
// 	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
// 	if err != nil {
// 		return true, err
// 	}
// 	defer resultsIterator.Close()

// 	var farms []*Farm
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return true, fmt.Errorf("failed to iterate farm: %v", err)
// 		}
// 		// fmt.log("queryResponse.Value", queryResponse.Value)
// 		var farm Farm
// 		err = json.Unmarshal(queryResponse.Value, &farm)
// 		if err != nil {
// 			return true, fmt.Errorf("unmarshall farm data: %v", err)
// 		}
// 		farms = append(farms, &farm)
// 	}
// 	var result bool
// 	if len(farms) != 0 {
// 		result = true
// 	} else {
// 		result = false
// 	}
// 	return result, nil
// }

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// setAssetStateBasedEndorsement adds an endorsement policy to an asset so that the passed orgs need to agree upon transfer
func setAssetStateBasedEndorsement(ctx contractapi.TransactionContextInterface, assetID string, orgsToEndorse []string) error {
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
	// Log the change
	ctx.GetStub().SetEvent("EndorsementPolicyChanged", policy)
	return nil
}

func generateUniqueAssetID() string {
	// Implement your logic for generating a unique asset ID
	// Example: timestamp + random number
	return strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(1000))
}
