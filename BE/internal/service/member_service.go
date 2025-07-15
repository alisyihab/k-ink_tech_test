package service

import (
	"context"
	"errors"
	"registration-smart-reward/internal/dto"
	"registration-smart-reward/internal/model"
	"registration-smart-reward/internal/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemberService interface {
	Register(ctx context.Context, input *dto.RegisterMemberRequest) error
	ListPaginated(ctx context.Context, page, limit int, search string) ([]model.Member, int64, error)
}

type memberService struct {
	memberRepo   repo.MemberRepository
	paketRepo    repo.PaketRepository
	aktivasiRepo repo.PaketAktivasiRepository
}

func NewMemberService(mr repo.MemberRepository, pr repo.PaketRepository, ar repo.PaketAktivasiRepository) MemberService {
	return &memberService{
		memberRepo:   mr,
		paketRepo:    pr,
		aktivasiRepo: ar,
	}
}

func (s *memberService) Register(ctx context.Context, input *dto.RegisterMemberRequest) error {
	// Cek data duplikat
	dup, err := s.memberRepo.ExistsDuplicate(ctx, input.IDMember, input.NoKTP, input.Email, input.PhoneNumber, input.BankAccountNo)
	if err != nil {
		return err
	}
	if dup {
		return errors.New("id_member, no_ktp, email, no_hp, atau rekening sudah digunakan")
	}

	// Parse tanggal lahir (kalau masih mau disimpan sebagai string, bisa skip ini step)
	_, err = time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return errors.New("format tanggal lahir tidak valid (contoh: 2000-01-02)")
	}

	// Ambil sponsor
	var sponsorID *primitive.ObjectID
	if input.SponsorIDMember != "" {
		sponsor, err := s.memberRepo.FindByIDMember(ctx, input.SponsorIDMember)
		if err != nil {
			return errors.New("sponsor tidak ditemukan")
		}
		sponsorID = &sponsor.ID
	}

	// Ambil upline
	var uplineID *primitive.ObjectID
	if input.UplineIDMember != "" {
		upline, err := s.memberRepo.FindByIDMember(ctx, input.UplineIDMember)
		if err != nil {
			return errors.New("upline tidak ditemukan")
		}
		uplineID = &upline.ID
	}

	// Ambil info paket
	paket, err := s.paketRepo.FindByID(ctx, input.PaketID)
	if err != nil || paket == nil {
		return errors.New("paket tidak ditemukan")
	}

	// Snapshot langsung ke member
	snapshot := model.PaketSnapshot{
		Wilayah: paket.Wilayah,
		Tipe:    paket.Tipe,
		Level:   paket.Level,
		Nama:    paket.Nama,
		Harga:   paket.Harga,
	}

	now := time.Now()

	member := model.Member{
		ID:              primitive.NewObjectID(),
		IDMember:        input.IDMember,
		Name:            input.Name,
		Gender:          input.Gender,
		BirthPlace:      input.BirthPlace,
		BirthDate:       input.BirthDate,
		PhoneNumber:     input.PhoneNumber,
		Email:           input.Email,
		NoKTP:           input.NoKTP,
		BankName:        input.BankName,
		BankAccountNo:   input.BankAccountNo,
		AlamatLengkap:   input.AlamatLengkap,
		ManagerArea:     input.ManagerArea,
		KirimKeStockist: input.KirimKeStockist,
		TempatDaftar:    input.TempatDaftar,
		TanggalDaftar:   now,
		SponsorID:       sponsorID,
		SponsorName:     input.SponsorName,
		UplineID:        uplineID,
		UplineName:      input.UplineName,
		PaketID:         input.PaketID,
		PaketSnapshot:   snapshot,
		Path:            "",
		Level:           0,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := s.memberRepo.Insert(ctx, &member); err != nil {
		return err
	}

	return nil
}

func (s *memberService) ListPaginated(ctx context.Context, page, limit int, search string) ([]model.Member, int64, error) {
	return s.memberRepo.FindPaginated(ctx, page, limit, search)
}
