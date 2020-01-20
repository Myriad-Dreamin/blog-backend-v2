package model

import (
	splayer "github.com/Myriad-Dreamin/blog-backend-v2/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Article = splayer.Article
type ArticleDB = splayer.ArticleDB

func NewArticleDB(m module.Module) (*ArticleDB, error) {
	return splayer.NewArticleDB(m)
}

func GetArticleDB(m module.Module) (*ArticleDB, error) {
	return splayer.GetArticleDB(m)
}
