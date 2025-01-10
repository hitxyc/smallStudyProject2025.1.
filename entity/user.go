package entity

type UserEntity struct {
	UID      int    `gorm:"primary_key" form:"uid" json:"uid"`
	UserName string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}

// 返回表名
func (*UserEntity) TableName() string {
	return "user_entity"
}
