package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type Farmer struct {
	FarmerName         string                    `json:"FarmerName"`
	ContactInformation *FarmerContactInformation `json:"ContactInformation"`
}

type FarmerContactInformation struct {
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

type Location struct {
	Address     string              `json:"Address"`
	Coordinates *LocationCoordinate `json:"Coordinates"`
}

type LocationCoordinate struct {
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

type CultivationPractices struct {
	SoilType         string   `json:"SoilType"`
	IrrigationMethod string   `json:"IrrigationMethod"`
	FertilizersUsed  []string `json:"FertilizersUsed"`
	PesticidesUsed   []string `json:"PesticidesUsed"`
}
type BlockchainInfo struct {
	TransactionID string `json:"TransactionID"`
	BlockNumber   string `json:"BlockNumber"`
}
type Farm struct {
	FarmID               string                `json:"FarmID"`
	Farmer               *Farmer               `json:"Farmer"`
	Location             *Location             `json:"Location"`
	FarmSize             string                `json:"FarmSize"`
	CultivationPractices *CultivationPractices `json:"CultivationPractices"`
	Certifications       []string              `json:"Certifications"`
	BlockchainInfo       *BlockchainInfo       `json:"BlockchainInfo"`
	IsDelete             int                   `json:"IsDelete"`
	DocType              string                `json:"DocType"`
}

type CropDetails struct {
	CropID          string          `json:"CropID"`
	FarmBy          string          `json:"FarmBy"`
	CropType        string          `json:"CropType"`
	PlantingDate    string          `json:"PlantingDate"`
	PesticidesUsed  []string        `json:"PesticidesUsed"`
	CropCondition   string          `json:"CropCondition"`
	Certification   []string        `json:"Certification"`
	BlockchainInfos *BlockchainInfo `json:"BlockchainInfos"`
	IsDelete        int             `json:"IsDelete"`
	DocType         string          `json:"DocType"`
}

type FertilizerPesticideEvent struct {
	CropID            string          `json:"CropID"`
	EventID           string          `json:"EventID"`
	EventType         string          `json:"EventType"`
	EventDate         string          `json:"EventDate"`
	Details           string          `json:"Details"`
	ResponsibleParty  string          `json:"ResponsibleParty"`
	QuantityUsed      string          `json:"QuantityUsed"`
	Unit              int             `json:"Unit"`
	ApplicationMethod string          `json:"ApplicationMethod"`
	BlockchainInfos   *BlockchainInfo `json:"BlockchainInfos"`
}

type IrrigationEvent struct {
	CropID            string          `json:"CropID"`
	EventID           string          `json:"EventID"`
	EventType         string          `json:"EventType"`
	EventDate         string          `json:"EventDate"`
	Details           string          `json:"Details"`
	ResponsibleParty  string          `json:"ResponsibleParty"`
	QuantityUsed      string          `json:"QuantityUsed"`
	Unit              int             `json:"Unit"`
	ApplicationMethod string          `json:"ApplicationMethod"`
	BlockchainInfos   *BlockchainInfo `json:"BlockchainInfos"`
}

type QualityAssessment struct {
	Size             float64 `json:"Size"`
	Color            float64 `json:"Color"`
	OverallCondition float64 `json:"OverallCondition"`
}
type HarvestingEvent struct {
	FarmID            string             `json:"FarmID"`
	CropID            string             `json:"CropID"`
	EventID           string             `json:"EventID"`
	EventType         string             `json:"EventType"`
	EventDate         string             `json:"EventDate"`
	QuantityHarvested int                `json:"QuantityHarvested"`
	HarvestedBy       string             `json:"HarvestedBy"`
	HarvestConditions string             `json:"HarvestConditions"`
	StorageConditions string             `json:"StorageConditions"`
	QualityAssessment *QualityAssessment `json:"QualityAssessment"`
	BlockchainInfos   *BlockchainInfo    `json:"BlockchainInfos"`
}
type NutritionalContent struct {
	VitaminC string `json:"VitaminC"`
	Iron     string `json:"Iron"`
	Calcium  string `json:"Calcium"`
}
type TestResults struct {
	PesticideResidue       string              `json:"PesticideResidue"`
	NutritionalContent     *NutritionalContent `json:"CropID"`
	MicrobialContamination string              `json:"MicrobialContamination"`
	AllergenPresence       string              `json:"AllergenPresence"`
}
type LabTestingEvent struct {
	FarmID          string          `json:"FarmID"`
	CropID          string          `json:"CropID"`
	EventID         string          `json:"EventID"`
	EventType       string          `json:"EventType"`
	EventDate       string          `json:"EventDate"`
	Details         string          `json:"Details"`
	TestedBy        string          `json:"TestedBy"`
	TestResults     *TestResults    `json:"TestResults"`
	BlockchainInfos *BlockchainInfo `json:"BlockchainInfos"`
}
type TraceEvent struct {
	MetaInfo    *MetaInfo    `json:"metaInfo"`
	Header      *Header      `json:"header"`
	Products    *Products    `json:"Products"`
	HarvestData *HarvestData `json:"HarvestData"`
}

type ProductSegregate struct {
	MetaInfo *MetaInfo `json:"metaInfo"`
	Header   *Header   `json:"header"`
	Products *Products `json:"Products"`
}

type Header struct {
	EventID    string `json:"eventId"`
	EventWhy   string `json:"eventWhy"`
	EventWhen  string `json:"eventWhen"`
	EventWhere string `json:"eventWhere"`
}

type Products struct {
	FarmId       string  `json:"farmId"`
	ProductId    string  `json:"productId"`
	BatchNumber  string  `json:"batchNumber"`
	Quantity     string  `json:"quantity"`
	ProductName  string  `json:"productName"`
	TotalWeight  float32 `json:"weight"`
	LogisticData string  `json:"logisticData"`
	TraceId      string  `json:"traceId"`
	Status       string  `json:"status"`
	IpfsCid      string  `json:"ipfsCid"`
}

type HarvestData struct {
	HarvestId                  string `json:"harvestId"`
	CropInformation            string `json:"cropInformation"`
	LabData                    string `json:"labData"`
	PlantingDate               string `json:"plantingDate"`
	HarvestingDate             string `json:"harvestingDate"`
	FertilizerUsage            string `json:"fertilizerUsage"`
	YieldPerAcre               string `json:"yieldPerAcre"`
	QualityAndSafety           string `json:"qualityAndSafety"`
	CertificationAndComplaince string `json:"certificationAndComplaince"`
	EnvironmentalDate          string `json:"environmentalDate"`
}

type MetaInfo struct {
	Owner        string `json:"owner"`
	Organisation string `json:"organisation"`
	MspId        string `json:"mspId"`
	EventType    string `json:"eventType"`
	Action       string `json:"action"`
	CreatedDate  string `json:"createdDate"`
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

type Attributes struct {
	ProductName string `json:"ProductName"`
	TotalWeight string `json:"TotalWeight"`
	Quantity    string `json:"Quantity"`
	EventWhy    string `json:"EventWhy"`
	EventWhen   string `json:"EventWhen"`
	EventWhere  string `json:"EventWhere"`
}

type Lab struct {
	PesticideResidueAnalysis int    `json:"PesticideResidueAnalysis"`
	MicrobialTesting         string `json:"MicrobialTesting"`
	AntioxidantLevels        string `json:"AntioxidantLevels"`
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    *TraceEvent `json:"record"`
	TxId      string      `json:"txId"`
	Timestamp time.Time   `json:"timestamp"`
	IsDelete  bool        `json:"isDelete"`
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

func (s *SmartContract) FoodTrazeCreate(ctx contractapi.TransactionContextInterface, status string, data1 string, data2 string, data3 string, data4 string, data5 string, data6 string, data7 string, data8 string, data9 string) (interface{}, error) {
	var response FoodTazeRes
	if status == "CropCreateEvent" {

		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		if err1 := json.Unmarshal([]byte(data8), &blockChainInfo); err1 != nil {
			// fmt.Println("Error parsing JSON:", err)
			return nil, err1
		}
		PesticidesUsed := strings.Split(data5, ",")
		arrCertificate := strings.Split(data7, ",")
		// feetFloat, _ := strconv.ParseFloat("3.2", 32)
		asset := CropDetails{
			FarmBy:          data1,
			CropID:          data2,
			CropType:        data3,
			PlantingDate:    data4,
			PesticidesUsed:  PesticidesUsed,
			CropCondition:   data6,
			Certification:   arrCertificate,
			BlockchainInfos: &blockChainInfo,
			IsDelete:        0,
			DocType:         "Crop",
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", asset.CropID)
		}

		result := ctx.GetStub().PutState(asset.CropID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.CropID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}

		response = FoodTazeRes{
			Status:  200,
			Message: "Crop Event Created Successfully.",
			Data:    result,
		}

	}

	if status == "FarmCreateEvent" {
		// Parse JSON data into Asset struct
		var farmerInformatioData FarmerContactInformation
		if err := json.Unmarshal([]byte(data2), &farmerInformatioData); err != nil {
			// fmt.Println("Error parsing JSON:", err)
			return nil, fmt.Errorf("the contact information error %v", err)
		}
		// Parse JSON data into Asset struct
		farmerData := Farmer{
			FarmerName:         data3,
			ContactInformation: &farmerInformatioData,
		}
		// Parse JSON data into Asset struct
		var locationCoridinates LocationCoordinate
		if err := json.Unmarshal([]byte(data4), &locationCoridinates); err != nil {
			// fmt.Println("Error parsing JSON:", err)
			return nil, fmt.Errorf("the location coordinate error %v", err)
		}
		locationData := Location{
			Address:     data5,
			Coordinates: &locationCoridinates,
		}
		// // Parse JSON data into Asset struct
		var cultivationPracticeData CultivationPractices
		if err1 := json.Unmarshal([]byte(data7), &cultivationPracticeData); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the cultivation practice error %v", err1)
		}
		arrCertificate := strings.Split(data8, ",")

		// // Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		if err1 := json.Unmarshal([]byte(data9), &blockChainInfo); err1 != nil {
			// fmt.Println("Error parsing JSON:", err)
			return nil, fmt.Errorf("the blockchain info error %v", err1)
		}

		asset := Farm{
			FarmID:               data1,
			Farmer:               &farmerData,
			Location:             &locationData,
			FarmSize:             data6,
			CultivationPractices: &cultivationPracticeData,
			Certifications:       arrCertificate,
			BlockchainInfo:       &blockChainInfo,
			IsDelete:             0,
			DocType:              "Farm",
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data1)
		}

		// farmKey, err := ctx.GetStub().CreateCompositeKey("Farm", []string{data1})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to create composite key: %v", err)
		// }

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(data1, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, data1, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Farm Event Created Successfully.",
			Data:    result,
		}
	}

	if status == "FertilizerPesticideEvent" {
		unit, _ := strconv.Atoi(data8)
		asset := FertilizerPesticideEvent{
			CropID:            data1,
			EventID:           data2,
			EventType:         data3,
			EventDate:         data4,
			Details:           data5,
			ResponsibleParty:  data6,
			QuantityUsed:      data7,
			Unit:              unit,
			ApplicationMethod: data9,
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}

		// farmKey, err := ctx.GetStub().CreateCompositeKey("Farm", []string{data1})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to create composite key: %v", err)
		// }

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(data2, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, data2, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Fertilizer Pesticide Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "IrrigationEvent" {
		unit, _ := strconv.Atoi(data8)
		asset := IrrigationEvent{
			CropID:            data1,
			EventID:           data2,
			EventType:         data3,
			EventDate:         data4,
			Details:           data5,
			ResponsibleParty:  data6,
			QuantityUsed:      data7,
			Unit:              unit,
			ApplicationMethod: data9,
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}

		// farmKey, err := ctx.GetStub().CreateCompositeKey("Farm", []string{data1})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to create composite key: %v", err)
		// }

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(data2, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, data2, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Irrigation Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "HarvestingEvent" {
		// // Parse JSON data into Asset struct
		var qualityAssessment QualityAssessment
		if err1 := json.Unmarshal([]byte(data9), &qualityAssessment); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the quality assessment error %v", err1)
		}
		quantity, _ := strconv.Atoi(data5)
		asset := HarvestingEvent{
			CropID:            data1,
			EventID:           data2,
			EventType:         data3,
			EventDate:         data4,
			QuantityHarvested: quantity,
			HarvestedBy:       data6,
			HarvestConditions: data7,
			StorageConditions: data8,
			QualityAssessment: &qualityAssessment,
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}

		// farmKey, err := ctx.GetStub().CreateCompositeKey("Farm", []string{data1})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to create composite key: %v", err)
		// }

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(data2, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, data2, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Harvesting Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "LabTestingEvent" {
		// // Parse JSON data into Asset struct
		var nutritionalContent NutritionalContent
		if err1 := json.Unmarshal([]byte(data7), &nutritionalContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the quality assessment error %v", err1)
		}
		// Parse JSON data into Asset struct
		testResultData := TestResults{
			PesticideResidue:       data6,
			NutritionalContent:     &nutritionalContent,
			MicrobialContamination: data8,
			AllergenPresence:       data9,
		}
		asset := LabTestingEvent{
			CropID:      data1,
			EventID:     data2,
			EventType:   data3,
			EventDate:   data4,
			Details:     data5,
			TestedBy:    data6,
			TestResults: &testResultData,
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}

		// farmKey, err := ctx.GetStub().CreateCompositeKey("Farm", []string{data1})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to create composite key: %v", err)
		// }

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(data2, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, data2, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Harvesting Event Created Successfully.",
			Data:    result,
		}
	}
	return response, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetAllFarms(ctx contractapi.TransactionContextInterface) ([]*Farm, error) {

	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\"}}", "Farm")
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var farms []*Farm
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate farm: %v", err)
		}
		// fmt.log("queryResponse.Value", queryResponse.Value)
		var farm Farm
		err = json.Unmarshal(queryResponse.Value, &farm)
		if err != nil {
			return nil, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farms = append(farms, &farm)
	}
	return farms, nil
}

// Function to retrieve all Struct1 data
func (s *SmartContract) GetAllStruct1(ctx contractapi.TransactionContextInterface) ([]*Farm, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var structs []*Farm
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var data Farm
		if err := json.Unmarshal(queryResponse.Value, &data); err != nil {
			return nil, err
		}
		structs = append(structs, &data)
	}
	return structs, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadFarmById(ctx contractapi.TransactionContextInterface, id string) (*Farm, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the farm %s does not exist", id)
	}

	var asset Farm
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the farm unmarshall error %v", err)
	}

	return &asset, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetAllCropsList(ctx contractapi.TransactionContextInterface) ([]*CropDetails, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\"}}", "Crop")
	// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*CropDetails
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset CropDetails
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil

}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadCropById(ctx contractapi.TransactionContextInterface, id string) (*CropDetails, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset CropDetails
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}

	return &asset, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetAllCropsByFarmId(ctx contractapi.TransactionContextInterface, farmId string) ([]*CropDetails, error) {
	exists, err2 := s.AssetExists(ctx, farmId)
	if err2 != nil {
		return nil, fmt.Errorf("the Farm Data %s exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the Farm Data %s not exists", farmId)
	}
	queryString := fmt.Sprintf("{\"selector\":{\"FarmBy\":\"%s\"}}", farmId)
	// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*CropDetails
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset CropDetails
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (s *SmartContract) CreateCrop(ctx contractapi.TransactionContextInterface, data1 string, data2 string, data3 string, data4 string, data5 string, data6 string, data7 string, data8 string) (interface{}, error) {

	// Parse JSON data into Asset struct
	var blockChainInfo BlockchainInfo
	if err1 := json.Unmarshal([]byte(data8), &blockChainInfo); err1 != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, err1
	}
	PesticidesUsed := strings.Split(data5, ",")
	arrCertificate := strings.Split(data7, ",")

	asset := CropDetails{
		FarmBy:          data1,
		CropID:          data2,
		CropType:        data3,
		PlantingDate:    data4,
		PesticidesUsed:  PesticidesUsed,
		CropCondition:   data6,
		Certification:   arrCertificate,
		BlockchainInfos: &blockChainInfo,
	}
	assetJSON, err4 := json.Marshal(asset)
	if err4 != nil {
		return nil, fmt.Errorf("the asset json %s already exists", asset.CropID)
	}

	result := ctx.GetStub().PutState(asset.CropID, assetJSON)

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org1MSP"}
	err1 := setAssetStateBasedEndorsement(ctx, asset.CropID, endorsingOrgs)
	if err1 != nil {
		return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	}

	rs := &TraceEventRes{
		EventID: asset.CropID,
		// TraceId: HeaderData.TraceID,
		Data: result,
	}
	return rs, nil
}
func (s *SmartContract) CreateFarm1(ctx contractapi.TransactionContextInterface, farmId string, farmer string, farmerContactInformation string, location string, locationCoridinate string, farmSize string, cultivationPractices string, certifications string, blockchainInfo string) (interface{}, error) {
	// Parse JSON data into Asset struct
	var farmerInformatioData FarmerContactInformation
	if err := json.Unmarshal([]byte(farmerContactInformation), &farmerInformatioData); err != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the farmData error", err)
	}
	// Parse JSON data into Asset struct
	farmerData := Farmer{
		FarmerName:         farmer,
		ContactInformation: &farmerInformatioData,
	}
	// Parse JSON data into Asset struct
	var locationCoridinates LocationCoordinate
	if err := json.Unmarshal([]byte(locationCoridinate), &locationCoridinates); err != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the farmData error1", err)
	}
	locationData := Location{
		Address:     location,
		Coordinates: &locationCoridinates,
	}
	// if err := json.Unmarshal([]byte(farmer), &farmerData); err != nil {
	// 	// fmt.Println("Error parsing JSON:", err)
	// 	return nil, fmt.Errorf("the farmData error", err)
	// }
	// // Parse JSON data into Asset struct
	// var locationData Location
	// if err1 := json.Unmarshal([]byte(location), &locationData); err1 != nil {
	// 	// fmt.Println("Error parsing JSON:", err)
	// 	return nil, err1
	// }
	// // Parse JSON data into Asset struct
	var cultivationPracticeData CultivationPractices
	if err1 := json.Unmarshal([]byte(cultivationPractices), &cultivationPracticeData); err1 != nil {
		// fmt.Println("Error parsing JSON1:", err1)
		return nil, fmt.Errorf("the farmData error2", err1)
	}
	// // Parse JSON data into Asset struct
	var blockChainInfo BlockchainInfo
	if err1 := json.Unmarshal([]byte(blockchainInfo), &blockChainInfo); err1 != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the farmData error3", err1)
	}

	arrCertificate := strings.Split(certifications, ",")

	asset := Farm{
		FarmID:               farmId,
		Farmer:               &farmerData,
		Location:             &locationData,
		FarmSize:             farmSize,
		CultivationPractices: &cultivationPracticeData,
		Certifications:       arrCertificate,
		BlockchainInfo:       &blockChainInfo,
	}
	assetJSON, err4 := json.Marshal(asset)
	if err4 != nil {
		return nil, fmt.Errorf("the asset json %s already exists", farmId)
	}

	result := ctx.GetStub().PutState(farmId, assetJSON)

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org1MSP"}
	err1 := setAssetStateBasedEndorsement(ctx, farmId, endorsingOrgs)
	if err1 != nil {
		return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	}

	rs := &TraceEventRes{
		EventID: farmId,
		// TraceId: HeaderData.TraceID,
		Data: result,
	}
	return rs, nil
}

