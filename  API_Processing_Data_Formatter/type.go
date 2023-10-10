package dpfm_api_output_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type DeliveryDocument struct {
	DeliveryDocument              *int     `json:"DeliveryDocument"`
	Buyer                         *int     `json:"Buyer"`
	Seller                        *int     `json:"Seller"`
	ContractType                  string   `json:"ContractType"`
	OrderValidityStartDate        *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate          *string  `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate      *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate        *string  `json:"InvoiceScheduleEndDate"`
	IssuingLocationTimeZone       string   `json:"IssuingLocationTimeZone"`
	ReceivingLocationTimeZone     string   `json:"ReceivingLocationTimeZone"`
	DocumentDate                  *string  `json:"DocumentDate"`
	PlannedGoodsIssueDate         *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime         *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate       *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime       *string  `json:"PlannedGoodsReceiptTime"`
	BillingDocumentDate           *string  `json:"BillingDocumentDate"`
	CompleteDeliveryIsDefined     *bool    `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus         string   `json:"OverallDeliveryStatus"`
	CreationDate                  *string  `json:"CreationDate"`
	CreationTime                  *string  `json:"CreationTime"`
	IssuingBlockReason            *bool    `json:"IssuingBlockReason"`
	ReceivingBlockReason          *bool    `json:"ReceivingBlockReason"`
	GoodsIssueOrReceiptSlipNumber string   `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus           string   `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus       string   `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockReason      *bool    `json:"HeaderBillingBlockReason"`
	HeaderGrossWeight             *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight               *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit              string   `json:"HeaderWeightUnit"`
	Incoterms                     string   `json:"Incoterms"`
	IsExportImportDelivery        *bool    `json:"IsExportImportDelivery"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	IssuingPlantBusinessPartner   string   `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                  string   `json:"IssuingPlant"`
	ReceivingPlantBusinessPartner string   `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                string   `json:"ReceivingPlant"`
	DeliverToParty                *int     `json:"DeliverToParty"`
	DeliverFromParty              *int     `json:"DeliverFromParty"`
	TransactionCurrency           string   `json:"TransactionCurrency"`
	OverallDelivReltdBillgStatus  string   `json:"OverallDelivReltdBillgStatus"`
}
