package models

type BaseMySQLID struct {
	ID uint `json:"id" gorm:"primarykey;column:id;type:bigint unsigned;not null"`
}

type BaseMySQLDates struct {
	CreatedAt string `json:"create_at" gorm:"column:created_at;type:timestamp(6);default:CURRENT_TIMESTAMP(6)"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;type:timestamp(6);default:NULL ON UPDATE CURRENT_TIMESTAMP(6)"`
}
