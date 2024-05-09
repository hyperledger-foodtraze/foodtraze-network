package chaincode

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

// *SmartContract provides functions for managing an Asset
// type *SmartContract struct {
// 	contractapi.Contract
// }

type TransformProduct struct {
	SKU         string `json:"SKU"`
	UPC         string `json:"UPC"`
	GTIN        string `json:"GTIN"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Quantity    string `json:"Quantity"`
	Unit        string `json:"Unit"`
}

type Supplier struct {
	Name    string `json:"Name"`
	Contact string `json:"Contact"`
}

type Ingredients struct {
	BatchNumber    string    `json:"BatchNumber"`
	Name           string    `json:"Name"`
	Unit           string    `json:"Unit"`
	CropFtlcId     string    `json:"CropFtlcId"`
	Quantity       string    `json:"Quantity"`
	RecevingFtlcId string    `json:"RecevingFtlcId"`
	Supplier       *Supplier `json:"Supplier"`
}

type Production struct {
	ProductionDate string `json:"ProductionDate"`
	ProductionTime string `json:"ProductionTime"`
	Location       string `json:"Location"`
	Equipment      string `json:"Equipment"`
}

type BatchLot struct {
	Number         string `json:"Number"`
	ExpirationDate string `json:"ExpirationDate"`
}

type QualityControlTests struct {
	Type   string `json:"Type"`
	Result string `json:"Result"`
}

type TransformShipping struct {
	ShipmentDate         string `json:"ShipmentDate"`
	TransportationMethod string `json:"TransportationMethod"`
	Destination          string `json:"Destination"`
}

type StorageConditions struct {
	Temperature string `json:"Temperature"`
	Humidity    string `json:"Humidity"`
}

type Recall struct {
	RecallStatus string `json:"RecallStatus"`
	RecallReason string `json:"RecallReason"`
}

type RegulatoryCompliance struct {
	Certifications   []string `json:"Certifications"`
	ComplianceStatus string   `json:"ComplianceStatus"`
}

type TransformationKdes struct {
	FTLCID               string                `json:"FTLCID"`
	TransformProduct     *TransformProduct     `json:"TransformProduct"`
	Ingredients          []Ingredients         `json:"Ingredients"`
	Production           *Production           `json:"Production"`
	BatchLot             *BatchLot             `json:"BatchLot"`
	QualityControlTests  []QualityControlTests `json:"QualityControlTests"`
	TranformShipping     *TransformShipping    `json:"TransformShipping"`
	StorageConditions    *StorageConditions    `json:"StorageConditions"`
	Recall               *Recall               `json:"Recall"`
	RegulatoryCompliance *RegulatoryCompliance `json:"RegulatoryCompliance"`
	Headers              *Header               `json:"Headers"`
	BlockchainInfos      *BlockchainInfo       `json:"BlockchainInfos"`
	DocType              string                `json:"DocType"`
	UserId               string                `json:"UserId"`
	AliasOrgName         string                `json:"AliasOrgName"`
}
type ProductItemInformation struct {
	BatchNumber          string `json:"BatchNumber"`
	ProductName          string `json:"ProductName"`
	Unit                 string `json:"Unit"`
	Quantity             string `json:"Quantity"`
	TransformationFtlcId string `json:"TransformationFtlcId"`
}
type ProcessorShippingingKdes struct {
	FTLCID              string                  `json:"FTLCID"`
	SenderInformation   *SenderInformation      `json:"SenderInformation"`
	ProductInformation  *ProductItemInformation `json:"ProductInformation"`
	ItemInformation     []Ingredients           `json:"ItemInformation"`
	ReceiverInformation *ReceiverInformation    `json:"ReceiverInformation"`
	CarrierInformation  *CarrierInformation     `json:"CarrierInformation"`
	Headers             *Header                 `json:"Headers"`
	BlockchainInfos     *BlockchainInfo         `json:"BlockchainInfos"`
	DocType             string                  `json:"DocType"`
	UserId              string                  `json:"UserId"`
	Status              string                  `json:"Status"`
	IsAccepted          string                  `json:"IsAccepted"`
	AliasOrgName        string                  `json:"AliasOrgName"`
}

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

func (s *SmartContract) FoodTrazeProcessorCreate(ctx contractapi.TransactionContextInterface, status string, data1 string, data2 string, data3 string, data4 string, data5 string, data6 string, data7 string, data8 string, data9 string, data10 string, data11 string, data12 string, data13 string, data14 string) (interface{}, error) {
	var response FoodTazeRes
	if status == "CropCreateEvent" {

		PesticidesUsed := strings.Split(data5, ",")
		arrCertificate := strings.Split(data7, ",")
		// feetFloat, _ := strconv.ParseFloat("3.2", 32)
		var ipfsCert []FarmFile
		if data12 != "" {
			if err1 := json.Unmarshal([]byte(data12), &ipfsCert); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		// timestamps, _ := ctx.GetStub().GetTxTimestamp()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Retrieve the block number from the transaction timestamp
		// blockNumber := timestamps.GetSeconds() / 10
		clientId, _ := ctx.GetClientIdentity().GetID()
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		// blockChainInfo.BlockNumber = clientId
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data13), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		asset := CropDetails{
			FTLCID:          data10,
			FarmBy:          data1,
			CropID:          data2,
			CropType:        data3,
			CropName:        data8,
			PlantingDate:    data4,
			HarvestingDate:  data9,
			PesticidesUsed:  PesticidesUsed,
			CropCondition:   data6,
			Certification:   arrCertificate,
			FarmFile:        ipfsCert,
			BlockchainInfos: &blockChainInfo,
			IsDelete:        0,
			UserId:          data11,
			DocType:         "Crop",
			Headers:         &headerContent,
			AliasOrgName:    "Producer",
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", asset.CropID)
		}

		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}

		response = FoodTazeRes{
			Status:  200,
			Message: "Crop Event Created Successfully.",
			Data:    result,
		}

	}

	return response, nil
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
