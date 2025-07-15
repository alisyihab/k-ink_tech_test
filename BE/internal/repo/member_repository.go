package repo

import (
	"context"
	"registration-smart-reward/internal/config"
	"registration-smart-reward/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemberRepository interface {
	Insert(ctx context.Context, member *model.Member) error
	ExistsDuplicate(ctx context.Context, idMember, ktp, email, phone, rekening string) (bool, error)
	FindByIDMember(ctx context.Context, idMember string) (*model.Member, error)
	FindPaginated(ctx context.Context, page, limit int, search string) ([]model.Member, int64, error)
	CountAll(ctx context.Context) (int64, error)
}

type memberRepository struct {
	col *mongo.Collection
}

func NewMemberRepository() MemberRepository {
	return &memberRepository{
		col: config.GetMongoDB().Collection("member"),
	}
}

func (r *memberRepository) Insert(ctx context.Context, member *model.Member) error {
	_, err := r.col.InsertOne(ctx, member)
	return err
}

// Cek apakah ada data yang bentrok (unique constraint manual)
func (r *memberRepository) ExistsDuplicate(ctx context.Context, idMember, ktp, email, phone, rekening string) (bool, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"id_member": idMember},
			{"no_ktp": ktp},
			{"email": email},
			{"phone_number": phone},
			{"bank_account_no": rekening},
		},
	}

	count, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Cari member berdasarkan kode id_member (dipakai di sponsor/upline)
func (r *memberRepository) FindByIDMember(ctx context.Context, idMember string) (*model.Member, error) {
	var member model.Member
	err := r.col.FindOne(ctx, bson.M{"id_member": idMember}).Decode(&member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) FindPaginated(ctx context.Context, page, limit int, search string) ([]model.Member, int64, error) {
	filter := bson.M{}
	if search != "" {
		filter["$or"] = []bson.M{
			{"id_member": bson.M{"$regex": search, "$options": "i"}},
			{"name": bson.M{"$regex": search, "$options": "i"}},
			{"email": bson.M{"$regex": search, "$options": "i"}},
			{"phone_number": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	// count total
	total, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit64).
		SetSort(bson.M{"created_at": -1})

	cursor, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var members []model.Member
	if err := cursor.All(ctx, &members); err != nil {
		return nil, 0, err
	}

	return members, total, nil
}

func (r *memberRepository) CountAll(ctx context.Context) (int64, error) {
	return r.col.CountDocuments(ctx, bson.M{})
}
