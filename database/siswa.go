package database

type Siswa struct {
	Id   int64  `gorm:"primarykey" json:"id"`
	Nama string `gorm:"type:varchar(255)" json:"nama"`
	Umur int64  `gorm:"type:int(2)" json:"umur"`
}
