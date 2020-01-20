package service

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	articleservice "github.com/Myriad-Dreamin/blog-backend-v2/service/article"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

// go:generate go run github.com/Myriad-Dreamin/minimum-lib/code-gen/test-gen -source ./ -destination ../../test/

type ArticleService = control.ArticleService

func NewArticleService(m module.Module) (ArticleService, error) {
	return articleservice.NewService(m)
}

func (s *Provider) ArticleService() ArticleService {
	return s.articleService
}
