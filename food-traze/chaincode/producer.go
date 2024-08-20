package chaincode

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/statebased"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type Farmer struct {
	FarmerName         string                    `json:"FarmerName"`
	DateOfBirth        string                    `json:"DateOfBirth"`
	Gender             string                    `json:"Gender"`
	Address            *Address                  `json:"Address"`
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

type Address struct {
	Street     string `json:"Street"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
	PostalCode string `json:"PostalCode"`
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
	TransactionID string    `json:"TransactionID"`
	ClientId      string    `json:"ClientId"`
	BlockNumber   int64     `json:"BlockNumber"`
	ChannelId     string    `json:"ChannelId"`
	Timestamp     time.Time `json:"Timestamp"`
	MspId         string    `json:"MspId"`
}
type IpfsImage struct {
	ImageName string `json:"ImageName"`
	ImageCid  string `json:"ImageCid"`
}
type IpfsFile struct {
	FileName string `json:"FileName"`
	FileCid  string `json:"FileCid"`
}

type FarmImage struct {
	ImageName    string `json:"ImageName"`
	ImagePath    string `json:"ImagePath"`
	ImageOrgName string `json:"ImageOrgName"`
	ImageSize    string `json:"ImageSize"`
	ImageDate    string `json:"ImageDate"`
}

type FarmFile struct {
	FileName    string `json:"FileName"`
	FilePath    string `json:"FilePath"`
	FileOrgName string `json:"FileOrgName"`
	FileSize    string `json:"FileSize"`
	FileDate    string `json:"FileDate"`
}

type Farm struct {
	FTLCID   string    `json:"FTLCID"`
	FarmID   string    `json:"FarmID"`
	FarmName string    `json:"FarmName"`
	Farmer   *Farmer   `json:"Farmer"`
	Location *Location `json:"Location"`
	FarmSize string    `json:"FarmSize"`
	// CultivationPractices              *CultivationPractices `json:"CultivationPractices"`
	SoilPhLevel                       string          `json:"SoilPhLevel"`
	Certifications                    []string        `json:"Certifications"`
	BlockchainInfos                   *BlockchainInfo `json:"BlockchainInfos"`
	IsDelete                          int             `json:"IsDelete"`
	DocType                           string          `json:"DocType"`
	IpfsImage                         []IpfsImage     `json:"IpfsImage"`
	IpfsFile                          []IpfsFile      `json:"IpfsFile"`
	FarmImage                         []FarmImage     `json:"FarmImage"`
	FarmFile                          []FarmFile      `json:"FarmFile"`
	UserId                            string          `json:"UserId"`
	Headers                           *Header         `json:"Headers"`
	AliasOrgName                      string          `json:"AliasOrgName"`
	UserName                          string          `json:"UserName"`
	SoilType                          string          `json:"SoilType"`
	IrrigationMethod                  string          `json:"IrrigationMethod"`
	AdditionalInformation             string          `json:"AdditionalInformation"`
	CropsGrown                        string          `json:"CropsGrown"`
	FarmingMethods                    string          `json:"FarmingMethods"`
	PestAndDiseaseManagementPractices string          `json:"PestAndDiseaseManagementPractices"`
	YearsOfFarmingExperience          string          `json:"YearsOfFarmingExperience"`
	Unit                              string          `json:"Unit"`
	FarmLogo                          []FarmImage     `json:"FarmLogo"`
}

type Image struct {
	ImageName    string `json:"ImageName"`
	ImagePath    string `json:"ImagePath"`
	ImageOrgName string `json:"ImageOrgName"`
	ImageSize    string `json:"ImageSize"`
	ImageDate    string `json:"ImageDate"`
}

type File struct {
	FileName    string `json:"FileName"`
	FilePath    string `json:"FilePath"`
	FileOrgName string `json:"FileOrgName"`
	FileSize    string `json:"FileSize"`
	FileDate    string `json:"FileDate"`
}

type CropDetails struct {
	FTLCID         string `json:"FTLCID"`
	CropID         string `json:"CropID"`
	FarmBy         string `json:"FarmBy"`
	FarmName       string `json:"FarmName"`
	CropType       string `json:"CropType"`
	CropName       string `json:"CropName"`
	PlantingDate   string `json:"PlantingDate"`
	HarvestingDate string `json:"HarvestingDate"`
	// PesticidesUsed  []string        `json:"PesticidesUsed"`
	SoilPh               string          `json:"SoilPh"`
	CropCondition        string          `json:"CropCondition"`
	Certification        []string        `json:"Certification"`
	FarmFile             []FarmFile      `json:"FarmFile"`
	BlockchainInfos      *BlockchainInfo `json:"BlockchainInfos"`
	IsDelete             int             `json:"IsDelete"`
	DocType              string          `json:"DocType"`
	UserId               string          `json:"UserId"`
	Headers              *Header         `json:"Headers"`
	AliasOrgName         string          `json:"AliasOrgName"`
	UserName             string          `json:"UserName"`
	CropImage            []Image         `json:"CropImage"`
	Status               string          `json:"Status"`
	Variety              string          `json:"Variety"`
	AreaUnderCultivation string          `json:"AreaUnderCultivation"`
	SoilType             string          `json:"SoilType"`
	TillingMethods       string          `json:"TillingMethods"`
	FertilisersUsed      string          `json:"FertilisersUsed"`
	Unit                 string          `json:"Unit"`
}

// type FertilizerEvent struct {
// 	FTLCID            string          `json:"FTLCID"`
// 	CropID            string          `json:"CropID"`
// 	EventID           string          `json:"EventID"`
// 	EventType         string          `json:"EventType"`
// 	EventDate         string          `json:"EventDate"`
// 	Details           string          `json:"Details"`
// 	ResponsibleParty  string          `json:"ResponsibleParty"`
// 	FarmName          string          `json:"FarmName"`
// 	QuantityUsed      int             `json:"QuantityUsed"`
// 	Unit              string          `json:"Unit"`
// 	ApplicationMethod string          `json:"ApplicationMethod"`
// 	BlockchainInfos   *BlockchainInfo `json:"BlockchainInfos"`
// 	UserId            string          `json:"UserId"`
// 	Headers           *Header         `json:"Headers"`
// 	AliasOrgName      string          `json:"AliasOrgName"`
// 	UserName          string          `json:"UserName"`
// 	FertilizerImage   []Image         `json:"FertilizerImage"`
// 	FertilizerName    string          `json:"FertilizerName"`
// 	CropName          string          `json:"CropName"`
// }

type FertilizerEvent struct {
	FTLCID            string `json:"FTLCID"`
	CropID            string `json:"CropID"`
	EventID           string `json:"EventID"`
	EventType         string `json:"EventType"`
	DateOfApplication string `json:"DateOfApplication"`
	// FertilizerUsed      string          `json:"FertilizerUsed"`
	ApplicationRate     string          `json:"ApplicationRate"`
	Unit                string          `json:"Unit"`
	MethodOfApplication string          `json:"MethodOfApplication"`
	GrowthStage         string          `json:"GrowthStage"`
	Notes               string          `json:"Notes"`
	TotalApplications   string          `json:"TotalApplications"`
	Location            string          `json:"Location"`
	BlockchainInfos     *BlockchainInfo `json:"BlockchainInfos"`
	UserId              string          `json:"UserId"`
	Headers             *Header         `json:"Headers"`
	AliasOrgName        string          `json:"AliasOrgName"`
	UserName            string          `json:"UserName"`
	FertilizerImage     []Image         `json:"FertilizerImage"`
	FertilizerName      string          `json:"FertilizerName"`
	CropName            string          `json:"CropName"`
}

//	type IrrigationEvent struct {
//		FTLCID           string          `json:"FTLCID"`
//		CropID           string          `json:"CropID"`
//		EventID          string          `json:"EventID"`
//		EventType        string          `json:"EventType"`
//		EventDate        string          `json:"EventDate"`
//		Details          string          `json:"Details"`
//		ResponsibleParty string          `json:"ResponsibleParty"`
//		FarmName         string          `json:"FarmName"`
//		QuantityUsed     int             `json:"QuantityUsed"`
//		Unit             string          `json:"Unit"`
//		WaterSource      string          `json:"WaterSource"`
//		IrrigationMethod string          `json:"IrrigationMethod"`
//		BlockchainInfos  *BlockchainInfo `json:"BlockchainInfos"`
//		UserId           string          `json:"UserId"`
//		Headers          *Header         `json:"Headers"`
//		AliasOrgName     string          `json:"AliasOrgName"`
//		UserName         string          `json:"UserName"`
//		IrrigationImage  []Image         `json:"IrrigationImage"`
//		CropName         string          `json:"CropName"`
//	}
type IrrigationEvent struct {
	FTLCID             string          `json:"FTLCID"`
	CropID             string          `json:"CropID"`
	EventID            string          `json:"EventID"`
	EventType          string          `json:"EventType"`
	DateOfIrrigation   string          `json:"DateOfIrrigation"`
	TimeOfIrrigation   string          `json:"TimeOfIrrigation"`
	FieldId            string          `json:"FieldId"`
	SourceOfWater      string          `json:"SourceOfWater"`
	StageOfCrop        string          `json:"StageOfCrop"`
	Comments           string          `json:"Comments"`
	MethodOfIrrigation string          `json:"MethodOfIrrigation"`
	BlockchainInfos    *BlockchainInfo `json:"BlockchainInfos"`
	UserId             string          `json:"UserId"`
	Headers            *Header         `json:"Headers"`
	AliasOrgName       string          `json:"AliasOrgName"`
	UserName           string          `json:"UserName"`
	IrrigationImage    []Image         `json:"IrrigationImage"`
	CropName           string          `json:"CropName"`
}

type QualityAssessment struct {
	Size             string `json:"Size"`
	Color            string `json:"Color"`
	OverallCondition string `json:"OverallCondition"`
}

// type HarvestingEvent struct {
// 	FTLCID            string             `json:"FTLCID"`
// 	FarmID            string             `json:"FarmID"`
// 	CropID            string             `json:"CropID"`
// 	EventID           string             `json:"EventID"`
// 	EventType         string             `json:"EventType"`
// 	EventDate         string             `json:"EventDate"`
// 	QuantityHarvested int                `json:"QuantityHarvested"`
// 	HarvestedBy       string             `json:"HarvestedBy"`
// 	FarmName          string             `json:"FarmName"`
// 	HarvestConditions string             `json:"HarvestConditions"`
// 	StorageConditions string             `json:"StorageConditions"`
// 	QualityAssessment *QualityAssessment `json:"QualityAssessment"`
// 	BlockchainInfos   *BlockchainInfo    `json:"BlockchainInfos"`
// 	UserId            string             `json:"UserId"`
// 	Headers           *Header            `json:"Headers"`
// 	AliasOrgName      string             `json:"AliasOrgName"`
// 	UserName          string             `json:"UserName"`
// 	HarvestingImage   []Image            `json:"HarvestingImage"`
// 	CropName          string             `json:"CropName"`
// }

type HarvestingEvent struct {
	FTLCID            string          `json:"FTLCID"`
	HarvestingID      string          `json:"HarvestingID"`
	FarmID            string          `json:"FarmID"`
	CropID            string          `json:"CropID"`
	EventID           string          `json:"EventID"`
	EventType         string          `json:"EventType"`
	EventDate         string          `json:"EventDate"`
	QuantityHarvested int             `json:"QuantityHarvested"`
	HarvestedBy       string          `json:"HarvestedBy"`
	FarmName          string          `json:"FarmName"`
	WeatherConditions string          `json:"WeatherConditions"`
	StorageConditions string          `json:"StorageConditions"`
	CropCondition     string          `json:"CropCondition"`
	BlockchainInfos   *BlockchainInfo `json:"BlockchainInfos"`
	UserId            string          `json:"UserId"`
	Headers           *Header         `json:"Headers"`
	AliasOrgName      string          `json:"AliasOrgName"`
	UserName          string          `json:"UserName"`
	HarvestingImage   []Image         `json:"HarvestingImage"`
	CropName          string          `json:"CropName"`
	AreaHarvested     string          `json:"AreaHarvested"`
	HarvestMethod     string          `json:"HarvestMethod"`
	Unit              string          `json:"Unit"`
}

//	type NutritionalContent struct {
//		VitaminC string `json:"VitaminC"`
//		Iron     string `json:"Iron"`
//		Calcium  string `json:"Calcium"`
//	}
type TestResults struct {
	PesticideResidue string `json:"PesticideResidue"`
	// NutritionalContent     *NutritionalContent `json:"NutritionalContent"`
	// NutritionalContent     string `json:"NutritionalContent"`
	MicrobialContamination string `json:"MicrobialContamination"`
	AllergenPresence       string `json:"AllergenPresence"`
}
type LabTestingEvent struct {
	FTLCID          string          `json:"FTLCID"`
	FarmID          string          `json:"FarmID"`
	FarmName        string          `json:"FarmName"`
	CropID          string          `json:"CropID"`
	EventID         string          `json:"EventID"`
	EventType       string          `json:"EventType"`
	EventDate       string          `json:"EventDate"`
	ResultFile      []File          `json:"ResultFile"`
	TestedBy        string          `json:"TestedBy"`
	TestResults     *TestResults    `json:"TestResults"`
	TypeOfSample    string          `json:"TypeOfSample"`
	BlockchainInfos *BlockchainInfo `json:"BlockchainInfos"`
	UserId          string          `json:"UserId"`
	Headers         *Header         `json:"Headers"`
	AliasOrgName    string          `json:"AliasOrgName"`
	UserName        string          `json:"UserName"`
	LabTestingImage []Image         `json:"LabTestingImage"`
	CropName        string          `json:"CropName"`
}
type PesticideEvent struct {
	FTLCID              string          `json:"FTLCID"`
	CropID              string          `json:"CropID"`
	EventID             string          `json:"EventID"`
	EventType           string          `json:"EventType"`
	DateOfApplication   string          `json:"DateOfApplication"`
	PesticideUsed       string          `json:"PesticideUsed"`
	ApplicationRate     string          `json:"ApplicationRate"`
	Unit                string          `json:"Unit"`
	MethodOfApplication string          `json:"MethodOfApplication"`
	BlockchainInfos     *BlockchainInfo `json:"BlockchainInfos"`
	UserId              string          `json:"UserId"`
	Headers             *Header         `json:"Headers"`
	AliasOrgName        string          `json:"AliasOrgName"`
	UserName            string          `json:"UserName"`
	PesticideImage      []Image         `json:"PesticideImage"`
	CropId              string          `json:"CropId"`
	CropName            string          `json:"CropName"`
	TargetPest          string          `json:"TargetPest"`
	GrowthStage         string          `json:"GrowthStage"`
	Notes               string          `json:"Notes"`
	TotalApplications   string          `json:"TotalApplications"`
}
type MiscellaneousEvents struct {
	DateOfEvent       string `json:"DateOfEvent"`
	StageOfGrowth     string `json:"StageOfGrowth"`
	EventType         string `json:"EventType"`
	DescriptionIssues string `json:"DescriptionIssues"`
	AffectedCrop      string `json:"AffectedCrop"`
	ActionsTaken      string `json:"ActionsTaken"`
}
type ObservationsEvent struct {
	FTLCID              string               `json:"FTLCID"`
	CropID              string               `json:"CropID"`
	CropName            string               `json:"CropName"`
	EventID             string               `json:"EventID"`
	EventType           string               `json:"EventType"`
	StageOfGrowth       string               `json:"StageOfGrowth"`
	DateOfObservation   string               `json:"DateOfObservation"`
	CropCondition       string               `json:"CropCondition"`
	DescriptionOfIssues string               `json:"DescriptionOfIssues"`
	ActionTaken         string               `json:"ActionTaken"`
	ObservationImage    []Image              `json:"ObservationImage"`
	MiscellaneousEvents *MiscellaneousEvents `json:"MiscellaneousEvents"`
	BlockchainInfos     *BlockchainInfo      `json:"BlockchainInfos"`
	Headers             *Header              `json:"Headers"`
	UserId              string               `json:"UserId"`
	AliasOrgName        string               `json:"AliasOrgName"`
	UserName            string               `json:"UserName"`
}
type Distribution struct {
	Distributor      *Distributor `json:"Distributor"`
	Destination      string       `json:"Destination"`
	DistributionDate string       `json:"DistributionDate"`
	DeliveryStatus   string       `json:"DeliveryStatus"`
}
type Distributor struct {
	ParticipantID   string `json:"ParticipantID"`
	DistributorName string `json:"DistributorName"`
	Location        string `json:"Location"`
}
type Participants struct {
	FarmID   []string `json:"FarmID"`
	FarmName string   `json:"FarmName"`
	CropID   []string `json:"CropID"`
	CropName string   `json:"CropName"`
}
type ProductDetail struct {
	FTLCID          string          `json:"FTLCID"`
	ProductID       string          `json:"ProductID"`
	ProductType     string          `json:"ProductType"`
	ProductName     string          `json:"ProductName"`
	BatchNumber     string          `json:"BatchNumber"`
	Quantity        string          `json:"Quantity"`
	Unit            string          `json:"Unit"`
	Distribution    *Distribution   `json:"Distribution"`
	Participants    *Participants   `json:"Participants"`
	BlockchainInfos *BlockchainInfo `json:"BlockchainInfos"`
	UserId          string          `json:"UserId"`
	UserName        string          `json:"UserName"`
}
type QuantityHarvested struct {
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
}
type HarvestWorkers struct {
	Name string `json:"Name"`
	Role string `json:"Role"`
}
type HarvestLocation struct {
	Latitude    string `json:"Latitude"`
	Longitude   string `json:"Longitude"`
	Description string `json:"Description"`
}

type Temperature struct {
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
}
type Humidity struct {
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
}
type Amount struct {
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
}
type Precipitation struct {
	Type   string  `json:"Type"`
	Amount *Amount `json:"Amount"`
}
type WeatherConditions struct {
	Temperature   *Temperature   `json:"Temperature"`
	Humidity      *Humidity      `json:"Humidity"`
	Precipitation *Precipitation `json:"Precipitation"`
}

type HarvestingKdes struct {
	FTLCID              string             `json:"FTLCID"`
	HarvestDate         string             `json:"HarvestDate"`
	CropType            string             `json:"CropType"`
	QuantityHarvested   *QuantityHarvested `json:"QuantityHarvested"`
	HarvestLocation     *HarvestLocation   `json:"HarvestLocation"`
	HarvestMethod       string             `json:"HarvestMethod"`
	HarvestWorkers      []HarvestWorkers   `json:"HarvestWorkers"`
	WeatherConditions   *WeatherConditions `json:"WeatherConditions"`
	Participants        *Participants      `json:"Participants"`
	Headers             *Header            `json:"Headers"`
	BlockchainInfos     *BlockchainInfo    `json:"BlockchainInfos"`
	DocType             string             `json:"DocType"`
	UserId              string             `json:"UserId"`
	AliasOrgName        string             `json:"AliasOrgName"`
	HarvestingKdesImage []Image            `json:"HarvestingKdesImage"`
	Status              string             `json:"Status"`
	HarvestingId        string             `json:"HarvestingId"`
}
type PackageItemInformation struct {
	BatchNumber  string `json:"BatchNumber"`
	Item         string `json:"Item"`
	Unit         string `json:"Unit"`
	CropFtlcId   string `json:"CropFtlcId"`
	Quantity     string `json:"Quantity"`
	HarvestingId string `json:"HarvestingId"`
	Weight       string `json:"Weight"`
}
type ShippingItemInformation struct {
	InitialPackageId string            `json:"InitialPackageId"`
	ItemInformation  []ItemInformation `json:"ItemInformation"`
}
type ItemInformation struct {
	BatchNumber string `json:"BatchNumber"`
	Item        string `json:"Item"`
	Unit        string `json:"Unit"`
	CropFtlcId  string `json:"CropFtlcId"`
	Quantity    string `json:"Quantity"`
}
type InitialPackagingKdes struct {
	FTLCID                 string `json:"FTLCID"`
	PackageIndentification string `json:"PackageIndentification"`
	Name                   string `json:"Name"`
	BatchNumber            string `json:"BatchNumber"`
	Description            string `json:"Description"`
	PackedDate             string `json:"PackedDate"`
	PackageMaterial        string `json:"PackageMaterial"`
	PackageMethod          string `json:"PackageMethod"`
	PackageByInformation   string `json:"PackageByInformation"`
	ShippingItemInformation
	// ItemInformation        []ItemInformation `json:"ItemInformation"`
	PackageItemInformation []PackageItemInformation `json:"PackageItemInformation"`
	Headers                *Header                  `json:"Headers"`
	BlockchainInfos        *BlockchainInfo          `json:"BlockchainInfos"`
	DocType                string                   `json:"DocType"`
	UserId                 string                   `json:"UserId"`
	AliasOrgName           string                   `json:"AliasOrgName"`
	InitialPackageImage    []Image                  `json:"InitialPackageImage"`
	Status                 string                   `json:"Status"`
}
type SenderInformation struct {
	ShipmentId          string `json:"ShipmentId"`
	BatchNumber         string `json:"BatchNumber"`
	Name                string `json:"Name"`
	Address             string `json:"Address"`
	ShipmentDate        string `json:"ShipmentDate"`
	PackageLatLan       string `json:"PackageLatLan"`
	PackageMaterial     string `json:"PackageMaterial"`
	PackageMethod       string `json:"PackageMethod"`
	PackedByInformation string `json:"PackedByInformation"`
}
type ReceiverInformation struct {
	ReceiverId      string `json:"ReceiverId"`
	ReceiverName    string `json:"ReceiverName"`
	ReceiverEmail   string `json:"ReceiverEmail"`
	ReceiverAddress string `json:"ReceiverAddress"`
}
type CarrierInformation struct {
	CompanyName    string `json:"CompanyName"`
	PhoneNumber    string `json:"PhoneNumber"`
	Email          string `json:"Email"`
	ContactPerson  string `json:"ContactPerson"`
	VechicalNumber string `json:"VechicalNumber"`
	VechicalType   string `json:"VechicalType"`
}
type ShippingingKdes struct {
	FTLCID            string             `json:"FTLCID"`
	SenderInformation *SenderInformation `json:"SenderInformation"`
	// ItemInformation     []ItemInformation    `json:"ItemInformation"`
	ShippingItemInformation []ShippingItemInformation `json:"ShippingItemInformation"`
	ReceiverInformation     *ReceiverInformation      `json:"ReceiverInformation"`
	CarrierInformation      *CarrierInformation       `json:"CarrierInformation"`
	Headers                 *Header                   `json:"Headers"`
	BlockchainInfos         *BlockchainInfo           `json:"BlockchainInfos"`
	DocType                 string                    `json:"DocType"`
	UserId                  string                    `json:"UserId"`
	Status                  string                    `json:"Status"`
	IsAccepted              string                    `json:"IsAccepted"`
	AliasOrgName            string                    `json:"AliasOrgName"`
	ShippingImage           []Image                   `json:"ShippingImage"`
}

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

//	type Ingredients struct {
//		BatchNumber    string    `json:"BatchNumber"`
//		Name           string    `json:"Name"`
//		Unit           string    `json:"Unit"`
//		CropFtlcId     string    `json:"CropFtlcId"`
//		Quantity       string    `json:"Quantity"`
//		RecevingFtlcId string    `json:"RecevingFtlcId"`
//		Supplier       *Supplier `json:"Supplier"`
//	}

type Ingredients struct {
	RecevingFtlcId  string            `json:"RecevingFtlcId"`
	ItemInformation []ItemInformation `json:"ItemInformation"`
	Supplier        *Supplier         `json:"Supplier"`
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
	TransformationImage  []Image               `json:"TransformationImage"`
	ProductImage         []Image               `json:"ProductImage"`
	Status               string                `json:"Status"`
}
type ProductItemInformation struct {
	BatchNumber          string `json:"BatchNumber"`
	ProductName          string `json:"ProductName"`
	Unit                 string `json:"Unit"`
	Quantity             string `json:"Quantity"`
	TransformationFtlcId string `json:"TransformationFtlcId"`
}

type ProcessorShippingingKdes struct {
	FTLCID                 string                  `json:"FTLCID"`
	SenderInformation      *SenderInformation      `json:"SenderInformation"`
	ProductInformation     *ProductItemInformation `json:"ProductInformation"`
	ItemInformation        []Ingredients           `json:"ItemInformation"`
	ReceiverInformation    *ReceiverInformation    `json:"ReceiverInformation"`
	CarrierInformation     *CarrierInformation     `json:"CarrierInformation"`
	Headers                *Header                 `json:"Headers"`
	BlockchainInfos        *BlockchainInfo         `json:"BlockchainInfos"`
	DocType                string                  `json:"DocType"`
	UserId                 string                  `json:"UserId"`
	Status                 string                  `json:"Status"`
	IsAccepted             string                  `json:"IsAccepted"`
	AliasOrgName           string                  `json:"AliasOrgName"`
	ProcessorShippingImage []Image                 `json:"ProcessorShippingImage"`
}
type DistributorShippingingKdes struct {
	FTLCID                   string                  `json:"FTLCID"`
	SenderInformation        *SenderInformation      `json:"SenderInformation"`
	ProductInformation       *ProductItemInformation `json:"ProductInformation"`
	ItemInformation          []Ingredients           `json:"ItemInformation"`
	ReceiverInformation      *ReceiverInformation    `json:"ReceiverInformation"`
	CarrierInformation       *CarrierInformation     `json:"CarrierInformation"`
	Headers                  *Header                 `json:"Headers"`
	BlockchainInfos          *BlockchainInfo         `json:"BlockchainInfos"`
	DocType                  string                  `json:"DocType"`
	UserId                   string                  `json:"UserId"`
	Status                   string                  `json:"Status"`
	IsAccepted               string                  `json:"IsAccepted"`
	AliasOrgName             string                  `json:"AliasOrgName"`
	DistributorShippingImage []Image                 `json:"DistributorShippingImage"`
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
	EventWhy      string `json:"eventWhy"`
	EventWhen     string `json:"eventWhen"`
	EventWhere    string `json:"eventWhere"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	UnixTimeStamp string `json:"UnixTimeStamp"`
	DeviceInfo    string `json:"deviceInfo"`
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

func (s *SmartContract) FoodTrazeCreate(ctx contractapi.TransactionContextInterface, status string, data1 string, data2 string, data3 string, data4 string, data5 string, data6 string, data7 string, data8 string, data9 string, data10 string, data11 string, data12 string, data13 string, data14 string, data15 string, data16 string, data17 string, data18 string, data19 string, data20 string, data21 string, data22 string, data23 string, data24 string, data25 string, data26 string, data27 string) (interface{}, error) {
	var response FoodTazeRes
	if status == "CropCreateEvent" {

		// PesticidesUsed := strings.Split(data5, ",")
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
		var image []Image
		if data16 != "" {
			if err1 := json.Unmarshal([]byte(data16), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := CropDetails{
			FTLCID:               data10,
			FarmBy:               data1,
			FarmName:             data14,
			CropID:               data2,
			CropType:             data3,
			CropName:             data8,
			PlantingDate:         data4,
			HarvestingDate:       data9,
			SoilPh:               data5,
			CropCondition:        data6,
			Certification:        arrCertificate,
			FarmFile:             ipfsCert,
			BlockchainInfos:      &blockChainInfo,
			IsDelete:             0,
			UserId:               data11,
			DocType:              "Crop",
			Headers:              &headerContent,
			AliasOrgName:         "Producer",
			UserName:             data15,
			CropImage:            image,
			Status:               "Created",
			Variety:              data17,
			AreaUnderCultivation: data18,
			SoilType:             data19,
			TillingMethods:       data20,
			FertilisersUsed:      data21,
			Unit:                 data22,
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

	if status == "FarmCreateEvent" {
		// Parse JSON data into Asset struct
		var farmerInformatioData FarmerContactInformation
		if err := json.Unmarshal([]byte(data2), &farmerInformatioData); err != nil {
			// fmt.Println("Error parsing JSON:", err)
			return nil, fmt.Errorf("the contact information error %v", err)
		}
		var farmerAddress Address
		if err := json.Unmarshal([]byte(data16), &farmerAddress); err != nil {
			// fmt.Println("Error parsing JSON:", err)
			return nil, fmt.Errorf("the contact information error %v", err)
		}
		// Parse JSON data into Asset struct
		farmerData := Farmer{
			FarmerName:         data3,
			DateOfBirth:        data17,
			Gender:             data18,
			ContactInformation: &farmerInformatioData,
			Address:            &farmerAddress,
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
		// var cultivationPracticeData CultivationPractices
		// if err1 := json.Unmarshal([]byte(data7), &cultivationPracticeData); err1 != nil {
		// 	// fmt.Println("Error parsing JSON1:", err1)
		// 	return nil, fmt.Errorf("the cultivation practice error %v", err1)
		// }
		arrCertificate := strings.Split(data8, ",")
		// var ipfsImg []IpfsImage
		// if data9 != "" {
		// 	if err1 := json.Unmarshal([]byte(data9), &ipfsImg); err1 != nil {
		// 		// fmt.Println("Error parsing JSON1:", err1)
		// 		return nil, fmt.Errorf("the ipfs image data error %v", err1)
		// 	}
		// }
		// var ipfsCert []IpfsFile
		// if data10 != "" {
		// 	if err1 := json.Unmarshal([]byte(data10), &ipfsCert); err1 != nil {
		// 		// fmt.Println("Error parsing JSON1:", err1)
		// 		return nil, fmt.Errorf("the ipfs image data error %v", err1)
		// 	}
		// }
		var ipfsImg []FarmImage
		if data9 != "" {
			if err1 := json.Unmarshal([]byte(data9), &ipfsImg); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var ipfsCert []FarmFile
		if data10 != "" {
			if err1 := json.Unmarshal([]byte(data10), &ipfsCert); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var farmLogo []FarmImage
		if data27 != "" {
			if err1 := json.Unmarshal([]byte(data27), &farmLogo); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		// ipfsCid := strings.Split(data9, ",")
		// // Parse JSON data into Asset struct
		// var blockChainInfo BlockchainInfo
		// if err1 := json.Unmarshal([]byte(data9), &blockChainInfo); err1 != nil {
		// 	// fmt.Println("Error parsing JSON:", err)
		// 	return nil, fmt.Errorf("the blockchain info error %v", err1)
		// }
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp

		var headerContent Header
		if err1 := json.Unmarshal([]byte(data14), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}

		asset := Farm{
			FTLCID:   data11,
			FarmID:   data1,
			FarmName: data13,
			Farmer:   &farmerData,
			Location: &locationData,
			FarmSize: data6,
			// CultivationPractices: &cultivationPracticeData,
			SoilPhLevel:     data7,
			Certifications:  arrCertificate,
			BlockchainInfos: &blockChainInfo,
			IsDelete:        0,
			DocType:         "Farm",
			// IpfsImage:            ipfsImg,
			// IpfsFile:             ipfsCert,
			FarmImage:                         ipfsImg,
			FarmFile:                          ipfsCert,
			UserId:                            data12,
			Headers:                           &headerContent,
			AliasOrgName:                      "Producer",
			UserName:                          data15,
			SoilType:                          data19,
			IrrigationMethod:                  data20,
			AdditionalInformation:             data21,
			CropsGrown:                        data22,
			FarmingMethods:                    data23,
			PestAndDiseaseManagementPractices: data24,
			YearsOfFarmingExperience:          data25,
			Unit:                              data26,
			FarmLogo:                          farmLogo,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
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
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data16), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var image []Image
		if data16 != "" {
			if err1 := json.Unmarshal([]byte(data17), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := FertilizerEvent{
			FTLCID:              data1,
			CropID:              data2,
			EventID:             data3,
			EventType:           data4,
			DateOfApplication:   data5,
			FertilizerName:      data6,
			ApplicationRate:     data7,
			Unit:                data8,
			MethodOfApplication: data9,
			GrowthStage:         data10,
			Notes:               data11,
			TotalApplications:   data12,
			Location:            data13,
			UserId:              data14,
			UserName:            data15,
			BlockchainInfos:     &blockChainInfo,
			Headers:             &headerContent,
			AliasOrgName:        "Producer",
			// FertilizerName:      data15,
			FertilizerImage: image,
			CropName:        data18,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)
		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Fertilizer Pesticide Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "PesticideEvent" {
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		// quantity, _ := strconv.Atoi(data7)
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data18), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var image []Image
		if data15 != "" {
			if err1 := json.Unmarshal([]byte(data15), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := PesticideEvent{
			FTLCID:              data1,
			EventID:             data2,
			EventType:           data3,
			CropID:              data4,
			CropName:            data5,
			DateOfApplication:   data6,
			PesticideUsed:       data7,
			ApplicationRate:     data8,
			Unit:                data9,
			MethodOfApplication: data10,
			TargetPest:          data11,
			GrowthStage:         data12,
			Notes:               data13,
			TotalApplications:   data14,
			UserId:              data16,
			UserName:            data17,
			PesticideImage:      image,
			BlockchainInfos:     &blockChainInfo,
			Headers:             &headerContent,
			AliasOrgName:        "Producer",
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)
		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Pesticide Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "IrrigationEvent" {
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		// quality, _ := strconv.Atoi(data7)
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data13), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var image []Image
		if data16 != "" {
			if err1 := json.Unmarshal([]byte(data15), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := IrrigationEvent{
			FTLCID:             data1,
			CropID:             data2,
			EventID:            data3,
			EventType:          data4,
			DateOfIrrigation:   data5,
			TimeOfIrrigation:   data6,
			FieldId:            data7,
			SourceOfWater:      data8,
			StageOfCrop:        data9,
			Comments:           data10,
			MethodOfIrrigation: data11,
			UserId:             data12,
			BlockchainInfos:    &blockChainInfo,
			Headers:            &headerContent,
			AliasOrgName:       "Producer",
			UserName:           data14,
			IrrigationImage:    image,
			CropName:           data16,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
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
		// var qualityAssessment QualityAssessment
		// if err1 := json.Unmarshal([]byte(data9), &qualityAssessment); err1 != nil {
		// 	// fmt.Println("Error parsing JSON1:", err1)
		// 	return nil, fmt.Errorf("the quality assessment error %v", err1)
		// }
		quantity, _ := strconv.Atoi(data5)

		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data12), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var image []Image
		if data15 != "" {
			if err1 := json.Unmarshal([]byte(data15), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := HarvestingEvent{
			FTLCID:            data10,
			CropID:            data1,
			EventID:           data2,
			EventType:         data3,
			EventDate:         data4,
			QuantityHarvested: quantity,
			HarvestedBy:       data6,
			FarmName:          data13,
			WeatherConditions: data7,
			StorageConditions: data8,
			// QualityAssessment: &qualityAssessment,
			CropCondition:   data9,
			UserId:          data11,
			BlockchainInfos: &blockChainInfo,
			Headers:         &headerContent,
			AliasOrgName:    "Producer",
			UserName:        data14,
			HarvestingImage: image,
			CropName:        data16,
			AreaHarvested:   data17,
			HarvestMethod:   data18,
			Unit:            data19,
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}
		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		// farmKey, err := ctx.GetStub().CreateCompositeKey("Farm", []string{data1})
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to create composite key: %v", err)
		// }

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs2 := []string{"Org1MSP"}
		err2 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs2)
		if err2 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err2)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Harvesting Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "LabTestingEvent" {
		// // Parse JSON data into Asset struct
		// var nutritionalContent NutritionalContent
		// if err1 := json.Unmarshal([]byte(data8), &nutritionalContent); err1 != nil {
		// 	// fmt.Println("Error parsing JSON1:", err1)
		// 	return nil, fmt.Errorf("the quality assessment error %v", err1)
		// }
		// Parse JSON data into Asset struct
		testResultData := TestResults{
			PesticideResidue: data8,
			// NutritionalContent:     data8,
			MicrobialContamination: data9,
			AllergenPresence:       data10,
		}
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data12), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var image []Image
		if data14 != "" {
			if err1 := json.Unmarshal([]byte(data14), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var file []File
		if data15 != "" {
			if err1 := json.Unmarshal([]byte(data15), &file); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := LabTestingEvent{
			FTLCID:          data1,
			CropID:          data2,
			EventID:         data3,
			EventType:       data4,
			EventDate:       data5,
			TypeOfSample:    data6,
			TestedBy:        data7,
			TestResults:     &testResultData,
			UserId:          data11,
			BlockchainInfos: &blockChainInfo,
			Headers:         &headerContent,
			AliasOrgName:    "Producer",
			UserName:        data13,
			LabTestingImage: image,
			ResultFile:      file,
			CropName:        data16,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Harvesting Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "ObservationsEvent" {
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		// quantity, _ := strconv.Atoi(data7)
		var miscellaneousEvents MiscellaneousEvents
		if err1 := json.Unmarshal([]byte(data11), &miscellaneousEvents); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data15), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var image []Image
		if data15 != "" {
			if err1 := json.Unmarshal([]byte(data12), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := ObservationsEvent{
			FTLCID:              data1,
			EventID:             data2,
			EventType:           data3,
			CropID:              data4,
			CropName:            data5,
			StageOfGrowth:       data6,
			DateOfObservation:   data7,
			CropCondition:       data8,
			DescriptionOfIssues: data9,
			ActionTaken:         data10,
			MiscellaneousEvents: &miscellaneousEvents,
			ObservationImage:    image,
			UserId:              data13,
			UserName:            data14,
			BlockchainInfos:     &blockChainInfo,
			Headers:             &headerContent,
			AliasOrgName:        "Producer",
		}
		assetJSON, err4 := json.Marshal(asset)
		if err4 != nil {
			return nil, fmt.Errorf("the asset json %s already exists", data2)
		}

		// result := ctx.GetStub().PutState(farmKey, assetJSON)
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)
		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Pesticide Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "ProductEvent" {
		// // Parse JSON data into Asset struct
		var distributorContent Distributor
		if err1 := json.Unmarshal([]byte(data8), &distributorContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error %v", err1)
		}
		var distribution Distribution
		distribution.DistributionDate = data7
		distribution.Distributor = &distributorContent
		distribution.Destination = data9
		distribution.DeliveryStatus = data10
		var participantContent Participants
		if err1 := json.Unmarshal([]byte(data11), &participantContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		asset := ProductDetail{
			FTLCID:          data12,
			ProductID:       data1,
			ProductType:     data2,
			ProductName:     data3,
			BatchNumber:     data4,
			Quantity:        data5,
			Unit:            data6,
			Distribution:    &distribution,
			Participants:    &participantContent,
			BlockchainInfos: &blockChainInfo,
			UserId:          data13,
			UserName:        data14,
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
		result := ctx.GetStub().PutState(data1, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, data1, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Harvesting Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "HarvestingKdesEvent" {
		// Parse JSON data into Asset struct
		var quantityHarvested QuantityHarvested
		if err1 := json.Unmarshal([]byte(data4), &quantityHarvested); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error %v", err1)
		}
		// Parse JSON data into Asset struct
		var harvestLocation HarvestLocation
		if err1 := json.Unmarshal([]byte(data5), &harvestLocation); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error %v", err1)
		}
		var harvestWorkers []HarvestWorkers
		if data9 != "" {
			if err1 := json.Unmarshal([]byte(data7), &harvestWorkers); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		// Parse JSON data into Asset struct
		var temperature Temperature
		if err1 := json.Unmarshal([]byte(data8), &temperature); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error %v", err1)
		}
		// Parse JSON data into Asset struct
		var humidity Humidity
		if err1 := json.Unmarshal([]byte(data9), &humidity); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error %v", err1)
		}
		// Parse JSON data into Asset struct
		var amount Amount
		if err1 := json.Unmarshal([]byte(data11), &amount); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error %v", err1)
		}
		var precipitation Precipitation
		precipitation.Type = data10
		precipitation.Amount = &amount

		var weatherConditions WeatherConditions
		weatherConditions.Temperature = &temperature
		weatherConditions.Humidity = &humidity
		weatherConditions.Precipitation = &precipitation

		var participantContent Participants
		if err1 := json.Unmarshal([]byte(data12), &participantContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}

		var headerContent Header
		if err1 := json.Unmarshal([]byte(data13), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		// headerContent.CreatedDate, _ = time.Parse("2006-01-02 15:04:05 ", headerContent.EventWhen)
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var image []Image
		if data15 != "" {
			if err1 := json.Unmarshal([]byte(data15), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := HarvestingKdes{
			FTLCID:              data1,
			HarvestDate:         data2,
			CropType:            data3,
			QuantityHarvested:   &quantityHarvested,
			HarvestLocation:     &harvestLocation,
			HarvestMethod:       data6,
			HarvestWorkers:      harvestWorkers,
			WeatherConditions:   &weatherConditions,
			Participants:        &participantContent,
			Headers:             &headerContent,
			BlockchainInfos:     &blockChainInfo,
			DocType:             "HarvestingKdes",
			UserId:              data14,
			AliasOrgName:        "Producer",
			HarvestingKdesImage: image,
			Status:              "Created",
			HarvestingId:        data16,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		// if participantContent.CropID != nil && len(participantContent.CropID) != 0 {
		// 	for _, cropId := range participantContent.CropID {
		// 		// Get crop details to update status consumed
		// 		itemInfoJSON, err := ctx.GetStub().GetState(cropId)
		// 		if err != nil {
		// 			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		// 		}
		// 		if itemInfoJSON == nil {
		// 			return false, fmt.Errorf("the farm %s does not exist", cropId)
		// 		}
		// 		var kdes CropDetails
		// 		err = json.Unmarshal(itemInfoJSON, &kdes)
		// 		if err != nil {
		// 			return false, fmt.Errorf("unmarshall farm data: %v", err)
		// 		}
		// 		kdes.Status = "Consumed"
		// 		assetJSON2, err4 := json.Marshal(kdes)
		// 		if err4 != nil {
		// 			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		// 		}
		// 		ctx.GetStub().PutState(cropId, assetJSON2)
		// 	}
		// }
		response = FoodTazeRes{
			Status:  200,
			Message: "Harvesting Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "InitialPackageKdesEvent" {
		// // Parse JSON data into Asset struct
		// var itemInformation []ItemInformation
		var itemInformation []PackageItemInformation
		if data10 != "" {
			if err1 := json.Unmarshal([]byte(data10), &itemInformation); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data11), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var image []Image
		if data13 != "" {
			if err1 := json.Unmarshal([]byte(data13), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := InitialPackagingKdes{
			FTLCID:                 data1,
			PackageIndentification: data2,
			Name:                   data3,
			BatchNumber:            data4,
			Description:            data5,
			PackedDate:             data6,
			PackageMaterial:        data7,
			PackageMethod:          data8,
			PackageByInformation:   data9,
			PackageItemInformation: itemInformation,
			Headers:                &headerContent,
			BlockchainInfos:        &blockChainInfo,
			UserId:                 data12,
			DocType:                "InitialPackageKdes",
			AliasOrgName:           "Producer",
			InitialPackageImage:    image,
			Status:                 "Created",
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		// if itemInformation != nil && len(itemInformation) != 0 {
		// 	for _, itemInfo := range itemInformation {
		// 		// Get crop details to update status consumed
		// 		itemInfoJSON, err := ctx.GetStub().GetState(itemInfo.HarvestingId)
		// 		if err != nil {
		// 			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		// 		}
		// 		if itemInfoJSON == nil {
		// 			return false, fmt.Errorf("the farm %s does not exist", itemInfo.HarvestingId)
		// 		}
		// 		var kdes HarvestingKdes
		// 		err = json.Unmarshal(itemInfoJSON, &kdes)
		// 		if err != nil {
		// 			return false, fmt.Errorf("unmarshall farm data: %v", err)
		// 		}
		// 		kdes.Status = "Consumed"
		// 		assetJSON2, err4 := json.Marshal(kdes)
		// 		if err4 != nil {
		// 			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		// 		}
		// 		ctx.GetStub().PutState(itemInfo.HarvestingId, assetJSON2)
		// 	}
		// }
		response = FoodTazeRes{
			Status:  200,
			Message: "Initial Package KDEs Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "ShippingKdesEvent" {
		// // Parse JSON data into Asset struct
		var senderContent SenderInformation
		if err1 := json.Unmarshal([]byte(data2), &senderContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var receiverContent ReceiverInformation
		if err1 := json.Unmarshal([]byte(data3), &receiverContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var itemInformation []ShippingItemInformation
		if data4 != "" {
			if err1 := json.Unmarshal([]byte(data4), &itemInformation); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var carrierContent CarrierInformation
		if err1 := json.Unmarshal([]byte(data5), &carrierContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data6), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		mspId, _ := ctx.GetClientIdentity().GetMSPID()
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		blockChainInfo.MspId = mspId
		// accept, _ := strconv.Atoi(data9)
		var image []Image
		if data10 != "" {
			if err1 := json.Unmarshal([]byte(data10), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := ShippingingKdes{
			FTLCID:            data1,
			SenderInformation: &senderContent,
			// ItemInformation:     itemInformation,
			ShippingItemInformation: itemInformation,
			ReceiverInformation:     &receiverContent,
			CarrierInformation:      &carrierContent,
			Headers:                 &headerContent,
			BlockchainInfos:         &blockChainInfo,
			UserId:                  data7,
			DocType:                 "ShippingKdes",
			Status:                  data8,
			IsAccepted:              data9,
			AliasOrgName:            "Producer",
			ShippingImage:           image,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org1MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		if itemInformation != nil && len(itemInformation) != 0 {
			for _, itemInfo := range itemInformation {
				// Get crop details to update status
				itemInfoJSON, err := ctx.GetStub().GetState(itemInfo.InitialPackageId)
				if err != nil {
					return false, fmt.Errorf("failed to read farm data from world state: %v", err)
				}
				if itemInfoJSON == nil {
					return false, fmt.Errorf("the farm %s does not exist", itemInfo.InitialPackageId)
				}
				var kdes InitialPackagingKdes
				err = json.Unmarshal(itemInfoJSON, &kdes)
				if err != nil {
					return false, fmt.Errorf("unmarshall farm data: %v", err)
				}
				kdes.Status = "Consumed"
				assetJSON2, err4 := json.Marshal(kdes)
				if err4 != nil {
					return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
				}
				ctx.GetStub().PutState(itemInfo.InitialPackageId, assetJSON2)
			}
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Shipping KDEs Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "TransformationKdesEvent" {
		// Parse JSON data into Asset struct
		var product TransformProduct
		if err1 := json.Unmarshal([]byte(data2), &product); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error1 %v", err1)
		}
		var ingredients []Ingredients
		// if data9 != "" {
		if err1 := json.Unmarshal([]byte(data3), &ingredients); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the ipfs image data error %v", err1)
		}
		// }
		// Parse JSON data into Asset struct
		var production Production
		if err1 := json.Unmarshal([]byte(data4), &production); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error2 %v", err1)
		}
		// Parse JSON data into Asset struct
		var batchLot BatchLot
		if err1 := json.Unmarshal([]byte(data5), &batchLot); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error3 %v", err1)
		}
		// Parse JSON data into Asset struct
		var qualityControlTests []QualityControlTests
		if err1 := json.Unmarshal([]byte(data6), &qualityControlTests); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the distributor error4 %v", err1)
		}

		var transformShipping TransformShipping
		if err1 := json.Unmarshal([]byte(data7), &transformShipping); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}

		var storageConditions StorageConditions
		if err1 := json.Unmarshal([]byte(data8), &storageConditions); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error5 %v", err1)
		}

		var recall Recall
		if err1 := json.Unmarshal([]byte(data9), &recall); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error6 %v", err1)
		}
		var regulatoryCertificate []string
		if err1 := json.Unmarshal([]byte(data10), &regulatoryCertificate); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error7 %v", err1)
		}
		var regulatoryCompliance RegulatoryCompliance
		regulatoryCompliance.Certifications = regulatoryCertificate
		regulatoryCompliance.ComplianceStatus = data11
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data12), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		var image []Image
		if data14 != "" {
			if err1 := json.Unmarshal([]byte(data14), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var prodImage []Image
		if data15 != "" {
			if err1 := json.Unmarshal([]byte(data15), &prodImage); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := TransformationKdes{
			FTLCID:           data1,
			TransformProduct: &product,
			// RecevingFtlcId:       data3,
			Ingredients:          ingredients,
			Production:           &production,
			BatchLot:             &batchLot,
			QualityControlTests:  qualityControlTests,
			TranformShipping:     &transformShipping,
			StorageConditions:    &storageConditions,
			Recall:               &recall,
			RegulatoryCompliance: &regulatoryCompliance,
			Headers:              &headerContent,
			BlockchainInfos:      &blockChainInfo,
			DocType:              "TransformationKdes",
			UserId:               data13,
			AliasOrgName:         "Processor",
			TransformationImage:  image,
			ProductImage:         prodImage,
			Status:               "Created",
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org2MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		if ingredients != nil && len(ingredients) != 0 {
			for _, itemInfo := range ingredients {
				// Get crop details to update status
				itemInfoJSON, err := ctx.GetStub().GetState(itemInfo.RecevingFtlcId)
				if err != nil {
					return false, fmt.Errorf("failed to read farm data from world state: %v", err)
				}
				if itemInfoJSON == nil {
					return false, fmt.Errorf("the farm %s does not exist", itemInfo.RecevingFtlcId)
				}
				var kdes ShippingingKdes
				err = json.Unmarshal(itemInfoJSON, &kdes)
				if err != nil {
					return false, fmt.Errorf("unmarshall farm data: %v", err)
				}
				kdes.Status = "Consumed"
				assetJSON2, err4 := json.Marshal(kdes)
				if err4 != nil {
					return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
				}
				ctx.GetStub().PutState(itemInfo.RecevingFtlcId, assetJSON2)
			}
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Traformation KDEs Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "ProcessorShippingKdesEvent" {
		// // Parse JSON data into Asset struct
		var senderContent SenderInformation
		if err1 := json.Unmarshal([]byte(data2), &senderContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var receiverContent ReceiverInformation
		if err1 := json.Unmarshal([]byte(data3), &receiverContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var productInformation ProductItemInformation
		if data4 != "" {
			if err1 := json.Unmarshal([]byte(data4), &productInformation); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var itemInformation []Ingredients
		// if data5 != "" {
		if err1 := json.Unmarshal([]byte(data5), &itemInformation); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the ipfs image data error %v", err1)
		}
		// }
		var carrierContent CarrierInformation
		if err1 := json.Unmarshal([]byte(data6), &carrierContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data7), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		mspId, _ := ctx.GetClientIdentity().GetMSPID()
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		blockChainInfo.MspId = mspId
		var image []Image
		if data11 != "" {
			if err1 := json.Unmarshal([]byte(data11), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		// accept, _ := strconv.Atoi(data9)
		asset := ProcessorShippingingKdes{
			FTLCID:                 data1,
			SenderInformation:      &senderContent,
			ProductInformation:     &productInformation,
			ItemInformation:        itemInformation,
			ReceiverInformation:    &receiverContent,
			CarrierInformation:     &carrierContent,
			Headers:                &headerContent,
			BlockchainInfos:        &blockChainInfo,
			UserId:                 data8,
			DocType:                "ProcessorShippingKdes",
			Status:                 data9,
			IsAccepted:             data10,
			AliasOrgName:           "Processor",
			ProcessorShippingImage: image,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org2MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		if itemInformation != nil && len(itemInformation) != 0 {
			for _, itemInfo := range itemInformation {
				// Get crop details to update status
				itemInfoJSON, err := ctx.GetStub().GetState(itemInfo.RecevingFtlcId)
				if err != nil {
					return false, fmt.Errorf("failed to read data from world state: %v", err)
				}
				if itemInfoJSON == nil {
					return false, fmt.Errorf("the %s does not exist", itemInfo.RecevingFtlcId)
				}
				var kdes TransformationKdes
				err = json.Unmarshal(itemInfoJSON, &kdes)
				if err != nil {
					return false, fmt.Errorf("unmarshall data: %v", err)
				}
				kdes.Status = "Consumed"
				assetJSON2, err4 := json.Marshal(kdes)
				if err4 != nil {
					return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
				}
				ctx.GetStub().PutState(itemInfo.RecevingFtlcId, assetJSON2)
			}
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Processor Shipping KDEs Event Created Successfully.",
			Data:    result,
		}
	}
	if status == "DistributorShippingKdesEvent" {
		// // Parse JSON data into Asset struct
		var senderContent SenderInformation
		if err1 := json.Unmarshal([]byte(data2), &senderContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var receiverContent ReceiverInformation
		if err1 := json.Unmarshal([]byte(data3), &receiverContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var productInformation ProductItemInformation
		if data4 != "" {
			if err1 := json.Unmarshal([]byte(data4), &productInformation); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		var itemInformation []Ingredients
		// if data5 != "" {
		if err1 := json.Unmarshal([]byte(data5), &itemInformation); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the ipfs image data error %v", err1)
		}
		// }
		var carrierContent CarrierInformation
		if err1 := json.Unmarshal([]byte(data6), &carrierContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		var headerContent Header
		if err1 := json.Unmarshal([]byte(data7), &headerContent); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return nil, fmt.Errorf("the participant error %v", err1)
		}
		mspId, _ := ctx.GetClientIdentity().GetMSPID()
		channelId := ctx.GetStub().GetChannelID()
		transactionId := ctx.GetStub().GetTxID()
		clientId, _ := ctx.GetClientIdentity().GetID()
		timestamp, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().UTC().Format("2006-01-02 15:04:05"))
		// Parse JSON data into Asset struct
		var blockChainInfo BlockchainInfo
		blockChainInfo.TransactionID = transactionId
		blockChainInfo.ClientId = clientId
		blockChainInfo.ChannelId = channelId
		blockChainInfo.Timestamp = timestamp
		blockChainInfo.MspId = mspId
		// accept, _ := strconv.Atoi(data9)
		var image []Image
		if data11 != "" {
			if err1 := json.Unmarshal([]byte(data11), &image); err1 != nil {
				// fmt.Println("Error parsing JSON1:", err1)
				return nil, fmt.Errorf("the ipfs image data error %v", err1)
			}
		}
		asset := DistributorShippingingKdes{
			FTLCID:                   data1,
			SenderInformation:        &senderContent,
			ProductInformation:       &productInformation,
			ItemInformation:          itemInformation,
			ReceiverInformation:      &receiverContent,
			CarrierInformation:       &carrierContent,
			Headers:                  &headerContent,
			BlockchainInfos:          &blockChainInfo,
			UserId:                   data8,
			DocType:                  "DistributorShippingKdes",
			Status:                   data9,
			IsAccepted:               data10,
			AliasOrgName:             "Distributor",
			DistributorShippingImage: image,
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
		result := ctx.GetStub().PutState(asset.FTLCID, assetJSON)

		// Changes the endorsement policy to the new owner org
		endorsingOrgs := []string{"Org4MSP"}
		err1 := setAssetStateBasedEndorsement(ctx, asset.FTLCID, endorsingOrgs)
		if err1 != nil {
			return "", fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		}
		response = FoodTazeRes{
			Status:  200,
			Message: "Processor Shipping KDEs Event Created Successfully.",
			Data:    result,
		}
	}
	return response, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetAllFarms(ctx contractapi.TransactionContextInterface) ([]map[string]interface{}, error) {

	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\"}}", "Farm")
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var farms []map[string]interface{}
	for resultsIterator.HasNext() {
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

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadFarmById(ctx contractapi.TransactionContextInterface, id string) (map[string]interface{}, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the farm %s does not exist", id)
	}

	// var data map[string]interface{}
	// err = json.Unmarshal(assetJSON, &data)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to unmarshal data: %s", err.Error())
	// }

	// value, ok := .(string)
	// if !ok {
	// 	return "", fmt.Errorf("invalid data format for key %s", id)
	// }
	// Define an empty interface{} to hold the unmarshalled JSON data
	var jsonData map[string]interface{}

	// Unmarshal the byte array into the empty interface
	err = json.Unmarshal(assetJSON, &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("failed to unmarshal data: %s", err.Error())
	}

	// Print the unmarshalled JSON data
	fmt.Println(jsonData)

	return jsonData, nil

	// var asset Farm
	// err = json.Unmarshal(assetJSON, &asset)
	// if err != nil {
	// 	return nil, fmt.Errorf("the farm unmarshall error %v", err)
	// }

	// return &asset, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetAllCropsList(ctx contractapi.TransactionContextInterface) ([]map[string]interface{}, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\"}}", "Crop")
	// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []map[string]interface{}
	// var assets []*CropDetails
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
		assets = append(assets, asset)
	}

	return assets, nil

}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadCropById(ctx contractapi.TransactionContextInterface, id string) (map[string]interface{}, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset map[string]interface{}
	// var asset CropDetails
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}

	return asset, nil
}

// GetAllFarms returns all assets found in world state
func (s *SmartContract) GetAllCropsByFarmId(ctx contractapi.TransactionContextInterface, farmId string) ([]map[string]interface{}, error) {
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
	var assets []map[string]interface{}
	// var assets []*CropDetails
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset map[string]interface{}
		// var asset CropDetails
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

// type TrazeDetail struct {
// 	Farm       map[string]interface{}
// 	Crop       map[string]interface{}
// 	Fertilizer map[string]interface{}
// 	Irrigation map[string]interface{}
// 	Harvesting map[string]interface{}
// 	LabTesting map[string]interface{}
// 	Product    map[string]interface{}
// }

// type HarvestTrazeDetail struct {
// 	Crop       map[string]interface{}
// 	Fertilizer map[string]interface{}
// 	Irrigation map[string]interface{}
// 	Harvesting map[string]interface{}
// 	LabTesting map[string]interface{}
// }

// GetAllHarvest returns all assets found in world state
func (s *SmartContract) GetAllHarvestByCropId(ctx contractapi.TransactionContextInterface, cropId string) (map[string]interface{}, error) {
	exists, err2 := s.AssetExists(ctx, cropId)
	if err2 != nil {
		return nil, fmt.Errorf("the Farm Data %s exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("the Farm Data %s not exists", cropId)
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
	var asset CropDetails
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Crop"] = asset

	// Started To check Type as Fertilization
	queryString1 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Fertilization")
	resultsIterator1, err := ctx.GetStub().GetQueryResult(queryString1)
	if err != nil {
		return nil, fmt.Errorf("the selector %s not exists", cropId)
	}
	defer resultsIterator1.Close()
	var count1 int
	var queryResponse1 *queryresult.KV
	for resultsIterator1.HasNext() {
		queryResponse1, err = resultsIterator1.Next()
		if err != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		count1++
	}
	var asset1 map[string]interface{}
	if count1 != 0 {

		err3 := json.Unmarshal(queryResponse1.Value, &asset1)
		if err3 != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		// Assign value for Fertilization
		assets["Fertilization"] = asset1
	} else {
		assets["Fertilization"] = asset1
	}

	// Started To check Type as Irrigation
	queryString2 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Irrigation")
	resultsIterator2, err := ctx.GetStub().GetQueryResult(queryString2)
	if err != nil {
		return nil, fmt.Errorf("the selector %s not exists", cropId)
	}
	defer resultsIterator2.Close()
	var count2 int
	var queryResponse2 *queryresult.KV
	for resultsIterator2.HasNext() {
		queryResponse2, err = resultsIterator2.Next()
		if err != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		count2++
	}
	var asset2 map[string]interface{}
	if count2 != 0 {

		err3 := json.Unmarshal(queryResponse2.Value, &asset2)
		if err3 != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		// Assign value for Irrigation
		assets["Irrigation"] = asset2
	} else {
		assets["Irrigation"] = asset2
	}

	// Started To check Type as Harvesting
	queryString3 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Harvesting")
	resultsIterator3, err := ctx.GetStub().GetQueryResult(queryString3)
	if err != nil {
		return nil, fmt.Errorf("the selector %s not exists", cropId)
	}
	defer resultsIterator3.Close()
	var count3 int
	var queryResponse3 *queryresult.KV
	for resultsIterator3.HasNext() {
		queryResponse3, err = resultsIterator3.Next()
		if err != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		count3++
	}
	var asset3 map[string]interface{}
	if count3 != 0 {
		// var asset3 map[string]interface{}
		err3 := json.Unmarshal(queryResponse3.Value, &asset3)
		if err3 != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		// Assign value for Harvesting
		assets["Harvesting"] = asset3
	} else {
		assets["Harvesting"] = asset3
	}

	// Started To check Type as LabTesting
	queryString4 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "LabTesting")
	resultsIterator4, err := ctx.GetStub().GetQueryResult(queryString4)
	if err != nil {
		return nil, fmt.Errorf("the selector %s not exists", cropId)
	}
	defer resultsIterator4.Close()
	var count4 int
	var queryResponse4 *queryresult.KV
	for resultsIterator4.HasNext() {
		queryResponse4, err = resultsIterator4.Next()
		if err != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		count4++
	}
	var asset4 map[string]interface{}
	if count4 != 0 {
		// var asset4 map[string]interface{}
		err3 := json.Unmarshal(queryResponse4.Value, &asset4)
		if err3 != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		// Assign value for LabTesting
		assets["LabTesting"] = asset4
	} else {
		assets["LabTesting"] = asset4
	}
	// Started To check Type as Pesticide
	queryString5 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Pesticide")
	resultsIterator5, err := ctx.GetStub().GetQueryResult(queryString5)
	if err != nil {
		return nil, fmt.Errorf("the selector %s not exists", cropId)
	}
	defer resultsIterator5.Close()
	var count5 int
	var queryResponse5 *queryresult.KV
	for resultsIterator5.HasNext() {
		queryResponse5, err = resultsIterator5.Next()
		if err != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		count5++
	}
	var asset5 map[string]interface{}
	if count5 != 0 {
		// var asset4 map[string]interface{}
		err3 := json.Unmarshal(queryResponse5.Value, &asset5)
		if err3 != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		// Assign value for LabTesting
		assets["Pesticide"] = asset5
	} else {
		assets["Pesticide"] = asset5
	}

	// Started To check Type as Observations
	queryString6 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Observations")
	resultsIterator6, err := ctx.GetStub().GetQueryResult(queryString6)
	if err != nil {
		return nil, fmt.Errorf("the selector %s not exists", cropId)
	}
	defer resultsIterator6.Close()
	var count6 int
	var queryResponse6 *queryresult.KV
	for resultsIterator6.HasNext() {
		queryResponse6, err = resultsIterator6.Next()
		if err != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		count6++
	}
	var asset6 map[string]interface{}
	if count6 != 0 {
		// var asset4 map[string]interface{}
		err3 := json.Unmarshal(queryResponse6.Value, &asset6)
		if err3 != nil {
			return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
		}
		// Assign value for LabTesting
		assets["Observations"] = asset6
	} else {
		assets["Observations"] = asset6
	}
	return assets, nil
}

// GetAllHarvest returns all assets found in world state
func (s *SmartContract) GetAllCropEventByCropId(ctx contractapi.TransactionContextInterface, cropId string) (map[string]interface{}, error) {
	exists, err2 := s.AssetExists(ctx, cropId)
	if err2 != nil {
		return nil, fmt.Errorf("The Crop Data %s exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("The Crop Data %s not exists", cropId)
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
	var asset CropDetails
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Crop"] = asset

	assetJSON1, err := ctx.GetStub().GetState(asset.FarmBy)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON1 == nil {
		return nil, fmt.Errorf("the farm %s does not exist", asset.FarmBy)
	}
	// var response FoodTazeRes
	// var data TrazeDetail
	var asset1 Farm
	err = json.Unmarshal(assetJSON1, &asset1)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Farm"] = asset1

	// Started To check Type as Fertilization
	// queryString1 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Fertilization")
	// resultsIterator1, err := ctx.GetStub().GetQueryResult(queryString1)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator1.Close()
	// var count1 int
	// var queryResponse1 *queryresult.KV
	// for resultsIterator1.HasNext() {
	// 	queryResponse1, err = resultsIterator1.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count1++
	// }
	// var asset1 map[string]interface{}
	// if count1 != 0 {

	// 	err3 := json.Unmarshal(queryResponse1.Value, &asset1)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for Fertilization
	// 	assets["Fertilization"] = asset1
	// } else {
	// 	assets["Fertilization"] = asset1
	// }

	var data []map[string]interface{}
	// Started To check Type as Fertilization
	filter := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"},\"sort\":[{\"Headers.eventWhen\":\"desc\"},{\"Headers.UnixTimeStamp\":\"desc\"}]}", cropId, "Fertilization")
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
	assets["Fertilization"] = data

	//Started To check Type as Irrigation
	// queryString2 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Irrigation")
	// resultsIterator2, err := ctx.GetStub().GetQueryResult(queryString2)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator2.Close()
	// var count2 int
	// var queryResponse2 *queryresult.KV
	// for resultsIterator2.HasNext() {
	// 	queryResponse2, err = resultsIterator2.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count2++
	// }
	// var asset2 map[string]interface{}
	// if count2 != 0 {

	// 	err3 := json.Unmarshal(queryResponse2.Value, &asset2)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for Irrigation
	// 	assets["Irrigation"] = asset2
	// } else {
	// 	assets["Irrigation"] = asset2
	// }

	var data2 []map[string]interface{}
	// Started To check Type as Fertilization
	filter2 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"},\"sort\":[{\"Headers.eventWhen\":\"desc\"},{\"Headers.UnixTimeStamp\":\"desc\"}]}", cropId, "Irrigation")
	resultsIterator2, err := ctx.GetStub().GetQueryResult(filter2)
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
	assets["Irrigation"] = data2

	// Started To check Type as Harvesting
	// queryString3 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Harvesting")
	// resultsIterator3, err := ctx.GetStub().GetQueryResult(queryString3)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator3.Close()
	// var count3 int
	// var queryResponse3 *queryresult.KV
	// for resultsIterator3.HasNext() {
	// 	queryResponse3, err = resultsIterator3.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count3++
	// }
	// var asset3 map[string]interface{}
	// if count3 != 0 {
	// 	// var asset3 map[string]interface{}
	// 	err3 := json.Unmarshal(queryResponse3.Value, &asset3)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for Harvesting
	// 	assets["Harvesting"] = asset3
	// } else {
	// 	assets["Harvesting"] = asset3
	// }

	var data3 []map[string]interface{}
	// Started To check Type as Fertilization
	filter3 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"},\"sort\":[{\"Headers.eventWhen\":\"desc\"},{\"Headers.UnixTimeStamp\":\"desc\"}]}", cropId, "Harvesting")
	resultsIterator3, err := ctx.GetStub().GetQueryResult(filter3)
	if err != nil {
		return nil, err
	}
	defer resultsIterator3.Close()

	// var assets []map[string]interface{}
	for resultsIterator3.HasNext() {
		queryResponse3, err := resultsIterator3.Next()
		if err != nil {
			return nil, err
		}

		var asset3 map[string]interface{}
		err = json.Unmarshal(queryResponse3.Value, &asset3)
		if err != nil {
			return nil, err
		}
		data3 = append(data3, asset3)
	}
	assets["Harvesting"] = data3

	// Started To check Type as LabTesting
	// queryString4 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "LabTesting")
	// resultsIterator4, err := ctx.GetStub().GetQueryResult(queryString4)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator4.Close()
	// var count4 int
	// var queryResponse4 *queryresult.KV
	// for resultsIterator4.HasNext() {
	// 	queryResponse4, err = resultsIterator4.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count4++
	// }
	// var asset4 map[string]interface{}
	// if count4 != 0 {
	// 	// var asset4 map[string]interface{}
	// 	err3 := json.Unmarshal(queryResponse4.Value, &asset4)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for LabTesting
	// 	assets["LabTesting"] = asset4
	// } else {
	// 	assets["LabTesting"] = asset4
	// }
	var data4 []map[string]interface{}
	// Started To check Type as Fertilization
	filter4 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"},\"sort\":[{\"Headers.eventWhen\":\"desc\"},{\"Headers.UnixTimeStamp\":\"desc\"}]}", cropId, "LabTesting")
	resultsIterator4, err := ctx.GetStub().GetQueryResult(filter4)
	if err != nil {
		return nil, err
	}
	defer resultsIterator4.Close()

	// var assets []map[string]interface{}
	for resultsIterator4.HasNext() {
		queryResponse4, err := resultsIterator4.Next()
		if err != nil {
			return nil, err
		}

		var asset4 map[string]interface{}
		err = json.Unmarshal(queryResponse4.Value, &asset4)
		if err != nil {
			return nil, err
		}
		data4 = append(data4, asset4)
	}
	assets["LabTesting"] = data4
	// Started To check Type as Pesticide
	// queryString5 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Pesticide")
	// resultsIterator5, err := ctx.GetStub().GetQueryResult(queryString5)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator5.Close()
	// var count5 int
	// var queryResponse5 *queryresult.KV
	// for resultsIterator5.HasNext() {
	// 	queryResponse5, err = resultsIterator5.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count5++
	// }
	// var asset5 map[string]interface{}
	// if count5 != 0 {
	// 	// var asset4 map[string]interface{}
	// 	err3 := json.Unmarshal(queryResponse5.Value, &asset5)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for LabTesting
	// 	assets["Pesticide"] = asset5
	// } else {
	// 	assets["Pesticide"] = asset5
	// }

	var data5 []map[string]interface{}
	// Started To check Type as Fertilization
	// filter5 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Pesticide")
	filter5 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"},\"sort\":[{\"Headers.eventWhen\":\"desc\"},{\"Headers.UnixTimeStamp\":\"desc\"}]}", cropId, "Pesticide")
	resultsIterator5, err := ctx.GetStub().GetQueryResult(filter5)
	if err != nil {
		return nil, err
	}
	defer resultsIterator5.Close()

	// var assets []map[string]interface{}
	for resultsIterator5.HasNext() {
		queryResponse5, err := resultsIterator5.Next()
		if err != nil {
			return nil, err
		}

		var asset5 map[string]interface{}
		err = json.Unmarshal(queryResponse5.Value, &asset5)
		if err != nil {
			return nil, err
		}
		data5 = append(data5, asset5)
	}
	assets["Pesticide"] = data5

	// Started To check Type as Observations
	// queryString6 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Observations")
	// resultsIterator6, err := ctx.GetStub().GetQueryResult(queryString6)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator6.Close()
	// var count6 int
	// var queryResponse6 *queryresult.KV
	// for resultsIterator6.HasNext() {
	// 	queryResponse6, err = resultsIterator6.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count6++
	// }
	// var asset6 map[string]interface{}
	// if count6 != 0 {
	// 	// var asset4 map[string]interface{}
	// 	err3 := json.Unmarshal(queryResponse6.Value, &asset6)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for LabTesting
	// 	assets["Observations"] = asset6
	// } else {
	// 	assets["Observations"] = asset6
	// }
	var data6 []map[string]interface{}
	// Started To check Type as Fertilization
	filter6 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"},\"sort\":[{\"Headers.eventWhen\":\"desc\"},{\"Headers.UnixTimeStamp\":\"desc\"}]}", cropId, "Observations")
	resultsIterator6, err := ctx.GetStub().GetQueryResult(filter6)
	if err != nil {
		return nil, err
	}
	defer resultsIterator6.Close()

	// var assets []map[string]interface{}
	for resultsIterator6.HasNext() {
		queryResponse6, err := resultsIterator6.Next()
		if err != nil {
			return nil, err
		}

		var asset6 map[string]interface{}
		err = json.Unmarshal(queryResponse6.Value, &asset6)
		if err != nil {
			return nil, err
		}
		data6 = append(data6, asset6)
	}
	assets["Observations"] = data6
	return assets, nil
}

// GetAllHarvest returns all assets found in world state
func (s *SmartContract) GetAllCropEventByCropIdSort(ctx contractapi.TransactionContextInterface, cropId string, filters string) (map[string]interface{}, error) {
	exists, err2 := s.AssetExists(ctx, cropId)
	if err2 != nil {
		return nil, fmt.Errorf("The Crop Data %s exist error", err2)
	}
	if !exists {
		return nil, fmt.Errorf("The Crop Data %s not exists", cropId)
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
	var asset CropDetails
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Crop"] = asset

	assetJSON1, err := ctx.GetStub().GetState(asset.FarmBy)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON1 == nil {
		return nil, fmt.Errorf("the farm %s does not exist", asset.FarmBy)
	}
	// var response FoodTazeRes
	// var data TrazeDetail
	var asset1 Farm
	err = json.Unmarshal(assetJSON1, &asset1)
	if err != nil {
		return nil, fmt.Errorf("the unmarshall error %s", err)
	}
	// Assign value for crop
	assets["Farm"] = asset1

	// Started To check Type as Fertilization
	// queryString1 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Fertilization")
	// resultsIterator1, err := ctx.GetStub().GetQueryResult(queryString1)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator1.Close()
	// var count1 int
	// var queryResponse1 *queryresult.KV
	// for resultsIterator1.HasNext() {
	// 	queryResponse1, err = resultsIterator1.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count1++
	// }
	// var asset1 map[string]interface{}
	// if count1 != 0 {

	// 	err3 := json.Unmarshal(queryResponse1.Value, &asset1)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for Fertilization
	// 	assets["Fertilization"] = asset1
	// } else {
	// 	assets["Fertilization"] = asset1
	// }

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

// GetAllCropAndEventByBatchId returns all assets found in world state
func (s *SmartContract) GetAllCropAndEventByBatchId(ctx contractapi.TransactionContextInterface, filters string) (map[string]interface{}, error) {
	// exists, err2 := s.AssetExists(ctx, cropId)
	// if err2 != nil {
	// 	return nil, fmt.Errorf("The Crop Data %s exist error", err2)
	// }
	// if !exists {
	// 	return nil, fmt.Errorf("The Crop Data %s not exists", cropId)
	// }

	assets := make(map[string]interface{})

	// assetJSON, err := ctx.GetStub().GetState(cropId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read from world state: %v", err)
	// }
	// if assetJSON == nil {
	// 	return nil, fmt.Errorf("the crop %s does not exist", cropId)
	// }
	// // var response FoodTazeRes
	// // var data TrazeDetail
	// var asset CropDetails
	// err = json.Unmarshal(assetJSON, &asset)
	// if err != nil {
	// 	return nil, fmt.Errorf("the unmarshall error %s", err)
	// }
	// // Assign value for crop
	// assets["Crop"] = asset

	// assetJSON1, err := ctx.GetStub().GetState(asset.FarmBy)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read from world state: %v", err)
	// }
	// if assetJSON1 == nil {
	// 	return nil, fmt.Errorf("the farm %s does not exist", asset.FarmBy)
	// }
	// // var response FoodTazeRes
	// // var data TrazeDetail
	// var asset1 Farm
	// err = json.Unmarshal(assetJSON1, &asset1)
	// if err != nil {
	// 	return nil, fmt.Errorf("the unmarshall error %s", err)
	// }
	// // Assign value for crop
	// assets["Farm"] = asset1

	// Started To check Type as Fertilization
	// queryString1 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Fertilization")
	// resultsIterator1, err := ctx.GetStub().GetQueryResult(queryString1)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator1.Close()
	// var count1 int
	// var queryResponse1 *queryresult.KV
	// for resultsIterator1.HasNext() {
	// 	queryResponse1, err = resultsIterator1.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count1++
	// }
	// var asset1 map[string]interface{}
	// if count1 != 0 {

	// 	err3 := json.Unmarshal(queryResponse1.Value, &asset1)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for Fertilization
	// 	assets["Fertilization"] = asset1
	// } else {
	// 	assets["Fertilization"] = asset1
	// }

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
	assets["PackageData"] = data
	return assets, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadTrazeById(ctx contractapi.TransactionContextInterface, id string, status string) (map[string]interface{}, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}
	// var response FoodTazeRes
	data := make(map[string]interface{})
	if status == "CropEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}

		data["Crop"] = asset
	}
	// return &data, nil
	// response = FoodTazeRes{
	// 	Status:  200,
	// 	Message: "Crop detail retrived Successfully.",
	// 	Data:    asset,queryString2 := fmt.Sprintf("{\"selector\":{\"CropID\":\"%s\",\"EventType\":\"%s\"}}", cropId, "Irrigation")
	// resultsIterator2, err := ctx.GetStub().GetQueryResult(queryString2)
	// if err != nil {
	// 	return nil, fmt.Errorf("the selector %s not exists", cropId)
	// }
	// defer resultsIterator2.Close()
	// var count2 int
	// var queryResponse2 *queryresult.KV
	// for resultsIterator2.HasNext() {
	// 	queryResponse2, err = resultsIterator2.Next()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	count2++
	// }
	// var asset2 map[string]interface{}
	// if count2 != 0 {

	// 	err3 := json.Unmarshal(queryResponse2.Value, &asset2)
	// 	if err3 != nil {
	// 		return nil, fmt.Errorf("the queryResponse %s not exists", cropId)
	// 	}
	// 	// Assign value for Irrigation
	// 	assets["Irrigation"] = asset2
	// } else {
	// 	assets["Irrigation"] = asset2
	// }

	if status == "FarmEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Farm"] = asset
	}

	if status == "FertilizerPesticideEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Fertilizer"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Fertilizer Pesticide detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "PesticideEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Pesticide"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Fertilizer Pesticide detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "IrrigationEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Irrigation"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Irrigation detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "HarvestingEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Harvesting"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Harvesting event retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "LabTestingEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["LabTesting"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Labtesting detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "ObservationsEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Observations"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Harvesting event retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "ProductEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["Product"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "HarvestingKdesEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["HarvestingKdes"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "InitialPackageKdesEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["InitialPackageKdes"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "ShippingKdesEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["ShippingKdes"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "TransformationKdesEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["TransformationKdes"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "ProcessorShippingKdesEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["ProcessorShippingKdes"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "DistributorShippingKdesEvent" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the unmarshall error %s", err)
		}
		data["DistributorShippingKdes"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	if status == "TrazeDetail" {
		var asset map[string]interface{}
		err = json.Unmarshal(assetJSON, &asset)
		if err != nil {
			return nil, fmt.Errorf("the product detail unmarshall error %s", err)
		}
		data["Data"] = asset
		// response = FoodTazeRes{
		// 	Status:  200,
		// 	Message: "Product detail retrived Successfully.",
		// 	Data:    asset,
		// }
	}
	return data, nil
	// return &response, nil
}

// type TrazeList struct {
// 	Farm       []map[string]interface{}
// 	Crop       []map[string]interface{}
// 	Fertilizer []map[string]interface{}
// 	Irrigation []map[string]interface{}
// 	Harvesting []map[string]interface{}
// 	LabTesting []map[string]interface{}
// 	Product    []map[string]interface{}
// }

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetAllTrazeByEvent(ctx contractapi.TransactionContextInterface, status string, userId string, filter string) ([]map[string]interface{}, error) {

	var data []map[string]interface{}
	if status == "CropEvent" {
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

	}

	if status == "FarmEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "Farm", userId)
		// queryString := fmt.Sprintf(`{"selector":{"FarmID":"%s"}}`, farmId)
		resultsIterator, err := ctx.GetStub().GetQueryResult(filter)
		if err != nil {
			return nil, fmt.Errorf("the querystring error %s", err)
		}
		defer resultsIterator.Close()

		// var Farm []map[string]interface{}
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return nil, fmt.Errorf("the json error %s", err)
			}

			var asset map[string]interface{}
			err = json.Unmarshal(queryResponse.Value, &asset)
			if err != nil {
				return nil, fmt.Errorf("the unmarshall error %s", err)
			}
			data = append(data, asset)
		}

	}

	if status == "FertilizerPesticideEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"EventType\":\"%s\",\"UserId\":\"%s\"}}", "Fertilization", userId)
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
	}
	if status == "PesticideEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"EventType\":\"%s\",\"UserId\":\"%s\"}}", "Fertilization", userId)
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
	}
	if status == "IrrigationEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"EventType\":\"%s\",\"UserId\":\"%s\"}}", "Irrigation", userId)
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
	}
	if status == "HarvestingEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"EventType\":\"%s\",\"UserId\":\"%s\"}}", "Harvesting", userId)
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

		// data["Harvesting"] = assets
	}
	if status == "LabTestingEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"EventType\":\"%s\",\"UserId\":\"%s\"}}", "LabTesting", userId)
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

		// data["LabTesting"] = assets
	}
	if status == "ObservationsEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"EventType\":\"%s\",\"UserId\":\"%s\"}}", "Fertilization", userId)
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
	}
	if status == "ProductEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"ProductType\":\"%s\",\"UserId\":\"%s\"}}", "Product", userId)
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

		// data["Product"] = assets
	}
	if status == "HarvestingKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "HarvestingKdes", userId)

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

		// data["Product"] = assets
	}
	if status == "InitialPackageKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "InitialPackageKdes", userId)
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

		// data["Product"] = assets
	}
	if status == "ShippingKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "ShippingKdes", userId)
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

		// data["Product"] = assets
	}
	if status == "ReceivingKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "ShippingKdes", userId)
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

		// data["Product"] = assets
	}
	if status == "TransformationKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "TransformationKdes", userId)
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

		// data["Product"] = assets
	}
	if status == "ProcessorShippingKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "ProcessorShippingKdes", userId)
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

		// data["Product"] = assets
	}
	if status == "DistributorShippingKdesEvent" {
		// queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\"}}", "ProcessorShippingKdes", userId)
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

		// data["Product"] = assets
	}
	return data, nil
	// return &response, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetAllShippingkdesByEventStatus(ctx contractapi.TransactionContextInterface, types string, userId string, status string) ([]map[string]interface{}, error) {

	var data []map[string]interface{}
	// if types == "ShippingKdes" {
	queryString := fmt.Sprintf("{\"selector\":{\"DocType\":\"%s\",\"UserId\":\"%s\",\"Status\":\"%s\"}}", types, userId, status)
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

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) DeleteTrazeImageAndFileById(ctx contractapi.TransactionContextInterface, id, status, types, data string) (bool, error) {
	exists, err2 := s.AssetExists(ctx, id)
	if err2 != nil {
		return false, fmt.Errorf("the traze data %s exist error", err2)
	}
	if !exists {
		return false, fmt.Errorf("the traze data %s not exists", id)
	}
	if status == "logo" && types == "FarmEvent" {
		var farmLogo []FarmImage
		if err1 := json.Unmarshal([]byte(data), &farmLogo); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var farm Farm
		err = json.Unmarshal(assetJSON, &farm)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farm.FarmLogo = farmLogo

		assetJSON2, err4 := json.Marshal(farm)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "FarmEvent" {
		var ipfsImage []FarmImage
		if err1 := json.Unmarshal([]byte(data), &ipfsImage); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var farm Farm
		err = json.Unmarshal(assetJSON, &farm)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farm.FarmImage = ipfsImage

		assetJSON2, err4 := json.Marshal(farm)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "file" && types == "FarmEvent" {
		var ipfsCert []FarmFile
		if err1 := json.Unmarshal([]byte(data), &ipfsCert); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var farm Farm
		err = json.Unmarshal(assetJSON, &farm)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farm.FarmFile = ipfsCert

		assetJSON2, err4 := json.Marshal(farm)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "CropEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var crop CropDetails
		err = json.Unmarshal(assetJSON, &crop)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		crop.CropImage = image

		assetJSON2, err4 := json.Marshal(crop)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "file" && types == "CropEvent" {
		var ipfsCert []FarmFile
		if err1 := json.Unmarshal([]byte(data), &ipfsCert); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var crop CropDetails
		err = json.Unmarshal(assetJSON, &crop)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		crop.FarmFile = ipfsCert

		assetJSON2, err4 := json.Marshal(crop)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "FertilizerPesticideEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var fertilizer FertilizerEvent
		err = json.Unmarshal(assetJSON, &fertilizer)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		fertilizer.FertilizerImage = image

		assetJSON2, err4 := json.Marshal(fertilizer)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "PesticideEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var Pesticide PesticideEvent
		err = json.Unmarshal(assetJSON, &Pesticide)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		Pesticide.PesticideImage = image

		assetJSON2, err4 := json.Marshal(Pesticide)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "IrrigationEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var irrigation IrrigationEvent
		err = json.Unmarshal(assetJSON, &irrigation)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		irrigation.IrrigationImage = image

		assetJSON2, err4 := json.Marshal(irrigation)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "HarvestingEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read harvest data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the harvest %s does not exist", id)
		}
		var harvest HarvestingEvent
		err = json.Unmarshal(assetJSON, &harvest)
		if err != nil {
			return false, fmt.Errorf("unmarshall harvest data: %v", err)
		}
		harvest.HarvestingImage = image

		assetJSON2, err4 := json.Marshal(harvest)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "LabTestingEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read lab data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the lab %s does not exist", id)
		}
		var lab LabTestingEvent
		err = json.Unmarshal(assetJSON, &lab)
		if err != nil {
			return false, fmt.Errorf("unmarshall lab data: %v", err)
		}
		lab.LabTestingImage = image

		assetJSON2, err4 := json.Marshal(lab)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "file" && types == "LabTestingEvent" {
		var file []File
		if err1 := json.Unmarshal([]byte(data), &file); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read lab data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the lab %s does not exist", id)
		}
		var lab LabTestingEvent
		err = json.Unmarshal(assetJSON, &lab)
		if err != nil {
			return false, fmt.Errorf("unmarshall lab data: %v", err)
		}
		lab.ResultFile = file

		assetJSON2, err4 := json.Marshal(lab)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "HarvestingKdesEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read harvesting data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the harvesting %s does not exist", id)
		}
		var harvesting HarvestingKdes
		err = json.Unmarshal(assetJSON, &harvesting)
		if err != nil {
			return false, fmt.Errorf("unmarshall harvesting data: %v", err)
		}
		harvesting.HarvestingKdesImage = image

		assetJSON2, err4 := json.Marshal(harvesting)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "InitialPackageKdesEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read initial data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the initial %s does not exist", id)
		}
		var initial InitialPackagingKdes
		err = json.Unmarshal(assetJSON, &initial)
		if err != nil {
			return false, fmt.Errorf("unmarshall initial data: %v", err)
		}
		initial.InitialPackageImage = image

		assetJSON2, err4 := json.Marshal(initial)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "ShippingKdesEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read shipping data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the shipping %s does not exist", id)
		}
		var shipping ShippingingKdes
		err = json.Unmarshal(assetJSON, &shipping)
		if err != nil {
			return false, fmt.Errorf("unmarshall shipping data: %v", err)
		}
		shipping.ShippingImage = image

		assetJSON2, err4 := json.Marshal(shipping)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "TransformationKdesEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read farm data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the farm %s does not exist", id)
		}
		var transform TransformationKdes
		err = json.Unmarshal(assetJSON, &transform)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		transform.TransformationImage = image

		assetJSON2, err4 := json.Marshal(transform)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "ProcessorShippingKdesEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read processor data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the processor %s does not exist", id)
		}
		var processor ProcessorShippingingKdes
		err = json.Unmarshal(assetJSON, &processor)
		if err != nil {
			return false, fmt.Errorf("unmarshall processor data: %v", err)
		}
		processor.ProcessorShippingImage = image

		assetJSON2, err4 := json.Marshal(processor)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	} else if status == "image" && types == "DistributorShippingKdesEvent" {
		var image []Image
		if err1 := json.Unmarshal([]byte(data), &image); err1 != nil {
			// fmt.Println("Error parsing JSON1:", err1)
			return false, fmt.Errorf("the ipfs image data error %v", err1)
		}
		assetJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return false, fmt.Errorf("failed to read distributor data from world state: %v", err)
		}
		if assetJSON == nil {
			return false, fmt.Errorf("the distributor %s does not exist", id)
		}
		var distributor DistributorShippingingKdes
		err = json.Unmarshal(assetJSON, &distributor)
		if err != nil {
			return false, fmt.Errorf("unmarshall distributor data: %v", err)
		}
		distributor.DistributorShippingImage = image
		assetJSON2, err4 := json.Marshal(distributor)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)
	}
	// Changes the endorsement policy to the new owner org
	endorsingOrgs := []string{"Org1MSP"}
	err1 := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
	if err1 != nil {
		return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	}

	return true, nil
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

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) TrazeKdesTransfer(ctx contractapi.TransactionContextInterface, id string, typeOrg string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the farm %s does not exist", id)
	}
	if typeOrg == "Producer" {
		var kdes ShippingingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.Status = "Transferred"
		kdes.AliasOrgName = "Processor"
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
	if typeOrg == "Processor" {
		var kdes ProcessorShippingingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.Status = "Transferred"
		kdes.AliasOrgName = "Distributor"
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
	if typeOrg == "Distributor" {
		var kdes DistributorShippingingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.Status = "Transferred"
		kdes.AliasOrgName = "Retailer"
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
func (s *SmartContract) TrazeKdesTransfer1(ctx contractapi.TransactionContextInterface, id string, value string, typeOrg string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the farm %s does not exist", id)
	}
	if typeOrg == "Producer" {
		var kdes ShippingingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.IsAccepted = value

		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// ctx.GetStub().SetEndorsementPolicy([]byte(newPolicy))
		// Changes the endorsement policy to the new owner org
		// endorsingOrgs := []string{"Org2MSP"}
		// err1 := setAssetStateBasedEndorsement(ctx, kdes.FTLCID, endorsingOrgs)
		// if err1 != nil {
		// 	return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		// }
	}
	if typeOrg == "Processor" {
		var kdes ProcessorShippingingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.IsAccepted = value

		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// ctx.GetStub().SetEndorsementPolicy([]byte(newPolicy))
		// Changes the endorsement policy to the new owner org
		// endorsingOrgs := []string{"Org4MSP"}
		// err1 := setAssetStateBasedEndorsement(ctx, kdes.FTLCID, endorsingOrgs)
		// if err1 != nil {
		// 	return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		// }
	}
	if typeOrg == "Distributor" {
		var kdes DistributorShippingingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.IsAccepted = value

		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// ctx.GetStub().SetEndorsementPolicy([]byte(newPolicy))
		// Changes the endorsement policy to the new owner org
		// endorsingOrgs := []string{"Org4MSP"}
		// err1 := setAssetStateBasedEndorsement(ctx, kdes.FTLCID, endorsingOrgs)
		// if err1 != nil {
		// 	return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		// }
	}
	return true, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) UpdateKdesStatus(ctx contractapi.TransactionContextInterface, id string, data string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the farm %s does not exist", id)
	}
	var kdes ShippingingKdes
	err = json.Unmarshal(assetJSON, &kdes)
	if err != nil {
		return false, fmt.Errorf("unmarshall farm data: %v", err)
	}
	kdes.Status = data

	assetJSON2, err4 := json.Marshal(kdes)
	if err4 != nil {
		return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
	}
	ctx.GetStub().PutState(id, assetJSON2)
	// ctx.GetStub().SetEndorsementPolicy([]byte(newPolicy))
	// Changes the endorsement policy to the new owner org
	// endorsingOrgs := []string{"Org2MSP"}
	// err1 := setAssetStateBasedEndorsement(ctx, id, endorsingOrgs)
	// if err1 != nil {
	// 	return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
	// }
	// }
	return true, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) UpdateConsumedStatus(ctx contractapi.TransactionContextInterface, id string, types string, data string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read farm data from world state: %v", err)
	}
	if assetJSON == nil {
		return false, fmt.Errorf("the farm %s does not exist", id)
	}
	if types == "HarvestingKdes" {
		var kdes HarvestingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.Status = data

		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// ctx.GetStub().SetEndorsementPolicy([]byte(newPolicy))
		// Changes the endorsement policy to the new owner org
		// endorsingOrgs := []string{"Org2MSP"}
		// err1 := setAssetStateBasedEndorsement(ctx, kdes.FTLCID, endorsingOrgs)
		// if err1 != nil {
		// 	return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		// }
	}
	if types == "InitialPackageKdes" {
		var kdes InitialPackagingKdes
		err = json.Unmarshal(assetJSON, &kdes)
		if err != nil {
			return false, fmt.Errorf("unmarshall farm data: %v", err)
		}
		kdes.Status = data

		assetJSON2, err4 := json.Marshal(kdes)
		if err4 != nil {
			return false, fmt.Errorf("the asset json %s already exists", assetJSON2)
		}
		ctx.GetStub().PutState(id, assetJSON2)

		// ctx.GetStub().SetEndorsementPolicy([]byte(newPolicy))
		// Changes the endorsement policy to the new owner org
		// endorsingOrgs := []string{"Org1MSP"}
		// err1 := setAssetStateBasedEndorsement(ctx, kdes.FTLCID, endorsingOrgs)
		// if err1 != nil {
		// 	return false, fmt.Errorf("failed setting state based endorsement for new owner: %v", err1)
		// }
	}
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
	if err != nil {
		return nil, err
	}
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

	var farms []*Farm
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return true, fmt.Errorf("failed to iterate farm: %v", err)
		}
		// fmt.log("queryResponse.Value", queryResponse.Value)
		var farm Farm
		err = json.Unmarshal(queryResponse.Value, &farm)
		if err != nil {
			return true, fmt.Errorf("unmarshall farm data: %v", err)
		}
		farms = append(farms, &farm)
	}
	var result bool
	if len(farms) != 0 {
		result = true
	} else {
		result = false
	}
	return result, nil
}

func (s *SmartContract) CreateCrop(ctx contractapi.TransactionContextInterface, data1 string, data2 string, data3 string, data4 string, data5 string, data6 string, data7 string, data8 string) (interface{}, error) {

	// Parse JSON data into Asset struct
	var blockChainInfo BlockchainInfo
	if err1 := json.Unmarshal([]byte(data8), &blockChainInfo); err1 != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, err1
	}
	// PesticidesUsed := strings.Split(data5, ",")
	arrCertificate := strings.Split(data7, ",")

	asset := CropDetails{
		FarmBy:       data1,
		CropID:       data2,
		CropType:     data3,
		PlantingDate: data4,
		// PesticidesUsed:  PesticidesUsed,
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
		return nil, fmt.Errorf("the farmData error %v", err)
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
		return nil, fmt.Errorf("the farmData error1 %v", err)
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
		return nil, fmt.Errorf("the farmData error2 %v", err1)
	}
	// // Parse JSON data into Asset struct
	var blockChainInfo BlockchainInfo
	if err1 := json.Unmarshal([]byte(blockchainInfo), &blockChainInfo); err1 != nil {
		// fmt.Println("Error parsing JSON:", err)
		return nil, fmt.Errorf("the farmData error3 %v", err1)
	}

	arrCertificate := strings.Split(certifications, ",")

	asset := Farm{
		FarmID:   farmId,
		Farmer:   &farmerData,
		Location: &locationData,
		FarmSize: farmSize,
		// CultivationPractices: &cultivationPracticeData,
		Certifications:  arrCertificate,
		BlockchainInfos: &blockChainInfo,
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