func (s *SmartContract) CreateAssetHarvest(ctx contractapi.TransactionContextInterface, data1 string, data2 string, data3 string) (interface{}, error) {

	// Parse JSON data into Asset struct
	var HeaderData Header
	if err := json.Unmarshal([]byte(data1), &HeaderData); err != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the headers error", err)
	}
	// Parse JSON data into Asset struct
	var metaData MetaInfo
	if err1 := json.Unmarshal([]byte(data2), &metaData); err1 != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, err1
	}
	// metaData.EventType = "cropInformation"
	// metaData.CreatedDate, _ = time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
	var HarvestDatas HarvestData
	if err3 := json.Unmarshal([]byte(data3), &HarvestDatas); err3 != nil {
		fmt.Println("Error parsing JSON:", err3)
		return nil, fmt.Errorf("the HarvestDatas error", err3)
	}
	exists, err2 := s.AssetExists(ctx, HeaderData.EventID)
	if err2 != nil {
		return nil, fmt.Errorf("the Harvest Datas exist error", err2)
	}
	if exists {
		return nil, fmt.Errorf("the Harvest Datas %s already exists", HeaderData.EventID)
	}
	asset := TraceEvent{
		Header:      &HeaderData,
		MetaInfo:    &metaData,
		HarvestData: &HarvestDatas,
	}
	assetJSON, err4 := json.Marshal(asset)
	if err4 != nil {
		return nil, fmt.Errorf("the asset json %s already exists", asset)
	}

	result := ctx.GetStub().PutState(HeaderData.EventID, assetJSON)

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org1MSP"}
	err1 := setAssetStateBasedEndorsement(ctx, HeaderData.EventID, endorsingOrgs)
	if err1 != nil {
		return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	}

	rs := &TraceEventRes{
		EventID: HeaderData.EventID,
		// TraceId: HeaderData.TraceID,
		Data: result,
	}
	return rs, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllHarvestorEvents(ctx contractapi.TransactionContextInterface) ([]*TraceEvent, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*TraceEvent
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset TraceEvent
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (s *SmartContract) UpdateAssetProducts(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	// Parse JSON data into Asset struct
	var ProductData Products
	if err := json.Unmarshal([]byte(data), &ProductData); err != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the headers error", err)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	asset.Products = &ProductData

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return nil, err
	}

	ctx.GetStub().PutState(id, assetJSON)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvesterData(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	// Parse JSON data into Asset struct
	var HarvestDatas HarvestData
	if err3 := json.Unmarshal([]byte(data), &HarvestDatas); err3 != nil {
		fmt.Println("Error parsing JSON:", err3)
		return nil, fmt.Errorf("the HarvestDatas error", err3)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	asset.HarvestData = &HarvestDatas

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return nil, err
	}

	ctx.GetStub().PutState(id, assetJSON)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
	}
	return rs, nil
}

