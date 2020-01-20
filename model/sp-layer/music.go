package splayer

import (
	dblayer "github.com/Myriad-Dreamin/blog-backend-v2/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Music = dblayer.Music

type MusicDB struct {
	dblayer.MusicDB
}

func NewMusicDB(m module.Module) (*MusicDB, error) {
	cdb, err := dblayer.NewMusicDB(m)
	if err != nil {
		return nil, err
	}
	db := new(MusicDB)
	db.MusicDB = *cdb
	return db, nil
}

func GetMusicDB(m module.Module) (*MusicDB, error) {
	cdb, err := dblayer.GetMusicDB(m)
	if err != nil {
		return nil, err
	}
	db := new(MusicDB)
	db.MusicDB = *cdb
	return db, nil
}

func (s *Provider) MusicDB() *MusicDB {
	return s.musicDB
}
