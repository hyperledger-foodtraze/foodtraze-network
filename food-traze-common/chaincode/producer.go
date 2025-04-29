package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
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
	resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
	if err != nil {
		return nil, fmt.Errorf("unmarshall farm data1: %v", err)
	}
	defer resultsIterator.Close()

	// var assets []map[string]interface{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("unmarshall farm data2: %v", err)
		}

		var asset map[string]interface{}
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("unmarshall farm data3: %v", err)
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

// GetAllHarvest returns all assets found in world state
func (s *SmartContract) GetAllProductEventsById(ctx contractapi.TransactionContextInterface, cropId string, filters string) (map[string]interface{}, error) {
	exists, err2 := s.AssetExists(ctx, cropId)
	if err2 != nil {
		return nil, fmt.Errorf("the product data %s exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the product data %s not exists", cropId)
	}

	assets := make(map[string]interface{})

	assetJSON, err := ctx.GetStub().GetState(cropId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the crop %s does not exist", cropId)
	}
	// var response FoodTazeRes
	// var data TrazeDetail
	var asset map[string]interface{}
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Product"] = asset

	assetJSON1, err := ctx.GetStub().GetState(asset["ParentId"].(string))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON1 == nil {
		return nil, fmt.Errorf("the farm %s does not exist", asset["ParentId"].(string))
	}
	// var response FoodTazeRes
	// var data TrazeDetail
	var asset1 map[string]interface{}
	err = json.Unmarshal(assetJSON1, &asset1)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Owner"] = asset1

	var data []map[string]interface{}
	// Started To check Type as Fertilization
	// filter := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Fertilization")
	resultsIterator, err := ctx.GetStub().GetQueryResult(filters)
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
	assets["Events"] = data
	return assets, nil
}

