
package control

import (
    "github.com/Myriad-Dreamin/minimum-lib/controller"
    "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
    "github.com/Myriad-Dreamin/blog-backend-v2/model/db-layer"

)

var _ controller.MContext


type MusicService interface {
    MusicServiceSignatureXXX() interface{}
    ListMusics(c controller.MContext)
    PostMusic(c controller.MContext)
    InspectMusic(c controller.MContext)
    GetMusic(c controller.MContext)
    PutMusic(c controller.MContext)
    Delete(c controller.MContext)

}
type ListMusicsRequest = gorm_crud_dao.Filter

type ListMusicsReply struct {
    Code int `json:"code" form:"code"`
    Musics []dblayer.Music `json:"musics" form:"musics"`
}

type PostMusicRequest struct {
    RecommendType int `json:"recommend_type" form:"recommend_type"`
    Category string `json:"category" form:"category"`
    Title string `json:"title" form:"title"`
    Artist string `json:"artist" form:"artist"`
    TrackName string `json:"track_name" form:"track_name"`
    ReferenceID string `json:"reference_i_d" form:"reference_i_d"`
    Comment string `json:"comment" form:"comment"`
}

type PostMusicReply struct {
    Code int `json:"code" form:"code"`
    Music *dblayer.Music `json:"music" form:"music"`
}

type InspectMusicReply struct {
    Code int `json:"code" form:"code"`
    Music *dblayer.Music `json:"music" form:"music"`
}

type GetMusicReply struct {
    Code int `json:"code" form:"code"`
    Music *dblayer.Music `json:"music" form:"music"`
}

type PutMusicRequest struct {

}
func PSerializeListMusicsReply(_code int, _musics []dblayer.Music) *ListMusicsReply {

    return &ListMusicsReply{
        Code: _code,
        Musics: _musics,
    }
}
func SerializeListMusicsReply(_code int, _musics []dblayer.Music) ListMusicsReply {

    return ListMusicsReply{
        Code: _code,
        Musics: _musics,
    }
}
func _packSerializeListMusicsReply(_code int, _musics []dblayer.Music) ListMusicsReply {

    return ListMusicsReply{
        Code: _code,
        Musics: _musics,
    }
}
func PackSerializeListMusicsReply(_code []int, _musics [][]dblayer.Music) (pack []ListMusicsReply) {
	for i := range _code {
		pack = append(pack, _packSerializeListMusicsReply(_code[i], _musics[i]))
	}
	return
}
func PSerializePostMusicRequest(music *dblayer.Music) *PostMusicRequest {

    return &PostMusicRequest{
        RecommendType: music.RecommendType,
        Category: music.Category,
        Title: music.Title,
        Artist: music.Artist,
        TrackName: music.TrackName,
        ReferenceID: music.ReferenceID,
        Comment: music.Comment,
    }
}
func SerializePostMusicRequest(music *dblayer.Music) PostMusicRequest {

    return PostMusicRequest{
        RecommendType: music.RecommendType,
        Category: music.Category,
        Title: music.Title,
        Artist: music.Artist,
        TrackName: music.TrackName,
        ReferenceID: music.ReferenceID,
        Comment: music.Comment,
    }
}
func _packSerializePostMusicRequest(music *dblayer.Music) PostMusicRequest {

    return PostMusicRequest{
        RecommendType: music.RecommendType,
        Category: music.Category,
        Title: music.Title,
        Artist: music.Artist,
        TrackName: music.TrackName,
        ReferenceID: music.ReferenceID,
        Comment: music.Comment,
    }
}
func PackSerializePostMusicRequest(music []*dblayer.Music) (pack []PostMusicRequest) {
	for i := range music {
		pack = append(pack, _packSerializePostMusicRequest(music[i]))
	}
	return
}
func PSerializePostMusicReply(_code int, _music *dblayer.Music) *PostMusicReply {

    return &PostMusicReply{
        Code: _code,
        Music: _music,
    }
}
func SerializePostMusicReply(_code int, _music *dblayer.Music) PostMusicReply {

    return PostMusicReply{
        Code: _code,
        Music: _music,
    }
}
func _packSerializePostMusicReply(_code int, _music *dblayer.Music) PostMusicReply {

    return PostMusicReply{
        Code: _code,
        Music: _music,
    }
}
func PackSerializePostMusicReply(_code []int, _music []*dblayer.Music) (pack []PostMusicReply) {
	for i := range _code {
		pack = append(pack, _packSerializePostMusicReply(_code[i], _music[i]))
	}
	return
}
func PSerializeInspectMusicReply(_code int, _music *dblayer.Music) *InspectMusicReply {

    return &InspectMusicReply{
        Code: _code,
        Music: _music,
    }
}
func SerializeInspectMusicReply(_code int, _music *dblayer.Music) InspectMusicReply {

    return InspectMusicReply{
        Code: _code,
        Music: _music,
    }
}
func _packSerializeInspectMusicReply(_code int, _music *dblayer.Music) InspectMusicReply {

    return InspectMusicReply{
        Code: _code,
        Music: _music,
    }
}
func PackSerializeInspectMusicReply(_code []int, _music []*dblayer.Music) (pack []InspectMusicReply) {
	for i := range _code {
		pack = append(pack, _packSerializeInspectMusicReply(_code[i], _music[i]))
	}
	return
}
func PSerializeGetMusicReply(_code int, _music *dblayer.Music) *GetMusicReply {

    return &GetMusicReply{
        Code: _code,
        Music: _music,
    }
}
func SerializeGetMusicReply(_code int, _music *dblayer.Music) GetMusicReply {

    return GetMusicReply{
        Code: _code,
        Music: _music,
    }
}
func _packSerializeGetMusicReply(_code int, _music *dblayer.Music) GetMusicReply {

    return GetMusicReply{
        Code: _code,
        Music: _music,
    }
}
func PackSerializeGetMusicReply(_code []int, _music []*dblayer.Music) (pack []GetMusicReply) {
	for i := range _code {
		pack = append(pack, _packSerializeGetMusicReply(_code[i], _music[i]))
	}
	return
}
func PSerializePutMusicRequest() *PutMusicRequest {

    return &PutMusicRequest{

    }
}
func SerializePutMusicRequest() PutMusicRequest {

    return PutMusicRequest{

    }
}
func _packSerializePutMusicRequest() PutMusicRequest {

    return PutMusicRequest{

    }
}
func PackSerializePutMusicRequest() (pack []PutMusicRequest) {
	return
}
