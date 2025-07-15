package service

import (
	"context"
	"errors"
	"registration-smart-reward/internal/dto"
	"registration-smart-reward/internal/model"
	"registration-smart-reward/internal/repo"
	"time"
)

type PaketAktivasiService interface {
	Activate(ctx context.Context, input *dto.PaketAktivasiRequest) error
}

type paketAktivasiService struct {
	paketRepo    repo.PaketRepository
	aktivasiRepo repo.PaketAktivasiRepository
}

func NewPaketAktivasiService(paketRepo repo.PaketRepository, aktivasiRepo repo.PaketAktivasiRepository) PaketAktivasiService {
	return &paketAktivasiService{
		paketRepo:    paketRepo,
		aktivasiRepo: aktivasiRepo,
	}
}

func (s *paketAktivasiService) Activate(ctx context.Context, input *dto.PaketAktivasiRequest) error {
	paket, err := s.paketRepo.FindByID(ctx, input.PaketID)
	if err != nil || paket == nil {
		return errors.New("paket tidak ditemukan atau tidak aktif")
	}

	snapshot := model.PaketSnapshot{
		Wilayah: paket.Wilayah,
		Tipe:    paket.Tipe,
		Level:   paket.Level,
		Nama:    paket.Nama,
		Harga:   paket.Harga,
	}

	data := &model.PaketAktivasi{
		MemberID:    input.MemberID,
		PaketID:     input.PaketID,
		Snapshot:    snapshot,
		ActivatedAt: time.Now(),
	}

	return s.aktivasiRepo.Insert(ctx, data)
}
