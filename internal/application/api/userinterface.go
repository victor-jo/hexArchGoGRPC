package api

type User interface {
	Join(username string, password string)
	Find(id uint64)
}
