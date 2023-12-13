package model

type BlogPost struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (BlogPost) TableName() string {
	return "blog_post"
}
