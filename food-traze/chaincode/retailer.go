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

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset4 struct {
	AppraisedValue int        `json:"AppraisedValue"`
	Color          string     `json:"Color"`
	ID             string     `json:"ID"`
	Owner          string     `json:"Owner"`
	Size           int        `json:"Size"`
	Data           Attribute4 `json:"Attributes"`
	Test           Lab4       `json:"Labtest"`
	Val            string
}

type Attribute4 struct {
	Quantity  int    `json:"Quantity"`
	FruitType string `json:"FruitType"`
	FruitName string `json:"FruitName"`
}

type Lab4 struct {
	PesticideResidueAnalysis int    `json:"PesticideResidueAnalysis"`
	MicrobialTesting         string `json:"MicrobialTesting"`
	AntioxidantLevels        string `json:"AntioxidantLevels"`
}

// InitLedger adds a base set of assets to the ledger
// func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
// 	asset2 := Attribute4{
// 		Quantity:  9,
// 		FruitType: "deciduous",
// 		FruitName: "Apple",
// 	}

// 	assets := []Asset4{
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
// 		err = setAssetStateBasedEndorsement4(ctx, asset.ID, endorsingOrgs)
// 		if err != nil {
// 			return fmt.Errorf("failed setting state based endorsement for new owner: %v", err)
// 		}
// 	}

// 	return nil
// }

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAssetRetailer(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int, data string, lab string) error {

	val, ok, err := ctx.GetClientIdentity().GetAttributeValue("role")
	if err != nil {
		// There was an error trying to retrieve the attribute
	}
	if !ok {
		// The client identity does not possess the attribute
	}
	// Do something with the value of 'val'
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return fmt.Errorf("only approver attribute users can approve the contract")
	}
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	// Parse JSON data into Asset struct
	var Attributes Attribute4
	if err := json.Unmarshal([]byte(data), &Attributes); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}

	// Parse JSON lab data into Asset struct
	var Labtest Lab4
	if err1 := json.Unmarshal([]byte(lab), &Labtest); err1 != nil {
		fmt.Println("Error parsing JSON:", err1)
		return err
	}
	asset := Asset4{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
		Data:           Attributes,
		Test:           Labtest,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	result := ctx.GetStub().PutState(id, assetJSON)

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org1MSP"}
	err = setAssetStateBasedEndorsement4(ctx, asset.ID, endorsingOrgs)
	if err != nil {
		return fmt.Errorf("failed setting state based endorsement for new owner: %v", err)
	}

	return result
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAssetRetailer(ctx contractapi.TransactionContextInterface, id string) (*Asset4, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset4
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAssetRetailer(ctx contractapi.TransactionContextInterface, id string, color string, size int, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}
	// overwriting original asset with new asset
	asset := Asset4{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAssetRetailer(ctx contractapi.TransactionContextInterface, id string) error {
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
func (s *SmartContract) AssetExistsRetailer(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
func (s *SmartContract) TransferAssetRetailer(ctx contractapi.TransactionContextInterface, id string, newOwner string) (string, error) {
	asset, err := s.ReadAssetRetailer(ctx, id)
	if err != nil {
		return "", err
	}

	oldOwner := asset.Owner
	asset.Owner = newOwner

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{newOwner}
	err = setAssetStateBasedEndorsement4(ctx, asset.ID, endorsingOrgs)
	if err != nil {
		return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err)
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldOwner, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssetsRetailer(ctx contractapi.TransactionContextInterface) ([]*Asset4, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset4
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset4
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

// setAssetStateBasedEndorsement adds an endorsement policy to an asset so that the passed orgs need to agree upon transfer
func setAssetStateBasedEndorsement4(ctx contractapi.TransactionContextInterface, assetID string, orgsToEndorse []string) error {
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
