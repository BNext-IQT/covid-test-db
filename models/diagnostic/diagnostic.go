package diagnostic

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "gitlab.iqt.org/rashley/covid-test-db/models/poc"
    "gitlab.iqt.org/rashley/covid-test-db/models/company"
    "gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_type"
    "gitlab.iqt.org/rashley/covid-test-db/models/diagnostic_target_type"
    "gitlab.iqt.org/rashley/covid-test-db/models/regulatory_approval_type"
)

type Diagnostic struct {
    Id      			uuid.UUID     			                            `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    			string        			                            `json:"name" gorm:"column:name; type:string; not_null"` 
    Description 		string 		  			                            `json:"description" gorm:"column:description; type:string; not_null"` 
	CompanyId           uuid.UUID                                           `json:"companyId" gorm:"column:company_id; type:uuid;"`
    Company             company.Company                                     `json:"company" gorm:"foreignkey:CompanyId;"`
    DiagnosticTypeId  	uuid.UUID 	  			                            `json:"diagnosticTypeId" gorm:"column:diagnostic_type_id; type:uuid;`
	DiagnosticType  	diagnostic_type.DiagnosticType 	  		            `json:"diagnosticType" gorm:"foreignkey:diagnosticTypeId;`
	PocId   			uuid.UUID     			                            `json:"pocId" gorm:"column:poc_id; type:uuid;"`
    Poc     			poc.Poc 					                        `json:"poc" gorm:"foreignkey:PocId;"`
    RegulatoryApprovals	[]regulatory_approval_type.RegulatoryApprovalType   `gorm:"many2many:regulatory_approval_types;"`
    DiagnosticTargets	[]diagnostic_target_type.DiagnosticTargetType 	    `gorm:"many2many:diagnostic_targets;"`
}

func (Diagnostic) TableName() string {
    return "diagnostics"
}

func Create(
             db *gorm.DB, name string, description string, company company.Company, 
             diagnosticType diagnostic_type.DiagnosticType, poc poc.Poc, 
			 approvals []regulatory_approval_type.RegulatoryApprovalType, 
             targets []diagnostic_target_type.DiagnosticTargetType) (*Diagnostic, error) {
    var toInsert = &Diagnostic{
        Name: name,
        Description: description,
        CompanyId: company.Id,
        Company: company,
        DiagnosticTypeId: diagnosticType.Id,
        DiagnosticType: diagnosticType,
        PocId: poc.Id,
        Poc: poc,
        RegulatoryApprovals: approvals,
        DiagnosticTargets: targets,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *Diagnostic) (*Diagnostic, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*Diagnostic, error) {
    result :=  &Diagnostic{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]Diagnostic, error) {
    var results []Diagnostic =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}