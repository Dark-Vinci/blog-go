package entity

type User struct {
	ID       int64  `json:"id" gorm:"primary_key; auto_increment; type:int"`
	Name     string `json:"name" gorm:"type:varchar(255); not null;"`
	Email    string `json:"email" gorm:"uniqueIndex; type:varchar(255); not null"`
	Password string `json:"password" gorm:"type:varchar(255); not null"`
	Token    string `json:"token.omitempty" gorm:"-"`
}
