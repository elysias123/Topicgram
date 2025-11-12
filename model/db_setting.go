package model

type Setting struct {
	Key   string `gorm:"column:key; primaryKey; not null"`
	Value string `gorm:"column:value; not null"`
}

func (*Setting) TableName() string {
	return "settings"
}
