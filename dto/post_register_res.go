package dto

type RegisterRes struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (rr *RegisterRes) FromRegister(b *models.Register) *RegisterRes {
	return &RegisterRes{
		ID:    b.ID,
		Name:  b.Name,
		Email: b.Name,
	}
}
