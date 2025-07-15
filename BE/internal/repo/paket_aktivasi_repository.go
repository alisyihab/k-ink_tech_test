package repo

import (
	"context"
	"registration-smart-reward/internal/config"
	"registration-smart-reward/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type PaketAktivasiRepository interface {
	Insert(ctx context.Context, data *model.PaketAktivasi) error
}

type paketAktivasiRepo struct {
	col *mongo.Collection
}

func NewPaketAktivasiRepository() PaketAktivasiRepository {
	return &paketAktivasiRepo{
		col: config.GetMongoDB().Collection("paket_aktivasi"),
	}
}

func (r *paketAktivasiRepo) Insert(ctx context.Context, data *model.PaketAktivasi) error {
	_, err := r.col.InsertOne(ctx, data)
	return err
}
