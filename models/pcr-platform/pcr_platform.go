package pcr_platform

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PcrPlatform struct {
    Id      uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    string        `json:"name" gorm:"column:name; type:string; not_null"` 
}

func (PcrPlatform) TableName() string {
    return "pcr_platforms"
}

func Create(db *gorm.DB, name string) (*PcrPlatform, error) {
    var toInsert = &PcrPlatform{
        Name: name,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *PcrPlatform) (*PcrPlatform, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*PcrPlatform, error) {
    result :=  &PcrPlatform{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchByName(db *gorm.DB, name string) (*PcrPlatform, error) {
    result :=  &PcrPlatform{}

    err := db.Where("name = ?", name).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]PcrPlatform, error) {
    var results []PcrPlatform =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}