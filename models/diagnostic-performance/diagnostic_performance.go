package diagnostic_performance

import (
	"time"

    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type DiagnosticPerformance struct {
    DiagnosticId      	uuid.UUID     `json:"id" gorm:"column:id; type:uuid;"`
    SourceOfPerfData    string        `json:"sourceOfPerfData" gorm:"column:source_of_perf_data; type:string; null"`
    SourceDisplayName   string        `json:"sourceDisplayName" gorm:"column:source_display_name; type:string; null"`
    Tp 					int64         `json:"tp" gorm:"column:tp; type:int; null"`
	Fp 					int64         `json:"fp" gorm:"column:fp; type:int; null"`
	Tn 					int64         `json:"tn" gorm:"column:tn; type:int; null"`
	Fn 					int64         `json:"fn" gorm:"column:fn; type:int; null"`
	Sensitivity         float64       `json:"sensitivity" gorm:"column:sensitivity; type:numeric; null"`
	SensitivityCiLow	float64       `json:"sensitivityCiLow" gorm:"column:sensitivity_ci_low; type:numeric; null"`
	SensitivityCiHigh	float64       `json:"sensitivityCiHigh" gorm:"column:sensitivity_ci_high; type:numeric; null"`
	Specificity         float64       `json:"specificity" gorm:"column:specificity; type:numeric; null"`
	SpecificityCiLow	float64       `json:"specificityCiLow" gorm:"column:specificity_ci_low; type:numeric; null"`
	SpecificityCiHigh	float64       `json:"specificityCiHigh" gorm:"column:specificity_ci_high; type:numeric; null"`
	CreatedBy           string        `json:"createdBy" gorm:"column:created_by; type:string; null"`
    Created             time.Time     `json:"created" gorm:"column:created; type:datetimetz; not null"`
    UpdatedBy           string        `json:"updatedBy" gorm:"column:updated_by; type:string; null"`
    Updated             time.Time     `json:"updated" gorm:"column:updated; type:datetimetz; not null"`
}

func (DiagnosticPerformance) TableName() string {
    return "diagnostic_performance"
}

func Create(db *gorm.DB, diagnosticId uuid.UUID, sourceOfPerfData string, sourceDisplayName string,
    tp int64, fp int64, tn int64,fn int64, sensitivity float64,
	sensitivityCiLow float64, sensitivityCiHigh float64, specificity float64, specificityCiLow float64,
	specificityCiHigh float64, createdBy string, updatedBy string ) (*DiagnosticPerformance, error) {
    var toInsert = &DiagnosticPerformance{
        DiagnosticId: diagnosticId,
	    SourceOfPerfData: sourceOfPerfData,
	    SourceDisplayName: sourceDisplayName,
	    Tp: tp,
		Fp: fp,
		Tn: tn,
		Fn: fn,
		Sensitivity: sensitivity,
		SensitivityCiLow: sensitivityCiLow,
		SensitivityCiHigh: sensitivityCiHigh,
		Specificity: specificity,
		SpecificityCiLow: specificityCiLow,
		SpecificityCiHigh: specificityCiHigh,
		CreatedBy: createdBy,
	    Created: time.Now(),
	    UpdatedBy: updatedBy,
	    Updated: time.Now(),
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *DiagnosticPerformance) (*DiagnosticPerformance, error) {
	toUpdate.Updated = time.Now();
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*DiagnosticPerformance, error) {
    result :=  &DiagnosticPerformance{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchByName(db *gorm.DB, name string) (*DiagnosticPerformance, error) {
    result :=  &DiagnosticPerformance{}

    err := db.Where("source_display_name = ?", name).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]DiagnosticPerformance, error) {
    var results []DiagnosticPerformance =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}