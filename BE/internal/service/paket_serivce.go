package service

import (
	"context"
	"registration-smart-reward/internal/dto"
	"registration-smart-reward/internal/model"
	"registration-smart-reward/internal/repo"
	"time"
)

type PaketService interface {
	Create(ctx context.Context, input *dto.CreatePaketRequest) error
	ListActive(ctx context.Context) ([]model.Paket, error)
}

type paketService struct {
	repo repo.PaketRepository
}

func NewPaketService(r repo.PaketRepository) PaketService {
	return &paketService{repo: r}
}

func (s *paketService) Create(ctx context.Context, input *dto.CreatePaketRequest) error {
	paket := &model.Paket{
		Wilayah:   input.Wilayah,
		Tipe:      input.Tipe,
		Level:     input.Level,
		Nama:      input.Nama,
		Harga:     input.Harga,
		IsActive:  input.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.Insert(ctx, paket)
}

func (s *paketService) ListActive(ctx context.Context) ([]model.Paket, error) {
	return s.repo.FindAllActive(ctx)
}
