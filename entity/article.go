package entity

type Article struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"type:varchar(100)" json:"title"`
	Slug string `gorm:"type:varchar(150)" json:"slug"`
	Description string `gorm:"type:text" json:"description"`
	Image string `gorm:"type:varchar(150)" json:"image"`
	UserID uint64 `gorm:"not null" json:"-"`
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`


}