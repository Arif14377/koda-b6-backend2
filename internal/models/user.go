package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 2 di bawah ini adalah DTO, akan digunakan di ShouldBindJSON()
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
