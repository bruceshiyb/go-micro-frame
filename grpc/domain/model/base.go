package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID        int64         `gorm:"primarykey;type:int;comment:主键" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"-"`
	UpdatedAt time.Time      `gorm:"type:timestamp not null;default:current_timestamp;comment:修改时间" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"comment:删除时间" json:"-"`
}
