package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaketAktivasiRequest struct {
	MemberID primitive.ObjectID `json:"member_id" binding:"required"`
	PaketID  primitive.ObjectID `json:"paket_id"  binding:"required"`
}
