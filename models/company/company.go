package company

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "gitlab.iqt.org/rashley/covid-test-db/api/poc"
)

type Company struct {
    Id      uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    string        `json:"name" gorm:"column:name; type:string; not_null"` 
    City    string        `json:"city" gorm:"column:city; type:string;"`
    State   string        `json:"state" gorm:"column:state; type:string;"`
    Country string        `json:"country" gorm:"column:country; type:string;"`
    PocId   uuid.UUID     `json:"pocId" gorm:"column:poc_id; type:uuid;"`
    Poc     Poc           `json:"poc" gorm:"foreignkey:PocId"`
}

func (Company) TableName() string {
    return "companies"
}

func Create(db *gorm.DB, name string, city string, state string, country string, poc Poc) (*Company, error) {
    var toInsert = &Company{
        Name: name,
        City: city,
        State: state,
        Country: country,
        PocId: poc.Id
        Poc: poc
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

func FetchList(db *gorm.DB) ([]Company, error) {
    var results []Company =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}