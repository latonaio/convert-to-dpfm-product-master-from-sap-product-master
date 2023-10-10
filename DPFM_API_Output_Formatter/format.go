package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "convert-to-dpfm-product-master-from-sap-product-master/DPFM_API_Processing_Formatter"

)

func ConvertToGeneral(subfuncSDC *sub_func_complementer.SDC) *General {
	data := subfuncSDC.Message.General

	general := &General{
		Product:                       data.Product,
		ProductType:                   data.ProductType,
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		ProductGroup:                  data.ProductGroup,
		GrossWeight:                   data.GrossWeight,
		NetWeight:                     data.NetWeight,
		WeightUnit:                    data.WeightUnit,
		InternalCapacityQuantity:      data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		ProductStandardID:             data.ProductStandardID,
		IndustryStandardName:          data.IndustryStandardName,
		ItemCategory:                  data.ItemCategory,
		BarcodeType:                   data.BarcodeType,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return general
}

func ConvertToGeneralUpdates(generalUpdates *dpfm_api_processing_formatter.GeneralUpdates) *General {
	data := generalUpdates

	general := &General{
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		GrossWeight:                   data.GrossWeight,
		NetWeight:                     data.NetWeight,
		WeightUnit:                    data.WeightUnit,
		InternalCapacityQuantity:      data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		ProductStandardID:             data.ProductStandardID,
		IndustryStandardName:          data.IndustryStandardName,
		BarcodeType:                   data.BarcodeType,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return general
}

func ConvertToBusinessPartner(subfuncSDC *sub_func_complementer.SDC) *[]BusinessPartner {
	var businessPartner []BusinessPartner

	for _, data := range *subfuncSDC.Message.BusinessPartner {
		businessPartner = append(businessPartner, BusinessPartner{
			Product:                data.Product,
			BusinessPartner:        data.BusinessPartner,
			ValidityEndDate:        data.ValidityEndDate,
			ValidityStartDate:      data.ValidityStartDate,
			BusinessPartnerProduct: data.BusinessPartnerProduct,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}

	return &businessPartner
}

func ConvertToBusinessPartnerUpdates(businessPartnerUpdates *[]dpfm_api_processing_formatter.BusinessPartnerUpdates) *[]BusinessPartner {
	var businessPartner []BusinessPartner

	for _, data := range *businessPartnerUpdates {
		businessPartner = append(businessPartner, BusinessPartner{
			BusinessPartnerProduct: data.BusinessPartnerProduct,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}

	return &businessPartner
}

func ConvertToAllergent(subfuncSDC *sub_func_complementer.SDC) *[]Allergen {
	var allergen []Allergen

	for _, data := range *subfuncSDC.Message.Allergen {
		allergen = append(allergen, Allergen{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			Allergen:                                  data.Allergen,
			AllergenIsContained:                       data.AllergenIsContained,
		})
	}
	return &allergen
}

func ConvertToAllergenUpdates(allergenUpdates *[]dpfm_api_processing_formatter.AllergenUpdates) *[]Allergen {
	var allergen []Allergen

	for _, data := range *allergenUpdates {
		allergen = append(allergen, Allergen{
			Allergen:                                              data.Allergen,
			AllergenIsContained:                                   data.AllergenIsContained,
		})
	}

	return &allergen
}

func ConvertToNutritionalInfo(subfuncSDC *sub_func_complementer.SDC) *[]NutritionalInfo {
	var nutritionalInfo []NutritionalInfo

	for _, data := range *subfuncSDC.Message.NutritionalInfo {
		nutritionalInfo = append(nutritionalInfo, NutritionalInfo{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			Nutrient:                                  data.Nutrient,
			NutrientContent:                           data.NutrientContent,
			NutrientContentUnit:                       data.NutrientContentUnit,
		})
	}
	return &nutritionalInfo
}

func ConvertToNutritionalInfoUpdates(nutritionalInfoUpdates *[]dpfm_api_processing_formatter.NutritionalInfoUpdates) *[]NutritionalInfo {
	var nutritionalInfo []NutritionalInfo

	for _, data := range *nutritionalInfoUpdates {
		nutritionalInfo = append(nutritionalInfo, NutritionalInfo{
			Nutrient:                                              data.Nutrient,
			NutrientContent:                                       data.NutrientContent,
			NutrientContentUnit:                                   data.NutrientContentUnit,
		})
	}

	return &nutritionalInfo
}

func ConvertToCalories(subfuncSDC *sub_func_complementer.SDC) *[]Calories {
	var calories []Calories

	for _, data := range *subfuncSDC.Message.Calories {
		calories = append(calories, Calories{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			CaloryUnitQuantitiy:                       data.CaloryUnitQuantitiy,
			Calories:                                  data.Calories,
			CaloryUnit:                                data.CaloryUnit,
		})
	}
	return &calories
}

func ConvertToCaloriesUpdates(caloriesUpdates *[]dpfm_api_processing_formatter.Calories) *[]Calories {
	var calories []Calories

	for _, data := range *caloriesUpdates {
		calories = append(calories, Calories{
			CaloryUnitQuantitiy:                          data.CaloryUnitQuantitiy,
			Calories:                                     data.Calories,
			CaloryUnit:                                   data.CaloryUnit,
		})
	}

	return &calories
}


func ConvertToBPPlant(subfuncSDC *sub_func_complementer.SDC) *[]BPPlant {
	var bpPlant []BPPlant

	for _, data := range *subfuncSDC.Message.BPPlant {
		bpPlant = append(bpPlant, BPPlant{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			Plant:                                     data.Plant,
			AvailabilityCheckType:                     data.AvailabilityCheckType,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
			PlanningTimeFence:                         data.PlanningTimeFence,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
			MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
			IsBatchManagementRequired:                 data.IsBatchManagementRequired,
			BatchManagementPolicy:                     data.BatchManagementPolicy,
			InventoryUnit:                             data.InventoryUnit,
			ProfitCenter:                              data.ProfitCenter,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	return &bpPlant
}

func ConvertToBPPlantUpdates(bpPlantUpdates *[]dpfm_api_processing_formatter.BPPlantUpdates) *[]BPPlant {
	var bpPlant []BPPlant

	for _, data := range *bpPlantUpdates {
		bpPlant = append(bpPlant, BPPlant{
			AvailabilityCheckType:                     data.AvailabilityCheckType,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
			PlanningTimeFence:                         data.PlanningTimeFence,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
			MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
			IsBatchManagementRequired:                 data.IsBatchManagementRequired,
			BatchManagementPolicy:                     data.BatchManagementPolicy,
			InventoryUnit:                             data.InventoryUnit,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}

	return &bpPlant
}

func ConvertToProductDescByBP(subfuncSDC *sub_func_complementer.SDC) *[]ProductDescByBP {
	var productDescByBP []ProductDescByBP

	for _, data := range *subfuncSDC.Message.ProductDescByBP {
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}
	return &productDescByBP
}

func ConvertToProductDescByBPUpdates(productDescByBPUpdates *[]dpfm_api_processing_formatter.ProductDescByBPUpdates) *[]ProductDescByBP {
	var productDescByBP []ProductDescByBP

	for _, data := range *productDescByBPUpdates {
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}

	return &productDescByBP
}

func ConvertToStorageLocation(subfuncSDC *sub_func_complementer.SDC) *[]StorageLocation {
	var storageLocation []StorageLocation

	for _, data := range *subfuncSDC.Message.StorageLocation {
		storageLocation = append(storageLocation, StorageLocation{
			Product:              data.Product,
			BusinessPartner:      data.BusinessPartner,
			Plant:                data.Plant,
			StorageLocation:      data.StorageLocation,
			CreationDate:         data.CreationDate,
			InventoryBlockStatus: data.InventoryBlockStatus,
			IsMarkedForDeletion:  data.IsMarkedForDeletion,
		})
	}

	return &storageLocation
}

func ConvertToStorageLocationUpdates(storageLocationUpdates *[]dpfm_api_processing_formatter.StorageLocationUpdates) *[]StorageLocation {
	var storageLocation []StorageLocation

	for _, data := range *storageLocationUpdates {
		storageLocation = append(storageLocation, StorageLocation{
			InventoryBlockStatus: data.InventoryBlockStatus,
			IsMarkedForDeletion:  data.IsMarkedForDeletion,
		})
	}

	return &storageLocation
}

func ConvertToMRPArea(subfuncSDC *sub_func_complementer.SDC) *[]MRPArea {
	var mRPArea []MRPArea

	for _, data := range *subfuncSDC.Message.MRPArea {
		mRPArea = append(mRPArea, MRPArea{
			Product:                                  data.Product,
			BusinessPartner:                          data.BusinessPartner,
			Plant:                                    data.Plant,
			MRPArea:                                  data.MRPArea,
			StorageLocationForMRP:                    data.StorageLocationForMRP,
			MRPType:                                  data.MRPType,
			MRPController:                            data.MRPController,
			ReorderThresholdQuantity:                 data.ReorderThresholdQuantity,
			PlanningTimeFence:                        data.PlanningTimeFence,
			MRPPlanningCalendar:                      data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:            data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                           data.SafetyDuration,
			MaximumStockQuantityInBaseUnit:           data.MaximumStockQuantityInBaseUnit,
			MinumumDeliveryQuantityInBaseUnit:        data.MinumumDeliveryQuantityInBaseUnit,
			MinumumDeliveryLotSizeQuantityInBaseUnit: data.MinumumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	return &mRPArea
}

func ConvertToMRPAreaUpdates(mrpAreaUpdates *[]dpfm_api_processing_formatter.MRPAreaUpdates) *[]MRPArea {
	var mrpArea []MRPArea

	for _, data := range *mrpAreaUpdates {
		mrpArea = append(mrpArea, MRPArea{
			StorageLocationForMRP:                     data.StorageLocationForMRP,
			MRPType:                                   data.MRPType,
			MRPController:                             data.MRPController,
			ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
			PlanningTimeFence:                         data.PlanningTimeFence,
			MRPPlanningCalendar:                       data.MRPPlanningCalendar,
			SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
			SafetyDuration:                            data.SafetyDuration,
			MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
			MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
			MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
			StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
			DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
			MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
			MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
			DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
			StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}

	return &mrpArea
}

func ConvertToWorkScheduling(subfuncSDC *sub_func_complementer.SDC) *[]WorkScheduling {
	var workScheduling []WorkScheduling

	for _, data := range *subfuncSDC.Message.WorkScheduling {
		workScheduling = append(workScheduling, WorkScheduling{
			Product:                       data.Product,
			BusinessPartner:               data.BusinessPartner,
			Plant:                         data.Plant,
			ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
			ProductProcessingTime:         data.ProductProcessingTime,
			ProductionSupervisor:          data.ProductionSupervisor,
			ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
			ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
			PDTCompIsMarkedForBackflush:   data.PDTCompIsMarkedForBackflush,
			ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	return &workScheduling
}

func ConvertToWorkSchedulingUpdates(workSchedulingUpdates *[]dpfm_api_processing_formatter.WorkSchedulingUpdates) *[]WorkScheduling {
	var workScheduling []WorkScheduling

	for _, data := range *workSchedulingUpdates {
		workScheduling = append(workScheduling, WorkScheduling{
			ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
			ProductProcessingTime:         data.ProductProcessingTime,
			ProductionSupervisor:          data.ProductionSupervisor,
			ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
			ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
			PDTCompIsMarkedForBackflush:   data.PDTCompIsMarkedForBackflush,
			ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}

	return &workScheduling
}

func ConvertToAccounting(subfuncSDC *sub_func_complementer.SDC) *[]Accounting {
	var accounting []Accounting

	for _, data := range *subfuncSDC.Message.Accounting {
		accounting = append(accounting, Accounting{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Plant:               data.Plant,
			ValuationClass:      data.ValuationClass,
			CostingPolicy:       data.CostingPolicy,
			PriceUnitQty:        data.PriceUnitQty,
			StandardPrice:       data.StandardPrice,
			MovingAveragePrice:  data.MovingAveragePrice,
			PriceLastChangeDate: data.PriceLastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &accounting
}

func ConvertToAccountingUpdates(accountingUpdates *[]dpfm_api_processing_formatter.AccountingUpdates) *[]Accounting {
	var accounting []Accounting

	for _, data := range *accountingUpdates {
		accounting = append(accounting, Accounting{
			ValuationClass:      data.ValuationClass,
			CostingPolicy:       data.CostingPolicy,
			PriceUnitQty:        data.PriceUnitQty,
			StandardPrice:       data.StandardPrice,
			MovingAveragePrice:  data.MovingAveragePrice,
			PriceLastChangeDate: data.PriceLastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &accounting
}

func ConvertToTax(subfuncSDC *sub_func_complementer.SDC) *[]Tax {
	var tax []Tax

	for _, data := range *subfuncSDC.Message.Tax {
		tax = append(tax, Tax{
			Product:                  data.Product,
			Country:                  data.Country,
			ProductTaxCategory:       data.ProductTaxCategory,
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &tax
}

func ConvertToTaxUpdates(taxUpdates *[]dpfm_api_processing_formatter.TaxUpdates) *[]Tax {
	var tax []Tax

	for _, data := range *taxUpdates {
		tax = append(tax, Tax{
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &tax
}

func ConvertToProductDescription(subfuncSDC *sub_func_complementer.SDC) []ProductDescription {
	var productDescription []ProductDescription

	for _, data := range *subfuncSDC.Message.ProductDescription {
		productDescription = append(productDescription, ProductDescription{
			Product:            data.Product,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}

	return productDescription
}

func ConvertToProductDescriptionUpdates(productDescriptionUpdates *[]dpfm_api_processing_formatter.ProductDescriptionUpdates) *[]ProductDescription {
	var productDescription []ProductDescription

	for _, data := range *productDescriptionUpdates {
		productDescription = append(productDescription, ProductDescription{
			ProductDescription: data.ProductDescription,
		})
	}

	return &productDescription
}
