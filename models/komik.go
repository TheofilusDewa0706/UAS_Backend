package models

type Komik struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Nama        string `json:"nama"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	TahunTerbit int    `json:"tahun_terbit"`
	Publisher   string `json:"publisher"`
	Stok        int    `json:"stok"`
}
