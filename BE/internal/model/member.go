package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDMember      string             `bson:"id_member" json:"id_member"`
	Name          string             `bson:"name" json:"name"`
	AlamatLengkap string             `bson:"alamat_lengkap" json:"alamat_lengkap"`
	Gender        string             `bson:"gender" json:"gender"`
	NoKTP         string             `bson:"no_ktp" json:"no_ktp"`
	BirthPlace    string             `bson:"birth_place" json:"birth_place"`
	BirthDate     string             `bson:"birth_date" json:"birth_date"`
	PhoneNumber   string             `bson:"phone_number" json:"phone_number"`
	Email         string             `bson:"email" json:"email"`
	BankName      string             `bson:"bank_name" json:"bank_name"`
	BankAccountNo string             `bson:"bank_account_no" json:"bank_account_no"`

	UplineID    *primitive.ObjectID `bson:"upline_id,omitempty" json:"upline_id,omitempty"`
	UplineName  string              `bson:"upline_name" json:"upline_name"`
	SponsorID   *primitive.ObjectID `bson:"sponsor_id,omitempty" json:"sponsor_id,omitempty"`
	SponsorName string              `bson:"sponsor_name" json:"sponsor_name"`

	PaketID       primitive.ObjectID `bson:"paket_id" json:"paket_id"`
	PaketSnapshot PaketSnapshot      `bson:"paket_snapshot" json:"paket_snapshot"`

	ManagerArea     string `bson:"manager_area" json:"manager_area"`
	KirimKeStockist bool   `bson:"kirim_ke_stockist" json:"kirim_ke_stockist"`

	TempatDaftar  string    `bson:"tempat_daftar" json:"tempat_daftar"`
	TanggalDaftar time.Time `bson:"tanggal_daftar" json:"tanggal_daftar"`

	Path  string `bson:"path" json:"path"`
	Level int    `bson:"level" json:"level"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
