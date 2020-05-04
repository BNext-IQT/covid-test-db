package test

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "gitlab.iqt.org/rashley/covid-test-db/api/poc"
    "gitlab.iqt.org/rashley/covid-test-db/api/test_type"
    "gitlab.iqt.org/rashley/covid-test-db/api/test_target_type"
    "gitlab.iqt.org/rashley/covid-test-db/api/regulatory_approval_type"
)

type Test struct {
    Id      			uuid.UUID     			 `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    			string        			 `json:"name" gorm:"column:name; type:string; not_null"` 
    Description 		string 		  			 `json:"description" gorm:"column:description; type:string; not_null"` 
	TestTypeId  		uuid.UUID 	  			 `json:"testTypeId" gorm:"column:test_type_id; type:uuid;`
	TestType  			TestType 	  			 `json:"testType" gorm:"foreignkey:TestTypeId;`
	PocId   			uuid.UUID     			 `json:"pocId" gorm:"column:poc_id; type:uuid;"`
    Poc     			Poc 					 `json:"poc" gorm:"foreignkey:PocId;"`
    RegulatoryApprovals	[]RegulatoryApprovalType `gorm:"many2many:regulatory_approval_types;"`
    TestTargets			[]TestTargetType 		 `gorm:"many2many:test_target_types;"`
}

func (Test) TableName() string {
    return "tests"
}

func Create(db *gorm.DB, name string, description string, testType TestType, poc Poc, 
			approvals []RegulatoryApprovalType, targets []TestTargetType) (*Test, error) {
    var toInsert = &Test{
        Name: name,
        Description: description,
        TestTypeId: testType.Id,
        TestType: testType,
        PocId: poc.Id,
        Poc: poc,
        RegulatoryApprovals: approvals,
        TestTargets: targets,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *Test) (*Test, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*Test, error) {
    result :=  &Test{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]Test, error) {
    var results []Test =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}