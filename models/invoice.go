package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	ClientID  uint           `json:"client_id" gorm:"not null;index"`
	IssueDate time.Time      `json:"issue_date" gorm:"not null"`
	DueDate   time.Time      `json:"due_date" gorm:"not null"`
	Status    string         `json:"status" gorm:"not null;default:'draft'"`
	Notes     string         `json:"notes" gorm:"type:text"`
	Subtotal  float64        `json:"subtotal" gorm:"not null;default:0"`
	Tax       float64        `json:"tax" gorm:"not null;default:0"`
	TaxRate   float64        `json:"tax_rate" gorm:"not null;default:0"`
	Total     float64        `json:"total" gorm:"not null;default:0"`
	Currency  string         `json:"currency" gorm:"not null;default:'USD'"`
	Items     []InvoiceItem  `json:"items" gorm:"foreignKey:InvoiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
