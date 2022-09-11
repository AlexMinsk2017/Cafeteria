package web

import "Cafeteria/pkg/common/models/dto"

type User struct {
	Id       uint
	User     string
	Password string
}

type Users []*User

func (r *User) ToDTO() *dto.User {
	if r == nil {
		return nil
	}
	return &dto.User{
		Id:       r.Id,
		User:     r.User,
		Password: r.Password,
	}
}
func (r *User) FromDTO(v *dto.User) *User {
	if v == nil {
		return nil
	}
	return &User{
		Id:       v.Id,
		User:     v.User,
		Password: v.Password,
	}
}
