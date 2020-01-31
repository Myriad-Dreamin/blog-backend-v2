package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
)

type MusicCategories struct {
	artisan.VirtualService
	List    artisan.Category
	Post    artisan.Category
	Inspect artisan.Category
	IdGroup artisan.Category
}

func DescribeMusicService(base string) artisan.ProposingService {
	var musicModel = new(model.Music)
	var _musicModel = new(model.Music)
	svc := &MusicCategories{
		List: artisan.Ink().
			Path("music-list").
			Method(artisan.POST, "ListMusics",
				artisan.QT("ListMusicsRequest", model.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("musics", _musicModel)),
				),
			),
		Post: artisan.Ink().
			Path("music").
			Method(artisan.POST, "PostMusic",
				artisan.Request(
					artisan.SPsC(
						&musicModel.RecommendType, &musicModel.Category,
						&musicModel.Title, &musicModel.Artist,
						&musicModel.TrackName, &musicModel.ReferenceID,
						&musicModel.Comment,
					),
				),
				artisan.Reply(
					codeField,
					artisan.Param("music", &musicModel),
				),
			),
		Inspect: artisan.Ink().Path("music/:mid/inspect").
			Method(artisan.GET, "InspectMusic",
				artisan.Reply(
					codeField,
					artisan.Param("music", &musicModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("music/:mid").
			Method(artisan.GET, "GetMusic",
				artisan.Reply(
					codeField,
					artisan.Param("music", &musicModel),
				)).
			Method(artisan.PUT, "PutMusic",
				artisan.Request()).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("MusicService").Base(base).
		UseModel(artisan.Model(artisan.Name("music"), &musicModel))
	return svc
}
