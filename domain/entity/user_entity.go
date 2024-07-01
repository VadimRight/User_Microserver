package entity

type User struct {
	Id         string
	Username   string
	Email      string
	Password   string
	IsVerified bool
	IsActive   bool
}

func (u User) GetId() string {
	return u.Id
}

func (u User) GetName() string {
	return u.Username
}

func (u User) GetEmail() string {
	return u.Email
}
