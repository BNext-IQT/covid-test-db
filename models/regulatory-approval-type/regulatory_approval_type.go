package regulatory_approval_type

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type RegulatoryApprovalType struct {
    Id      uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    string        `json:"name" gorm:"column:name; type:string; not_null"` 
}

func (RegulatoryApprovalType) TableName() string {
    return "regulatory_approval_types"
}

func Create(db *gorm.DB, name string) (*RegulatoryApprovalType, error) {
    var toInsert = &RegulatoryApprovalType{
        Name: name,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *RegulatoryApprovalType) (*RegulatoryApprovalType, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*RegulatoryApprovalType, error) {
    result :=  &RegulatoryApprovalType{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]RegulatoryApprovalType, error) {
    var results []RegulatoryApprovalType =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}