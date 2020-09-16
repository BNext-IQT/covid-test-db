package diagnostic

import (
    "time"

    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/BNext-IQT/covid-test-db/models/poc"
    "github.com/BNext-IQT/covid-test-db/models/company"
    "github.com/BNext-IQT/covid-test-db/models/diagnostic_type"
    "github.com/BNext-IQT/covid-test-db/models/diagnostic_performance"
    "github.com/BNext-IQT/covid-test-db/models/diagnostic_target_type"
    "github.com/BNext-IQT/covid-test-db/models/regulatory_approval_type"
    "github.com/BNext-IQT/covid-test-db/models/sample_type"
    "github.com/BNext-IQT/covid-test-db/models/pcr_platform"
)

type Diagnostic struct {
    Id      			uuid.UUID     			                            `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    			string        			                            `json:"name" gorm:"column:name; type:string; not_null"`
    TestUrl             string                                              `json:"testUrl" gorm:"column:test_url; type:string; null"` 
    Description 		string 		  			                            `json:"description" gorm:"column:description; type:string; not_null"` 
	CompanyId           uuid.UUID                                           `json:"companyId" gorm:"column:company_id; type:uuid;"`
    Company             company.Company                                     `json:"company" gorm:"foreignkey:CompanyId;"`
    DiagnosticTypeId  	uuid.UUID 	  			                            `json:"diagnosticTypeId" gorm:"column:diagnostic_type_id; type:uuid;`
	DiagnosticType  	diagnostic_type.DiagnosticType 	  		            `json:"diagnosticType" gorm:"foreignkey:diagnosticTypeId;`
	PocId   			uuid.UUID     			                            `json:"pocId" gorm:"column:poc_id; type:uuid;"`
    Poc     			poc.Poc 					                        `json:"poc" gorm:"foreignkey:PocId;"`
    VerifiedLod         string                                              `json:"verifiedLod" gorm:"column:verified_lod; type:string; null"`
    AvgCt               float64                                             `json:"avgCt" gorm:"column:avg_ct; type:numeric; null"`
    PrepIntegrated      bool                                                `json:"prepIntegrated" gorm:"column:prep_integrated; type:bool; not_null"`
    TestsPerRun         int64                                               `json:"testsPerRun" gorm:"column:tests_per_run; type:int; null"`
    TestsPerKit         int64                                               `json:"testsPerKit" gorm:"column:tests_per_kit; type:int; null"`
    CatalogNo           string                                              `json:"catalogNo" gorm:"column:catalog_no; type:string; null"`
    PointOfCare         bool                                                `json:"pointOfCare" gorm:"column:point_of_care; type:bool; not null"`
    CostPerKit          float64                                             `json:"costPerKit" gorm:"column:cost_per_kit; type:numeric; null"`
    InStock             bool                                                `json:"inStock" gorm:"column:in_stock; type:bool; not_null"`
    LeadTime            int64                                               `json:"leadTime" gorm:"column:lead_time; type:int; null"`
    CreatedBy           string                                              `json:"createdBy" gorm:"column:created_by; type:string; null"`
    Created             time.Time                                           `json:"created" gorm:"column:created; type:datetimetz; not null"`
    UpdatedBy           string                                              `json:"updatedBy" gorm:"column:updated_by; type:string; null"`
    Updated             time.Time                                           `json:"updated" gorm:"column:updated; type:datetimetz; not null"`
    RegulatoryApprovals	[]regulatory_approval_type.RegulatoryApprovalType   `json:"regulatoryApprovals" gorm:"many2many:diagnostic_regulatory_approvals;"`
    DiagnosticTargets	[]diagnostic_target_type.DiagnosticTargetType 	    `json:"diagnosticTargets" gorm:"many2many:diagnostic_targets;"`
    SampleTypes         []sample_type.SampleType                            `json:"sampleTypes" gorm:"many2many:diagnostic_sample_types;"`
    PcrPlatforms        []pcr_platform.PcrPlatform                          `json:"pcrPlatforms" gorm:"many2many:diagnostic_pcr_platforms;"`
    DxPerformance       []diagnostic_performance.DiagnosticPerformance      `json:"dxPerformance" gorm:"foreignkey:DiagnosticId"`
}

func (Diagnostic) TableName() string {
    return "diagnostics"
}

func Create(
             db *gorm.DB, name string, description string, testUrl string, company company.Company, 
             diagnosticType diagnostic_type.DiagnosticType, poc poc.Poc, 
             verifiedLod string, avgCt float64, prepIntegrated bool,
             testsPerRun int64, testsPerKit int64, catalogNo string, pointOfCare bool, 
             costPerKit float64, inStock bool, leadTime int64,
			 approvals []regulatory_approval_type.RegulatoryApprovalType, 
             targets []diagnostic_target_type.DiagnosticTargetType,
             sampleTypes []sample_type.SampleType,
             pcrPlatforms []pcr_platform.PcrPlatform) (*Diagnostic, error) {
    var toInsert = &Diagnostic{
        Name: name,
        Description: description,
        TestUrl: testUrl,
        CompanyId: company.Id,
        Company: company,
        DiagnosticTypeId: diagnosticType.Id,
        DiagnosticType: diagnosticType,
        PocId: poc.Id,
        Poc: poc,
        VerifiedLod: verifiedLod,
        AvgCt: avgCt,
        PrepIntegrated: prepIntegrated,
        TestsPerRun: testsPerRun,
        TestsPerKit: testsPerKit,
        CatalogNo: catalogNo,
        PointOfCare: pointOfCare,
        CostPerKit: costPerKit,
        InStock: inStock,
        LeadTime: leadTime,
        RegulatoryApprovals: approvals,
        DiagnosticTargets: targets,
        SampleTypes: sampleTypes,
        PcrPlatforms: pcrPlatforms,
        Created: time.Now(),
        Updated: time.Now(),
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *Diagnostic) (*Diagnostic, error) {
    toUpdate.Updated = time.Now();
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*Diagnostic, error) {
    result :=  &Diagnostic{}

    err := db.Preload("Company").Preload("Poc").Preload("DiagnosticType").Preload("RegulatoryApprovals").
    Preload("DiagnosticTargets").Preload("SampleTypes").Preload("PcrPlatforms").
    Preload("DiagnosticPerformance").Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]Diagnostic, error) {
    var results []Diagnostic =  nil

    err := db.Preload("Company").Preload("Poc").Preload("DiagnosticType").Preload("RegulatoryApprovals").
    Preload("DiagnosticTargets").Preload("SampleTypes").Preload("PcrPlatforms").
    Preload("DiagnosticPerformance").Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}