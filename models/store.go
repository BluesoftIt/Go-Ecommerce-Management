package models

// store information structure
type Store struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varcher(255)" json:"name"`
	StoreId  string `gorm:"uniqueIndex;type:varcher(255)" json:"store_id"`
	Email    string `gorm:"uniqueIndex;type:varcher(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json: "token,omitempty"`
}
