// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMonterStat = "monter_stats"

// MonterStat mapped from table <monter_stats>
type MonterStat struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Level    *int32 `gorm:"column:level" json:"level"`
	StatsID  int32  `gorm:"column:stats_id;not null" json:"stats_id"`
	MonterID int32  `gorm:"column:monter_id;not null" json:"monter_id"`
}

// TableName MonterStat's table name
func (*MonterStat) TableName() string {
	return TableNameMonterStat
}
