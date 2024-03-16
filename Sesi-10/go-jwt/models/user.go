package models

type User struct {
	Fullname string    `gorm:"not null" json:"fullname" form:"fullname" valid:"required-Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required-Your  email is required, email-Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required-Your password is required,minstringlength(6)-Password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,Ondelete:SET NULL;" json:"products"`
}
