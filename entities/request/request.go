package request

type UserRequest struct {
	Fullname    string `url:"fullname" validate:"required,gte=0,lte=256" query:"fullname" param:"fullname" json:"fullname"`
	Email       string `url:"email" validate:"required,email,gte=0,lte=256" query:"Email" param:"Email" json:"Email"`
	Username    string `url:"username" validate:"required,gte=0,lte=256" query:"Username" param:"Username" json:"Username"`
	Password    string `url:"password" query:"password" param:"password" json:"password" validate:"required"`
	NewPassword string `url:"new_password" query:"new_password" param:"new_password" json:"new_password"`
	Disabled    bool   `url:"disabled" query:"disabled" param:"disabled" json:"disabled"`
}
