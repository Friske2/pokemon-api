// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameType = "types"

// Type mapped from table <types>
type Type struct {
	ID   int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name *string `gorm:"column:name" json:"name"`
}

// TableName Type's table name
func (*Type) TableName() string {
	return TableNameType
}
