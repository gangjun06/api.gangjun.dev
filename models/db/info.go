package db

type Info struct {
	ID    int    `gorm:"autoIncrement;primaryKey;not null"`
	Key   string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Value string `gorm:"not null"`
}
