package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	musicTraits Traits
)

func injectMusicTraits() error {
	musicTraits = NewTraits(Music{})
	return nil
}

func wrapToMusic(music interface{}, err error) (*Music, error) {
	if music == nil {
		return nil, err
	}
	return music.(*Music), err
}

type Music struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null" json:"id"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	RecommendType int    `dorm:"recommend_type" gorm:"column:recommend_type;not_null" json:"recommend_type"`
	Category      string `dorm:"category" gorm:"column:category;not_null" json:"category"`
	Title         string `dorm:"title" gorm:"column:title;not_null" json:"title"`
	Artist        string `dorm:"artist" gorm:"column:artist;not_null" json:"artist"`
	TrackName     string `dorm:"track_name" gorm:"column:track_name;not_null" json:"track_name"`
	ReferenceID   string `dorm:"reference_id" gorm:"column:reference_id;not_null" json:"reference_id"`
	Comment       string `dorm:"comment" gorm:"column:comment;type:text;not_null" json:"comment"`
}

// TableName specification
func (Music) TableName() string {
	return "music"
}

func (Music) migrate() error {
	return musicTraits.Migrate()
}

func (d Music) GetID() uint {
	return d.ID
}

func (d *Music) Create() (int64, error) {
	return musicTraits.Create(d)
}

func (d *Music) Update() (int64, error) {
	return musicTraits.Update(d)
}

func (d *Music) UpdateFields(fields []string) (int64, error) {
	return musicTraits.UpdateFields(d, fields)
}

func (d *Music) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return musicTraits.UpdateFields_(db, d, fields)
}

func (d *Music) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return musicTraits.UpdateFields__(db, d, fields)
}

func (d *Music) Delete() (int64, error) {
	return musicTraits.Delete(d)
}

type MusicDB struct{}

func NewMusicDB(_ module.Module) (*MusicDB, error) {
	return new(MusicDB), nil
}

func GetMusicDB(_ module.Module) (*MusicDB, error) {
	return new(MusicDB), nil
}

func (musicDB *MusicDB) Filter(f *Filter) (user []Music, err error) {
	err = musicTraits.Filter(f, &user)
	return
}

func (musicDB *MusicDB) FilterI(f interface{}) (interface{}, error) {
	return musicDB.Filter(f.(*Filter))
}

func (musicDB *MusicDB) ID(id uint) (music *Music, err error) {
	return wrapToMusic(musicTraits.ID(id))
}

func (musicDB *MusicDB) ID_(db *gorm.DB, id uint) (goods *Music, err error) {
	return wrapToMusic(musicTraits.ID_(db, id))
}

type MusicQuery struct {
	db *gorm.DB
}

func (musicDB *MusicDB) QueryChain() *MusicQuery {
	return &MusicQuery{db: p.GormDB}
}

func (musicDB *MusicQuery) Order(order string, reorder ...bool) *MusicQuery {
	musicDB.db = musicDB.db.Order(order, reorder...)
	return musicDB
}

func (musicDB *MusicQuery) Page(page, pageSize int) *MusicQuery {
	musicDB.db = musicDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return musicDB
}
func (musicDB *MusicQuery) BeforeID(id uint) *MusicQuery {
	musicDB.db = musicDB.db.Where("id <= ?", id)
	return musicDB
}

func (musicDB *MusicQuery) Preload() *MusicQuery {
	musicDB.db = musicDB.db.Set("gorm:auto_preload", true)
	return musicDB
}

func (musicDB *MusicQuery) Query() (musics []Music, err error) {
	err = musicDB.db.Find(&musics).Error
	return
}

func (musicDB *MusicQuery) Scan(desc interface{}) (err error) {
	err = musicDB.db.Scan(desc).Error
	return
}