func (s *SmartContract) CreateAssetProduct(ctx contractapi.TransactionContextInterface, data1 string, data2 string, data3 string) (interface{}, error) {

	// Parse JSON data into Asset struct
	var HeaderData Header
	if err := json.Unmarshal([]byte(data1), &HeaderData); err != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the headers error", err)
	}
	exists, err2 := s.AssetExists(ctx, HeaderData.EventID)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if exists {
		return nil, fmt.Errorf("the asset %s already exists", HeaderData.EventID)
	}
	// Parse JSON data into Asset struct
	var metaData MetaInfo
	if err1 := json.Unmarshal([]byte(data2), &metaData); err1 != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, err1
	}
	// metaData.EventType = "Product"
	// Parse JSON data into Asset struct
	var ProductData Products
	if err := json.Unmarshal([]byte(data3), &ProductData); err != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the headers error", err)
	}
	// var HarvestDatas HarvestData
	// if err3 := json.Unmarshal([]byte(data3), &HarvestDatas); err3 != nil {
	// 	fmt.Println("Error parsing JSON:", err3)
	// 	return nil, fmt.Errorf("the HarvestDatas error", err3)
	// }

	asset := TraceEvent{
		Header:   &HeaderData,
		MetaInfo: &metaData,
		Products: &ProductData,
	}
	assetJSON, err4 := json.Marshal(asset)
	if err4 != nil {
		return nil, fmt.Errorf("the asset json %s already exists", asset)
	}

	ctx.GetStub().PutState(HeaderData.EventID, assetJSON)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
	}
	return rs, nil
}

