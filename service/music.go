package service

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	musicservice "github.com/Myriad-Dreamin/blog-backend-v2/service/music"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type MusicService = control.MusicService

func NewMusicService(m module.Module) (MusicService, error) {
	return musicservice.NewService(m)
}

func (s *Provider) MusicService() MusicService {
	return s.musicService
}
