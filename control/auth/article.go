package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"strconv"
)

type articleEntity struct{}

func (articleEntity) CreateObj(articleID uint) string {
	return "article:" + strconv.Itoa(int(articleID))
}

func (articleEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "article", Action: "r/.*"}
}

func (articleEntity) AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ArticleEntity.CreateObj(aim), "r/.*")
}

func (articleEntity) RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ArticleEntity.CreateObj(aim), "r/.*")
}

func (articleEntity) HasReadPolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ArticleEntity.CreateObj(aim), "r/.*")
}

func (articleEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "article", Action: "w/.*"}
}

func (articleEntity) AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ArticleEntity.CreateObj(aim), "w/.*")
}

func (articleEntity) RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ArticleEntity.CreateObj(aim), "w/.*")
}

func (articleEntity) HasWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ArticleEntity.CreateObj(aim), "w/.*")
}

func (articleEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "article", Action: "w/s"}
}

func (articleEntity) AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ArticleEntity.CreateObj(aim), "w/s")
}

func (articleEntity) RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ArticleEntity.CreateObj(aim), "w/s")
}

func (articleEntity) HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ArticleEntity.CreateObj(aim), "w/s")
}
