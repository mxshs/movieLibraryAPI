package adaptermodels

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResetPassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
