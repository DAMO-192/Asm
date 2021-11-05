package dto

import "Asm/moled"

type Userdto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user *moled.User) Userdto {
	return Userdto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
