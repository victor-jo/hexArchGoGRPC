package user

type User struct {
}

func New() *User {
	return &User{}
}

func (user User) Join(username string, password string) {
	// 도메인 비즈니스 로직 부
	return
}

func (user User) Find(id uint64) {
	// 도메인 비즈니스 로직 부
	return
}
