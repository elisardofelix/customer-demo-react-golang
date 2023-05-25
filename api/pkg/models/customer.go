package models

type Customer struct {
	BaseMySQLID
	FirstName string `json:"first_name" gorm:"column:first_name;type:varchar(255);not null"`
	LastName  string `json:"last_name"  gorm:"column:last_name;type:varchar(255); not null"`
	Email     string `json:"email" gorm:"type:varchar(255)"`
	Gender    string `json:"gender" gorm:"type:varchar(255)"`
	Company   string `json:"company" gorm:"type:varchar(255)"`
	City      string `json:"city" gorm:"type:varchar(255)"`
	Title     string `json:"title" gorm:"type:varchar(255)"`
	BaseMySQLDates
}
