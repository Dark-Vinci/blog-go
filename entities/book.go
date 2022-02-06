package entity

type Book struct {
	ID          int64  `json:"id" gorm:"primary_key; auto_increment; type:int"`
	Title       string `json:"title" gorm:"type:varchar(255); not null;column:title"`
	Description string `json:"description" gorm:"column:description; type:text; not null"`
	UserID      int64  `json:"user_id" gorm:"not null"`
	User        User   `json:"user" gorm:"foreignKey:UserID; constarint:onUpdate:CASCADE, onDelete:CASCADE"`
}
