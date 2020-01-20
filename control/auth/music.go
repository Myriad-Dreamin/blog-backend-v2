package auth

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"strconv"
)

type musicEntity struct{}

func (musicEntity) CreateObj(musicID uint) string {
	return "music:" + strconv.Itoa(int(musicID))
}

func (musicEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "music", Action: "r/.*"}
}

func (musicEntity) AddReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, MusicEntity.CreateObj(aim), "r/.*")
}

func (musicEntity) RemoveReadPolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, MusicEntity.CreateObj(aim), "r/.*")
}

func (musicEntity) HasReadPolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, MusicEntity.CreateObj(aim), "r/.*")
}

func (musicEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "music", Action: "w/.*"}
}

func (musicEntity) AddWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, MusicEntity.CreateObj(aim), "w/.*")
}

func (musicEntity) RemoveWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, MusicEntity.CreateObj(aim), "w/.*")
}

func (musicEntity) HasWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, MusicEntity.CreateObj(aim), "w/.*")
}

func (musicEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "music", Action: "w/s"}
}

func (musicEntity) AddJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, MusicEntity.CreateObj(aim), "w/s")
}

func (musicEntity) RemoveJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, MusicEntity.CreateObj(aim), "w/s")
}

func (musicEntity) HasJustSimpleWritePolicy(e *Enforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, MusicEntity.CreateObj(aim), "w/s")
}
