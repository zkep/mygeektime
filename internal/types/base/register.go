package base

type RegisterRequest struct {
	Account         string `json:"account,omitempty" validate:"min=5,max=50"`
	Password        string `json:"password,omitempty" validate:"min=5,max=50"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"min=5,max=50"`
	Type            string `json:"type,omitempty" validate:"min=5,max=50"`
}