func (s *SmartContract) CreatedProductSegregate(ctx contractapi.TransactionContextInterface, id, eventId, quantity string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, eventId)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", eventId)
	}
	// Parse JSON data into Asset struct
	// var ProductData Products
	// if err := json.Unmarshal([]byte(data), &ProductData); err != nil {
	// 	// fmt.Println("Error parsing JSON:", err)
	// 	return nil, fmt.Errorf("the headers error", err)
	// }
	asset, err := s.ReadAssetEvent(ctx, eventId)
	if err != nil {
		return nil, err
	}

	strInit := regexp.MustCompile(`[^0-9 ]+`).ReplaceAllString(strings.ReplaceAll(string(asset.Products.Quantity), " ", ""), "")
	initialQuantity, _ := strconv.Atoi(strInit)
	strParam := regexp.MustCompile(`[^0-9 ]+`).ReplaceAllString(strings.ReplaceAll(string(quantity), " ", ""), "")
	paramQuantity, _ := strconv.Atoi(strParam)
	if initialQuantity >= paramQuantity {

		reduceValue := (initialQuantity - paramQuantity)
		asset.Products.Quantity = strconv.Itoa(reduceValue) + " kg"
		assetJSON, err1 := json.Marshal(asset)
		if err1 != nil {
			return nil, err1
		}
		ctx.GetStub().PutState(eventId, assetJSON)
		// prodQuantity := strconv.Itoa(paramQuantity)
		product := &ProductSegregate{
			Header:   asset.Header,
			MetaInfo: asset.MetaInfo,
			Products: asset.Products,
		}
		product.Header.EventID = id
		product.Products.Quantity = quantity
		product.Products.Status = "Created"
		product.Products.TraceId = eventId
		assetJSONLast, err1 := json.Marshal(product)
		if err1 != nil {
			return nil, err1
		}

		ctx.GetStub().PutState(id, assetJSONLast)
		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
		if err != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}

	} else {
		rs := &TraceEventErrorRes{
			Status:  0,
			Message: "Quantity was minimum not able to segregate",
			// TraceId: asset.Header.TraceID,
		}
		return rs, nil
	}

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
	}
	return rs, nil
}

