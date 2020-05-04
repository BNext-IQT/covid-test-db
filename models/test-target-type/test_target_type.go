package test_target_type

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type TestTargetType struct {
    Id      uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    string        `json:"name" gorm:"column:name; type:string; not_null"` 
}

func (TestTargetType) TableName() string {
    return "test_target_types"
}

func Create(db *gorm.DB, name string) (*TestTargetType, error) {
    var toInsert = &TestTargetType{
        Name: name,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *TestTargetType) (*TestTargetType, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*TestTargetType, error) {
    result :=  &TestTargetType{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]TestTargetType, error) {
    var results []TestTargetType =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}