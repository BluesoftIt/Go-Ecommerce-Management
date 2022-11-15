package models

// store information structure
type Reseller struct {
	ID           int64  `gorm:"primary_key:auto_increment" json:"id"`
	Name         string `gorm:"type:varcher(255)" json:"name"`
	StoreId      string `gorm:"uniqueIndex;type:varcher(255)" json:"store_id"`
	Profile      string `gorm:"type:varcher(255)" json:"profile"`
	Email        string `gorm:"uniqueIndex;type:varcher(255)" json:"email"`
	Password     string `gorm:"->;<-;not null" json:"-"`
	Token        string `gorm:"-" json: "token,omitempty"`
	MonthEarning int64  `gorm:"type:int" json:"month-earning"`
	YearEarning  int64  `gorm:"type:int" json:"year_earning"`
}
