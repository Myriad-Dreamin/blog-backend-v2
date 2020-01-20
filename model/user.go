package model

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	splayer "github.com/Myriad-Dreamin/blog-backend-v2/model/sp-layer"
)

type User = splayer.User
type UserDB = splayer.UserDB

func NewUserDB(m module.Module) (*UserDB, error) {
	return splayer.NewUserDB(m)
}

func GetUserDB(m module.Module) (*UserDB, error) {
	return splayer.GetUserDB(m)
}
