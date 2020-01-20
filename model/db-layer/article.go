package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	articleTraits Traits
)

func injectArticleTraits() error {
	articleTraits = NewTraits(Article{})
	return nil
}

func wrapToArticle(article interface{}, err error) (*Article, error) {
	if article == nil {
		return nil, err
	}
	return article.(*Article), err
}

type Article struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`

	Title string `dorm:"title" gorm:"title;not_null"`
	Intro string `dorm:"intro" gorm:"intro;type:text;not_null"`
	Category string `dorm:"category" gorm:"category;not_null"`
	FilePath string `dorm:"file_path" gorm:"file_path;not_null"`
}

// TableName specification
func (Article) TableName() string {
	return "article"
}

func (Article) migrate() error {
	return articleTraits.Migrate()
}

func (d Article) GetID() uint {
	return d.ID
}

func (d *Article) Create() (int64, error) {
	return articleTraits.Create(d)
}

func (d *Article) Update() (int64, error) {
	return articleTraits.Update(d)
}

func (d *Article) UpdateFields(fields []string) (int64, error) {
	return articleTraits.UpdateFields(d, fields)
}

func (d *Article) UpdateFields_(db *dorm.DB, fields []string) (int64, error) {
	return articleTraits.UpdateFields_(db, d, fields)
}

func (d *Article) UpdateFields__(db dorm.SQLCommon, fields []string) (int64, error) {
	return articleTraits.UpdateFields__(db, d, fields)
}

func (d *Article) Delete() (int64, error) {
	return articleTraits.Delete(d)
}

type ArticleDB struct{}

func NewArticleDB(_ module.Module) (*ArticleDB, error) {
	return new(ArticleDB), nil
}

func GetArticleDB(_ module.Module) (*ArticleDB, error) {
	return new(ArticleDB), nil
}

func (articleDB *ArticleDB) Filter(f *Filter) (user []Article, err error) {
	err = articleTraits.Filter(f, &user)
	return
}

func (articleDB *ArticleDB) FilterI(f interface{}) (interface{}, error) {
	return articleDB.Filter(f.(*Filter))
}

func (articleDB *ArticleDB) ID(id uint) (article *Article, err error) {
	return wrapToArticle(articleTraits.ID(id))
}

func (articleDB *ArticleDB) ID_(db *gorm.DB, id uint) (goods *Article, err error) {
	return wrapToArticle(articleTraits.ID_(db, id))
}

type ArticleQuery struct {
	db *gorm.DB
}

func (articleDB *ArticleDB) QueryChain() *ArticleQuery {
	return &ArticleQuery{db: p.GormDB}
}

func (articleDB *ArticleQuery) Order(order string, reorder ...bool) *ArticleQuery {
	articleDB.db = articleDB.db.Order(order, reorder...)
	return articleDB
}

func (articleDB *ArticleQuery) Page(page, pageSize int) *ArticleQuery {
	articleDB.db = articleDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return articleDB
}
func (articleDB *ArticleQuery) BeforeID(id uint) *ArticleQuery {
	articleDB.db = articleDB.db.Where("id <= ?", id)
	return articleDB
}

func (articleDB *ArticleQuery) Preload() *ArticleQuery {
	articleDB.db = articleDB.db.Set("gorm:auto_preload", true)
	return articleDB
}

func (articleDB *ArticleQuery) Query() (articles []Article, err error) {
	err = articleDB.db.Find(&articles).Error
	return
}

func (articleDB *ArticleQuery) Scan(desc interface{}) (err error) {
	err = articleDB.db.Scan(desc).Error
	return
}
