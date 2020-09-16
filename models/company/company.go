package company

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/BNext-IQT/covid-test-db/models/poc"
)

type Company struct {
    Id              uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name            string        `json:"name" gorm:"column:name; type:string; not_null"` 
    StreetAddress   string        `json:"streetAddress" gorm:"column:street_address; type:string;"`
    City            string        `json:"city" gorm:"column:city; type:string;"`
    State           string        `json:"state" gorm:"column:state; type:string;"`
    Country         string        `json:"country" gorm:"column:country; type:string;"`
    PostalCode      string        `json:"postalCode" gorm:"column:postal_code; type:string;"`
    Stage           string        `json:"stage" gorm:"column:stage; type:string;"`
    Valuation       string        `json:"valuation" gorm:"column:valuation; type:string;"`
    PocId           uuid.UUID     `json:"pocId" gorm:"column:poc_id; type:uuid;"`
    Poc             poc.Poc       `json:"poc" gorm:"foreignkey:PocId"`
}

func (Company) TableName() string {
    return "companies"
}

func Create(db *gorm.DB, name string, streetAddress string, city string, state string, 
        country string, postalCode string, stage string, valuation string, poc poc.Poc) (*Company, error) {
    var toInsert = &Company{
        Name: name,
        StreetAddress: streetAddress,
        City: city,
        State: state,
        Country: country,
        PostalCode: postalCode,
        Stage: stage,
        Valuation: valuation,
        PocId: poc.Id,
        Poc: poc,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *Company) (*Company, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*Company, error) {
    result :=  &Company{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchByName(db *gorm.DB, name string) (*Company, error) {
    result :=  &Company{}

    err := db.Where("name = ?", name).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]Company, error) {
    var results []Company =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}