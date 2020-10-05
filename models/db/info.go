package db

type Info struct {
	ID      int    `gorm:"autoIncrement;primaryKey;not null"`
	DataKey string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Value   string `gorm:"not null"`
}