// GetAllProductEventsByBatchId returns all assets found in world state
func (s *SmartContract) GetAllProductEventsByBatchIdIngredient(ctx contractapi.TransactionContextInterface, cropId string, label string) (map[string]interface{}, error) {
	assets := make(map[string]interface{})
	if label == "Batch" {
		// --------------------   Get ProductData   -----------------------
		filters := fmt.Sprintf("{\"selector\":{\"Data.BatchId\":\"%s\",\"DocType\":\"%s\"}}", cropId, "TransformationProduct")
		resultsIterator0, err := ctx.GetStub().GetQueryResult(filters)
		if err != nil {
			return nil, fmt.Errorf("GetQueryResult: %v", err)
		}
		defer resultsIterator0.Close()
		var asset0 map[string]interface{}
		// var assets []map[string]interface{}
		for resultsIterator0.HasNext() {
			queryResponse0, err := resultsIterator0.Next()
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(queryResponse0.Value, &asset0)
			if err != nil {
				return nil, fmt.Errorf("the unmarshall product error %s", err)
			}
			// data = append(data, asset)
		}
		assets["Product"] = asset0

		// --------------------   Get Product Trace Ledger Data   -----------------------
		assetJSON, err := ctx.GetStub().GetState(asset0["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON == nil {
			return nil, fmt.Errorf("the crop %s does not exist", cropId)
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		// Assign value for crop
		assets["Ledger"] = asset
		// --------------------   Get Farm Data   -----------------------
		assetJSON1, err := ctx.GetStub().GetState(asset["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state1: %v", err)
		}
		if assetJSON1 == nil {
			return nil, fmt.Errorf("the farm %s does not exist1", asset["ParentId"].(string))
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset1 map[string]interface{}
		err = json.Unmarshal(assetJSON1, &asset1)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error1 %s", err)
		}
		// Assign value for crop
		assets["Farm"] = asset1
		// --------------------   Get Events Data   -----------------------
		var data []map[string]interface{}
		// Started To check Type as Fertilization
		filter := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"}}", asset["FTLCID"], "Event")
		resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
		if err != nil {
			return nil, fmt.Errorf("error event GetQueryResult: %v", err)
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
				return nil, fmt.Errorf("the unmarshall event error %s", err)
			}
			data = append(data, asset)
		}
		assets["Events"] = data
		// --------------------   Get Ingredient Data   -----------------------
		var ingredient []map[string]interface{}
		fmt.Println("Inside Length")
		log.Printf("Inside Length")
		// var ids []string
		for _, ingred := range data {
			data1 := ingred["Data"].(map[string]interface{})
			batchIdsIface := data1["BatchId"]
			// if !ok {
			// 	log.Printf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	err0 := fmt.Errorf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	fmt.Errorf("the continue1 ingredient error %t", err0)
			// 	continue
			// }

			batchIds, ok := batchIdsIface.([]interface{}) // most likely type
			if !ok {
				return nil, fmt.Errorf("invalid BatchId type")
			}
			fmt.Println("batchId", batchIds)
			// Started To check Type as Fertilization
			for _, id := range batchIds {
				batchId := id.(string)
				filter1 := fmt.Sprintf("{\"selector\":{\"Data.BatchId\":\"%s\",\"DocType\":\"%s\"}}", batchId, "TransformationProduct")
				resultsIterator1, err := ctx.GetStub().GetQueryResult(filter1)
				if err != nil {
					return nil, fmt.Errorf("error ingredient GetQueryResult: %v", err)
				}
				defer resultsIterator1.Close()
				for resultsIterator1.HasNext() {
					queryResponse, err := resultsIterator1.Next()
					if err != nil {
						return nil, err
					}

					var asset1 map[string]interface{}
					err = json.Unmarshal(queryResponse.Value, &asset1)
					if err != nil {
						return nil, fmt.Errorf("the unmarshall ingredient error %s", err)
					}
					ingredient = append(ingredient, asset1)
					fmt.Println("Appended ingredient")
				}
			}
		}
		assets["Ingredient"] = ingredient

	} else if label == "FTLC" {
		// Check Exist
		exists, err2 := s.AssetExists(ctx, cropId)
		if err2 != nil {
			return nil, fmt.Errorf("the product data %s exist error", err2)
		}
		if !exists {
			return nil, fmt.Errorf("the product data %s not exists", cropId)
		}
		// --------------------   Get ProductData   -----------------------
		assetJSON0, err := ctx.GetStub().GetState(cropId)
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON0 == nil {
			return nil, fmt.Errorf("the crop %s does not exist", cropId)
		}
		var asset0 map[string]interface{}
		err = json.Unmarshal(assetJSON0, &asset0)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		assets["Product"] = asset0

		// --------------------   Get Product Trace Ledger Data   -----------------------
		assetJSON, err := ctx.GetStub().GetState(asset0["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON == nil {
			return nil, fmt.Errorf("the crop %s does not exist", cropId)
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		// Assign value for crop
		assets["Ledger"] = asset
		// --------------------   Get Farm Data   -----------------------
		assetJSON1, err := ctx.GetStub().GetState(asset["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON1 == nil {
			return nil, fmt.Errorf("the farm %s does not exist", asset["ParentId"].(string))
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset1 map[string]interface{}
		err = json.Unmarshal(assetJSON1, &asset1)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		// Assign value for crop
		assets["Farm"] = asset1
		// --------------------   Get Events Data   -----------------------
		var data []map[string]interface{}
		// Started To check Type as Fertilization
		filter := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"}}", asset["FTLCID"], "Event")
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
		assets["Events"] = data
		// --------------------   Get Ingredient Data   -----------------------
		var ingredient []map[string]interface{}
		fmt.Println("Inside Length")
		log.Printf("Inside Length")
		// var ids []string
		for _, ingred := range data {
			data1 := ingred["Data"].(map[string]interface{})
			batchIdsIface := data1["Batch Id"]
			// if !ok {
			// 	log.Printf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	err0 := fmt.Errorf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	fmt.Errorf("the continue1 ingredient error %t", err0)
			// 	continue
			// }

			batchIds, ok := batchIdsIface.([]interface{}) // most likely type
			if !ok {
				return nil, fmt.Errorf("invalid BatchId type")
			}
			fmt.Println("batchId", batchIds)
			// Started To check Type as Fertilization
			for _, id := range batchIds {
				batchId := id.(string)
				filter1 := fmt.Sprintf("{\"selector\":{\"Data.BatchId\":\"%s\",\"DocType\":\"%s\"}}", batchId, "TransformationProduct")
				resultsIterator1, err := ctx.GetStub().GetQueryResult(filter1)
				if err != nil {
					return nil, fmt.Errorf("error ingredient GetQueryResult: %v", err)
				}
				defer resultsIterator1.Close()
				for resultsIterator1.HasNext() {
					queryResponse, err := resultsIterator1.Next()
					if err != nil {
						return nil, err
					}

					var asset1 map[string]interface{}
					err = json.Unmarshal(queryResponse.Value, &asset1)
					if err != nil {
						return nil, fmt.Errorf("the unmarshall ingredient error %s", err)
					}
					ingredient = append(ingredient, asset1)
					fmt.Println("Appended ingredient")
				}
			}
		}
		assets["Ingredient"] = ingredient

	}

	return assets, nil
}

// GetAllProductEventsByBatchId returns all assets found in world state
func (s *SmartContract) GetAllProductEventsByBatchId(ctx contractapi.TransactionContextInterface, cropId string, label string) (map[string]interface{}, error) {
	assets := make(map[string]interface{})
	if label == "Batch" {
		// --------------------   Get ProductData   -----------------------
		filters := fmt.Sprintf("{\"selector\":{\"Data.BatchId\":\"%s\",\"DocType\":\"%s\"}}", cropId, "TransformationProduct")
		resultsIterator0, err := ctx.GetStub().GetQueryResult(filters)
		if err != nil {
			return nil, fmt.Errorf("GetQueryResult: %v", err)
		}
		defer resultsIterator0.Close()
		var asset0 map[string]interface{}
		// var assets []map[string]interface{}
		for resultsIterator0.HasNext() {
			queryResponse0, err := resultsIterator0.Next()
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(queryResponse0.Value, &asset0)
			if err != nil {
				return nil, fmt.Errorf("the unmarshall product error %s", err)
			}
			// data = append(data, asset)
		}
		assets["Product"] = asset0

		// --------------------   Get Product Trace Ledger Data   -----------------------
		assetJSON, err := ctx.GetStub().GetState(asset0["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON == nil {
			return nil, fmt.Errorf("the crop %s does not exist", cropId)
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		// Assign value for crop
		assets["Ledger"] = asset
		// --------------------   Get Farm Data   -----------------------
		assetJSON1, err := ctx.GetStub().GetState(asset["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state1: %v", err)
		}
		if assetJSON1 == nil {
			return nil, fmt.Errorf("the farm %s does not exist1", asset["ParentId"].(string))
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset1 map[string]interface{}
		err = json.Unmarshal(assetJSON1, &asset1)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error1 %s", err)
		}
		// Assign value for crop
		assets["Farm"] = asset1
		// --------------------   Get Events Data   -----------------------
		var data []map[string]interface{}
		// Started To check Type as Fertilization
		filter := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"},\"sort\": [{\"Data.Date\": \"desc\"}]}", asset["FTLCID"], "Event")
		resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
		if err != nil {
			return nil, fmt.Errorf("error event GetQueryResult: %v", err)
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
				return nil, fmt.Errorf("the unmarshall event error %s", err)
			}
			data = append(data, asset)
		}

		// --------------------   Get Ingredient Data   -----------------------
		var ingredient []map[string]interface{}
		fmt.Println("Inside Length")
		log.Printf("Inside Length")
		// var ids []string
		for _, ingred := range data {
			data1 := ingred["Data"].(map[string]interface{})
			batchIdsIface := data1["Batch Id"]
			// if !ok {
			// 	log.Printf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	err0 := fmt.Errorf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	fmt.Errorf("the continue1 ingredient error %t", err0)
			// 	continue
			// }

			batchIds, ok := batchIdsIface.([]interface{}) // most likely type
			if !ok {
				// return nil, fmt.Errorf("invalid BatchId type")
			}
			fmt.Println("batchId", batchIds)
			// Started To check Type as Fertilization
			for _, id := range batchIds {
				batchId := id.(string)
				filter1 := fmt.Sprintf("{\"selector\":{\"Data.BatchId\":\"%s\",\"DocType\":\"%s\"}}", batchId, "TransformationProduct")
				resultsIterator1, err := ctx.GetStub().GetQueryResult(filter1)
				if err != nil {
					return nil, fmt.Errorf("error ingredient GetQueryResult: %v", err)
				}
				defer resultsIterator1.Close()
				for resultsIterator1.HasNext() {
					queryResponse, err := resultsIterator1.Next()
					if err != nil {
						return nil, err
					}

					var asset1 map[string]interface{}
					err = json.Unmarshal(queryResponse.Value, &asset1)
					if err != nil {
						return nil, fmt.Errorf("the unmarshall ingredient error %s", err)
					}
					ingredient = append(ingredient, asset1)
					fmt.Println("Appended ingredient")
				}
			}
		}
		assets["Ingredient"] = ingredient

		// --------------------   Get Ingredient Events Data   -----------------------
		if len(ingredient) != 0 {
			for _, ingred := range ingredient {
				dataIng := ingred["Data"].(map[string]interface{})
				// Started To check Type as Fertilization
				filter := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"},\"sort\": [{\"Data.Date\": \"desc\"}]}", dataIng["ProductTraceLedgerId"], "Event")
				resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
				if err != nil {
					return nil, fmt.Errorf("error event GetQueryResult: %v", err)
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
						return nil, fmt.Errorf("the unmarshall event error %s", err)
					}
					data = append(data, asset)
				}
			}

			assets["Events"] = data
		}

	} else if label == "FTLC" {
		// Check Exist
		exists, err2 := s.AssetExists(ctx, cropId)
		if err2 != nil {
			return nil, fmt.Errorf("the product data %s exist error", err2)
		}
		if !exists {
			return nil, fmt.Errorf("the product data %s not exists", cropId)
		}
		// --------------------   Get ProductData   -----------------------
		assetJSON0, err := ctx.GetStub().GetState(cropId)
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON0 == nil {
			return nil, fmt.Errorf("the crop %s does not exist", cropId)
		}
		var asset0 map[string]interface{}
		err = json.Unmarshal(assetJSON0, &asset0)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		assets["Product"] = asset0

		// --------------------   Get Product Trace Ledger Data   -----------------------
		assetJSON, err := ctx.GetStub().GetState(asset0["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON == nil {
			return nil, fmt.Errorf("the crop %s does not exist", cropId)
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		// Assign value for crop
		assets["Ledger"] = asset
		// --------------------   Get Farm Data   -----------------------
		assetJSON1, err := ctx.GetStub().GetState(asset["ParentId"].(string))
		if err != nil {
			return nil, fmt.Errorf("failed to read from world state: %v", err)
		}
		if assetJSON1 == nil {
			return nil, fmt.Errorf("the farm %s does not exist", asset["ParentId"].(string))
		}
		// var response FoodTazeRes
		// var data TrazeDetail
		var asset1 map[string]interface{}
		err = json.Unmarshal(assetJSON1, &asset1)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		// Assign value for crop
		assets["Farm"] = asset1
		// --------------------   Get Events Data   -----------------------
		var data []map[string]interface{}
		// Started To check Type as Fertilization
		filter := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"},\"sort\": [{\"Data.Date\": \"desc\"}]}", asset["FTLCID"], "Event")
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
		assets["Events"] = data
		// --------------------   Get Ingredient Data   -----------------------
		var ingredient []map[string]interface{}
		fmt.Println("Inside Length")
		log.Printf("Inside Length")
		// var ids []string
		for _, ingred := range data {
			data1 := ingred["Data"].(map[string]interface{})
			batchIdsIface := data1["Batch Id"]
			// if !ok {
			// 	log.Printf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	err0 := fmt.Errorf("continue1: invalid batch ID type, ok1 = %t", ok)
			// 	fmt.Errorf("the continue1 ingredient error %t", err0)
			// 	continue
			// }

			batchIds, ok := batchIdsIface.([]interface{}) // most likely type
			if !ok {
				// return nil, fmt.Errorf("invalid BatchId type")
				continue
			}

			fmt.Println("batchId", batchIds)
			// Started To check Type as Fertilization
			for _, id := range batchIds {
				batchId := id
				filter1 := fmt.Sprintf("{\"selector\":{\"Data.BatchId\":\"%s\",\"DocType\":\"%s\"}}", batchId, "TransformationProduct")
				resultsIterator1, err := ctx.GetStub().GetQueryResult(filter1)
				if err != nil {
					return nil, fmt.Errorf("error ingredient GetQueryResult: %v", err)
				}
				defer resultsIterator1.Close()
				for resultsIterator1.HasNext() {
					queryResponse, err := resultsIterator1.Next()
					if err != nil {
						return nil, err
					}

					var asset1 map[string]interface{}
					err = json.Unmarshal(queryResponse.Value, &asset1)
					if err != nil {
						return nil, fmt.Errorf("the unmarshall ingredient error %s", err)
					}
					ingredient = append(ingredient, asset1)
					fmt.Println("Appended ingredient")
				}
			}
		}
		assets["Ingredient"] = ingredient

	}

	return assets, nil
}

// GetAllHarvest returns all assets found in world state
func (s *SmartContract) GetAllProductListEventsById(ctx contractapi.TransactionContextInterface, filter string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
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

		var data2 []map[string]interface{}
		// Started To check Type as Fertilization
		filterEvent := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"},\"sort\": [{\"Headers.eventWhen\": \"desc\"},{\"Headers.UnixTimeStamp\": \"desc\"}]}", asset["FTLCID"].(string), "Event")
		// filter := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Fertilization")
		resultsIterator2, err := ctx.GetStub().GetQueryResult(filterEvent)
		if err != nil {
			return nil, err
		}
		defer resultsIterator2.Close()

		// var assets []map[string]interface{}
		for resultsIterator2.HasNext() {
			queryResponse2, err := resultsIterator2.Next()
			if err != nil {
				return nil, err
			}

			var asset2 map[string]interface{}
			err = json.Unmarshal(queryResponse2.Value, &asset2)
			if err != nil {
				return nil, err
			}
			data2 = append(data2, asset2)
		}
		asset["Events"] = data2
		asset["EventsCount"] = len(data2)
		data = append(data, asset)
	}
	return data, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) TrazeKdesTransfer(ctx contractapi.TransactionContextInterface, id string, typeOrg string, toUserId int, toUserName string, ownerId string, toOwnerName string, fromUserId int, fromUserName string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the data %s does not exist", id)
	}
	var kdes map[string]interface{}
	err = json.Unmarshal(assetJSON, &kdes)
	if err != nil {
		return false, fmt.Errorf("unmarshall farm data: %v", err)
	}
	// Get Owner Name
	assetJSON2, err2 := ctx.GetStub().GetState(kdes["ParentId"].(string))
	if err2 != nil {
		return false, fmt.Errorf("failed to read data from world state: %v", err)
	}
	if assetJSON2 == nil {
		return false, fmt.Errorf("the data %s does not exist", id)
	}
	var owner map[string]interface{}
	err2 = json.Unmarshal(assetJSON2, &owner)
	if err2 != nil {
		return false, fmt.Errorf("unmarshall farm data: %v", err2)
	}

	if typeOrg == "Producer" {
		kdes["Status"] = "Transferred"
		kdes["UserId"] = toUserId
		kdes["UserName"] = toUserName
		kdes["ToParentId"] = ownerId
		kdes["FromUserName"] = fromUserName
		kdes["FromUserId"] = fromUserId
		kdes["FarmName"] = toOwnerName
		kdes["IsAccept"] = 1
		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
		fmt.Println("Producer1")
		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
		if err1 != nil {
			return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
	}
	if typeOrg == "Processor" {
		// Started To check Type as Fertilization
		filterEvent := fmt.Sprintf("{\"selector\":{\"ParentId\":\"%s\",\"DocType\":\"%s\"},\"sort\": [{\"Headers.eventWhen\": \"desc\"},{\"Headers.UnixTimeStamp\": \"desc\"}]}", id, "Event")
		resultsIterator2, err := ctx.GetStub().GetQueryResult(filterEvent)
		if err != nil {
			return false, err
		}
		defer resultsIterator2.Close()

		// var assets []map[string]interface{}
		for resultsIterator2.HasNext() {
			queryResponse2, err := resultsIterator2.Next()
			if err != nil {
				return false, err
			}

			var asset2 map[string]interface{}
			err = json.Unmarshal(queryResponse2.Value, &asset2)
			if err != nil {
				return false, err
			}
			asset2["UserId"] = toUserId
			asset2["UserName"] = toUserName
			assetJSON2, err4 := json.Marshal(asset2)
			if err4 != nil {
				return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
			}
			ctx.GetStub().PutState(asset2["FTLCID"].(string), assetJSON2)
		}
		kdes["Status"] = "Transferred"
		kdes["UserId"] = toUserId
		kdes["UserName"] = toUserName
		kdes["ToParentId"] = ownerId
		kdes["FromUserName"] = fromUserName
		kdes["FromUserId"] = fromUserId
		kdes["FarmName"] = toOwnerName
		kdes["IsAccept"] = 1
		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org2MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
		if err1 != nil {
			return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
	}
	if typeOrg == "Distributor" {
		kdes["Status"] = "Transferred"
		kdes["UserId"] = toUserId
		kdes["UserName"] = toUserName
		kdes["ParentId"] = ownerId
		kdes["FromUserName"] = fromUserName
		kdes["FromUserId"] = fromUserId
		kdes["FarmName"] = toOwnerName
		kdes["IsAccept"] = 1
		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org4MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
		if err1 != nil {
			return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
	}
	if typeOrg == "Retailer" {
		kdes["Status"] = "Transferred"
		kdes["UserId"] = toUserId
		kdes["UserName"] = toUserName
		kdes["ParentId"] = ownerId
		kdes["FromUserName"] = fromUserName
		kdes["FromUserId"] = fromUserId
		kdes["FarmName"] = toOwnerName
		kdes["IsAccept"] = 1
		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org5MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
		if err1 != nil {
			return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
	}
	return true, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) TransferedIsAccept(ctx contractapi.TransactionContextInterface, id string, value int) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the data %s does not exist", id)
	}
	var kdes map[string]interface{}
	err = json.Unmarshal(assetJSON, &kdes)
	if err != nil {
		return false, fmt.Errorf("unmarshall farm data: %v", err)
	}
	kdes["IsAccept"] = value
	assetJSON2, err4 := json.Marshal(kdes)
	if err4 != nil {
		return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
	}
	ctx.GetStub().PutState(id, assetJSON2)

	return true, nil
}
func (s *SmartContract) TransferedStatus(ctx contractapi.TransactionContextInterface, id string, value string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the data %s does not exist", id)
	}
	var kdes map[string]interface{}
	err = json.Unmarshal(assetJSON, &kdes)
	if err != nil {
		return false, fmt.Errorf("unmarshall farm data: %v", err)
	}
	kdes["Status"] = value
	assetJSON2, err4 := json.Marshal(kdes)
	if err4 != nil {
		return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
	}
	ctx.GetStub().PutState(id, assetJSON2)

	return true, nil
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
func (s *SmartContract) CheckFarmEmail(ctx contractapi.TransactionContextInterface, email string) (bool, error) {

	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"Farmer.ContactInformation.Email\":\"%s\"}}", "Farm", email)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return true, err
	}
	defer resultsIterator.Close()

	var farms []map[string]interface{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return true, fmt.Errorf("failed to iterate farm: %v", err)
		}
		// fmt.log("queryResponse.Value", queryResponse.Value)
		var farm map[string]interface{}
		err = json.Unmarshal(queryResponse.Value, &farm)
		if err != nil {
			return true, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farms = append(farms, farm)
	}
	var result bool
	if len(farms) != 0 {
		result = true
	} else {
		result = false
	}
	return result, nil
}

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
