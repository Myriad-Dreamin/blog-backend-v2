
package control

import (
    "github.com/Myriad-Dreamin/blog-backend-v2/model/db-layer"
    "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
    "github.com/Myriad-Dreamin/minimum-lib/controller"

)

var _ controller.MContext


type ArticleService interface {
    ArticleServiceSignatureXXX() interface{}
    ListArticles(c controller.MContext)
    PostArticle(c controller.MContext)
    InspectArticle(c controller.MContext)
    GetArticle(c controller.MContext)
    PutArticle(c controller.MContext)
    Delete(c controller.MContext)

}
type ListArticlesRequest = gorm_crud_dao.Filter

type ListArticlesReply struct {
    Code int `json:"code" form:"code"`
    Articles []dblayer.Article `json:"articles" form:"articles"`
}

type PostArticleRequest struct {

}

type PostArticleReply struct {
    Code int `json:"code" form:"code"`
    Article *dblayer.Article `json:"article" form:"article"`
}

type InspectArticleReply struct {
    Code int `json:"code" form:"code"`
    Article *dblayer.Article `json:"article" form:"article"`
}

type GetArticleReply struct {
    Code int `json:"code" form:"code"`
    Article *dblayer.Article `json:"article" form:"article"`
}

type PutArticleRequest struct {

}
func PSerializeListArticlesReply(_code int, _articles []dblayer.Article) *ListArticlesReply {

    return &ListArticlesReply{
        Code: _code,
        Articles: _articles,
    }
}
func SerializeListArticlesReply(_code int, _articles []dblayer.Article) ListArticlesReply {

    return ListArticlesReply{
        Code: _code,
        Articles: _articles,
    }
}
func _packSerializeListArticlesReply(_code int, _articles []dblayer.Article) ListArticlesReply {

    return ListArticlesReply{
        Code: _code,
        Articles: _articles,
    }
}
func PackSerializeListArticlesReply(_code []int, _articles [][]dblayer.Article) (pack []ListArticlesReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListArticlesReply(_code[i], _articles[i]))
	}
	return
}
func PSerializePostArticleRequest() *PostArticleRequest {

    return &PostArticleRequest{

    }
}
func SerializePostArticleRequest() PostArticleRequest {

    return PostArticleRequest{

    }
}
func _packSerializePostArticleRequest() PostArticleRequest {

    return PostArticleRequest{

    }
}
func PackSerializePostArticleRequest() (pack []PostArticleRequest) {
	return
}
func PSerializePostArticleReply(_code int, _article *dblayer.Article) *PostArticleReply {

    return &PostArticleReply{
        Code: _code,
        Article: _article,
    }
}
func SerializePostArticleReply(_code int, _article *dblayer.Article) PostArticleReply {

    return PostArticleReply{
        Code: _code,
        Article: _article,
    }
}
func _packSerializePostArticleReply(_code int, _article *dblayer.Article) PostArticleReply {

    return PostArticleReply{
        Code: _code,
        Article: _article,
    }
}
func PackSerializePostArticleReply(_code []int, _article []*dblayer.Article) (pack []PostArticleReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostArticleReply(_code[i], _article[i]))
	}
	return
}
func PSerializeInspectArticleReply(_code int, _article *dblayer.Article) *InspectArticleReply {

    return &InspectArticleReply{
        Code: _code,
        Article: _article,
    }
}
func SerializeInspectArticleReply(_code int, _article *dblayer.Article) InspectArticleReply {

    return InspectArticleReply{
        Code: _code,
        Article: _article,
    }
}
func _packSerializeInspectArticleReply(_code int, _article *dblayer.Article) InspectArticleReply {

    return InspectArticleReply{
        Code: _code,
        Article: _article,
    }
}
func PackSerializeInspectArticleReply(_code []int, _article []*dblayer.Article) (pack []InspectArticleReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectArticleReply(_code[i], _article[i]))
	}
	return
}
func PSerializeGetArticleReply(_code int, _article *dblayer.Article) *GetArticleReply {

    return &GetArticleReply{
        Code: _code,
        Article: _article,
    }
}
func SerializeGetArticleReply(_code int, _article *dblayer.Article) GetArticleReply {

    return GetArticleReply{
        Code: _code,
        Article: _article,
    }
}
func _packSerializeGetArticleReply(_code int, _article *dblayer.Article) GetArticleReply {

    return GetArticleReply{
        Code: _code,
        Article: _article,
    }
}
func PackSerializeGetArticleReply(_code []int, _article []*dblayer.Article) (pack []GetArticleReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetArticleReply(_code[i], _article[i]))
	}
	return
}
func PSerializePutArticleRequest() *PutArticleRequest {

    return &PutArticleRequest{

    }
}
func SerializePutArticleRequest() PutArticleRequest {

    return PutArticleRequest{

    }
}
func _packSerializePutArticleRequest() PutArticleRequest {

    return PutArticleRequest{

    }
}
func PackSerializePutArticleRequest() (pack []PutArticleRequest) {
	return
}
