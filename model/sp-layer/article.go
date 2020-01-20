package splayer

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	dblayer "github.com/Myriad-Dreamin/blog-backend-v2/model/db-layer"
)

type Article = dblayer.Article

type ArticleDB struct {
	dblayer.ArticleDB
}

func NewArticleDB(m module.Module) (*ArticleDB, error) {
	cdb, err := dblayer.NewArticleDB(m)
	if err != nil {
		return nil, err
	}
	db := new(ArticleDB)
	db.ArticleDB = *cdb
	return db, nil
}

func GetArticleDB(m module.Module) (*ArticleDB, error) {
	cdb, err := dblayer.GetArticleDB(m)
	if err != nil {
		return nil, err
	}
	db := new(ArticleDB)
	db.ArticleDB = *cdb
	return db, nil
}

func (s *Provider) ArticleDB() *ArticleDB {
	return s.articleDB
}
