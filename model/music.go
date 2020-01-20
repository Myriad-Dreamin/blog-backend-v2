package model

import (
	splayer "github.com/Myriad-Dreamin/blog-backend-v2/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Music = splayer.Music
type MusicDB = splayer.MusicDB

func NewMusicDB(m module.Module) (*MusicDB, error) {
	return splayer.NewMusicDB(m)
}

func GetMusicDB(m module.Module) (*MusicDB, error) {
	return splayer.GetMusicDB(m)
}
