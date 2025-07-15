package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaketAktivasi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MemberID    primitive.ObjectID `bson:"member_id" json:"member_id"`
	PaketID     primitive.ObjectID `bson:"paket_id" json:"paket_id"`
	Snapshot    PaketSnapshot      `bson:"snapshot" json:"snapshot"`
	ActivatedAt time.Time          `bson:"activated_at" json:"activated_at"`
}
