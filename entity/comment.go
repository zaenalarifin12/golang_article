package entity

type Comment struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Body string `gorm:"text" json:"body"`
	UserID uint64 `gorm:"not null" json:"-"`
	ArticleID uint64 `gorm:"not null" json:"article_id"`
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Article Article `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
