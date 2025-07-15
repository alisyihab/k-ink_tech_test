package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paket struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Wilayah   string             `bson:"wilayah" json:"wilayah"`     // WIL.A / WIL.B
	Tipe      string             `bson:"tipe" json:"tipe"`           // Fast Track / Smart Reward / dll
	Level     string             `bson:"level" json:"level"`         // Silver / Gold / Platinum
	Nama      string             `bson:"nama" json:"nama"`           // Nama paket (opsional)
	Harga     float64            `bson:"harga" json:"harga"`         // Harga paket
	IsActive  bool               `bson:"is_active" json:"is_active"` // Paket masih berlaku?
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
