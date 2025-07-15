package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterMemberRequest struct {
	IDMember        string             `json:"id_member" binding:"required"`
	Name            string             `json:"name" binding:"required"`
	Gender          string             `json:"gender" binding:"required"`
	BirthPlace      string             `json:"birth_place" binding:"required"`
	BirthDate       string             `json:"birth_date" binding:"required"`
	PhoneNumber     string             `json:"phone_number" binding:"required"`
	Email           string             `json:"email" binding:"required,email"`
	NoKTP           string             `json:"no_ktp" binding:"required"`
	BankName        string             `json:"bank_name"`
	BankAccountNo   string             `json:"bank_account_no"`
	SponsorIDMember string             `json:"sponsor_id_member"`
	SponsorName     string             `json:"sponsor_name"`
	UplineIDMember  string             `json:"upline_id_member"`
	UplineName      string             `json:"upline_name"`
	PaketID         primitive.ObjectID `json:"paket_id" binding:"required"`
	AlamatLengkap   string             `json:"alamat_lengkap" binding:"required"`
	ManagerArea     string             `json:"manager_area"`
	KirimKeStockist bool               `json:"kirim_ke_stockist"`
	TempatDaftar    string             `json:"tempat_daftar"`
}
