package domain

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneHumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneHumber: phoneNumber,
	}
}

func NewUserUninitialized(
	fullname string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullname,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	// todo
}