// AssetExists returns true when asset with given ID exists in world state
// func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
// 	assetJSON, err := ctx.GetStub().GetState(id)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to read from world state: %v", err)
// 	}

// 	return assetJSON != nil, nil
// }

// TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
func (s *SmartContract) TransferHarvestAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string, mspId string) (interface{}, error) {
	asset, err := s.ReadProductAsset(ctx, id)
	if err != nil {
		return "", err
	}
	asset.MetaInfo.Owner = newOwner
	asset.MetaInfo.Organisation = "Org2"
	asset.MetaInfo.MspId = mspId

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org2MSP"}
	err = setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
	if err != nil {
		return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err)
	}

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
	}

	return rs, nil
}

func (s *SmartContract) UpdateHarvestProduct(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {

	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.Products.LogisticData = data

	assetAsBytes, _ := json.Marshal(asset)

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to mint new tokens
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("failed to get MSPID: %v", err)
	}
	// if clientMSPID != "Org1MSP" {
	// 	return nil, fmt.Errorf("client is not authorized to mint new tokens")
	// }

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		TraceId: clientMSPID,
		Data:    nil,
	}
	return rs, nil
}

func (s *SmartContract) CreateAssetHarvestFertilizer(ctx contractapi.TransactionContextInterface, id string, data1 string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.FertilizerUsage = data1
	asset.MetaInfo.EventType = "fertilizerUsage"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) CreateHarvestPlantingDate(ctx contractapi.TransactionContextInterface, id string, data1 string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.PlantingDate = data1
	asset.MetaInfo.EventType = "plantingDate"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) CreateHarvestDate(ctx contractapi.TransactionContextInterface, id string, data1 string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.HarvestingDate = data1
	asset.MetaInfo.EventType = "harvestingDate"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvestCropInfo(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.CropInformation = data
	asset.MetaInfo.EventType = "cropInformation"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvestLabData(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.LabData = data
	asset.MetaInfo.EventType = "labData"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvestYieldPerAcre(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.YieldPerAcre = data
	asset.MetaInfo.EventType = "yieldPerAcre"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvestQualityAndSafety(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.QualityAndSafety = data
	asset.MetaInfo.EventType = "qualityAndSafety"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvestCertificationAndComplaince(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.CertificationAndComplaince = data
	asset.MetaInfo.EventType = "certificationAndComplaince"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

func (s *SmartContract) UpdateHarvestEnvironmentalDate(ctx contractapi.TransactionContextInterface, id string, data string) (interface{}, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return nil, fmt.Errorf("the asset exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s already exists", id)
	}
	asset, err := s.ReadAssetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	asset.HarvestData.EnvironmentalDate = data
	asset.MetaInfo.EventType = "environmentalDate"

	assetAsBytes, _ := json.Marshal(asset)

	ctx.GetStub().PutState(id, assetAsBytes)

	rs := &TraceEventRes{
		EventID: asset.Header.EventID,
		// TraceId: asset.Header.TraceID,
		Data: nil,
	}
	return rs, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAssetEvent(ctx contractapi.TransactionContextInterface, id string) (*TraceEvent, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset TraceEvent
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error", err)
	}

	return &asset, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadProductAsset(ctx contractapi.TransactionContextInterface, id string) (*ProductSegregate, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset ProductSegregate
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error", err)
	}

	return &asset, nil
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// GetAssetHistory returns the chain of custody for an asset since issuance.
func (s *SmartContract) GetAssetHistory(ctx contractapi.TransactionContextInterface, assetID string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", assetID)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset TraceEvent
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
			Record:    &asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
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

	return nil
}

func generateUniqueAssetID() string {
	// Implement your logic for generating a unique asset ID
	// Example: timestamp + random number
	return strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(1000))
}
