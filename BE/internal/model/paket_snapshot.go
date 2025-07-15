package model

type PaketSnapshot struct {
	Wilayah string  `bson:"wilayah" json:"wilayah"`
	Tipe    string  `bson:"tipe" json:"tipe"`
	Level   string  `bson:"level" json:"level"`
	Nama    string  `bson:"nama" json:"nama"`
	Harga   float64 `bson:"harga" json:"harga"`
}
