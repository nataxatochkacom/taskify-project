package models

type SharedAccess struct {
    ID        uint   `gorm:"primaryKey"`
    UUID      string `gorm:"column:uuid" json:"uuid"`
    Role      string `gorm:"column:role" json:"role"` // admin, manager, user, viewer
    OwnerID   uint   `gorm:"column:owner_id" json:"owner_id"`
    CreatedAt string `gorm:"column:created_at" json:"created_at"`
}
