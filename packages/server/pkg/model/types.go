package model

type Note struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Title       string `gorm:"type:varchar(50);not null;column:title;index:title" json:"title"`
	Description string `gorm:"type:text;column:description" json:"description"`
	Content     string `gorm:"type:text;column:content" json:"content"`
	CreateAt    int64  `gorm:"type:bigint;column:create_at" json:"createAt"`
	UpdatedAt   int64  `gorm:"type:bigint;column:updated_at" json:"updatedAt"`
}

func (n Note) TableName() string {
	return "notes"
}
