package models

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Poc struct {
    id   uuid.UUID `gorm:"type:uuid;primary_key;"`
    name string 
    phone string
    email string
}

// Functions of type `txnFunc` are passed as arguments to our
// `runTransaction` wrapper that handles transaction retries for us
// (see implementation below).
type txnFunc func(*gorm.DB) error

// This function is used for testing the transaction retry loop.  It
// can be deleted from production code.
var forceRetryLoop txnFunc = func(db *gorm.DB) error {

    // The first statement in a transaction can be retried transparently
    // on the server, so we need to add a dummy statement so that our
    // force_retry statement isn't the first one.
    if err := db.Exec("SELECT now()").Error; err != nil {
        return err
    }
    // Used to force a transaction retry.
    if err := db.Exec("SELECT crdb_internal.force_retry('1s'::INTERVAL)").Error; err != nil {
        return err
    }
    return nil
}

func Create(db *gorm.DB, name string, email string, phone string) (*Poc, error) {
    toInsert := &Poc{}
    id := uuid.New()
    toInsert.id = id
    toInsert.name = name
    toInsert.email = email
    toInsert.phone = phone

    err := db.Save(&toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*Poc, error) {
    result :=  &Poc{}

    err := db.First(&result, id).Error;

    if err != nil {
        result = nil
    }

    return result, err
}