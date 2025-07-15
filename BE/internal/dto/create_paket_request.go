package dto

type CreatePaketRequest struct {
	Wilayah  string  `json:"wilayah" binding:"required"`
	Tipe     string  `json:"tipe" binding:"required"`
	Level    string  `json:"level" binding:"required"`
	Nama     string  `json:"nama"`
	Harga    float64 `json:"harga" binding:"required"`
	IsActive bool    `json:"is_active"`
}
