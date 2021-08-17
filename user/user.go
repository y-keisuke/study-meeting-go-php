package user

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type Users []*User

var err error

func (u *User) Resister() error {
	//　登録処理
	ok := true

	if !ok {
		return err
	}

	return nil
}

func Fetch() Users {
	users := []*User{
		{
			"Suzuki",
			"Ichiro",
			"Tokyo",
			"ichiro@example.com",
		},
		{
			"Sato",
			"Taro",
			"Japan",
			"taro@example.com",
		},
	}

	return users
}