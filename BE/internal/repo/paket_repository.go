package repo

import (
	"context"
	"registration-smart-reward/internal/config"
	"registration-smart-reward/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaketRepository interface {
	Insert(ctx context.Context, paket *model.Paket) error
	FindAllActive(ctx context.Context) ([]model.Paket, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.Paket, error)
	CountAll(ctx context.Context) (int64, error)
}

type paketRepository struct {
	col *mongo.Collection
}

func NewPaketRepository() PaketRepository {
	return &paketRepository{
		col: config.GetMongoDB().Collection("paket"),
	}
}

func (r *paketRepository) Insert(ctx context.Context, paket *model.Paket) error {
	_, err := r.col.InsertOne(ctx, paket)
	return err
}

func (r *paketRepository) FindAllActive(ctx context.Context) ([]model.Paket, error) {
	filter := bson.M{"is_active": true}
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []model.Paket
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *paketRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Paket, error) {
	var paket model.Paket
	err := r.col.FindOne(ctx, bson.M{"_id": id, "is_active": true}).Decode(&paket)
	if err != nil {
		return nil, err
	}
	return &paket, nil
}

func (r *paketRepository) CountAll(ctx context.Context) (int64, error) {
	return r.col.CountDocuments(ctx, bson.M{})
}
