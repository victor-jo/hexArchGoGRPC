package ports

import "hex/internal/application/domain"

type UserDbPorts interface {
	CreateUser(username string, password string) (uint64, error)
	FindUser(id uint64) (domain.UserEntity, error)
}

type UserApiPorts interface {
	Join(username string, password string) (uint64, error)
	Find(id uint64) (domain.UserEntity, error)
}
