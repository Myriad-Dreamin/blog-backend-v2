package tests

import (
	"fmt"
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	"testing"
)

func testUserGet(t *testing.T) {
	srv := srv.Context(t).AssertNoError(true)
	resp := srv.Get("/v1/user/1")
	reply := srv.DecodeJSON(resp.Body(), new(control.GetUserReply)).(*control.GetUserReply)
	fmt.Println(reply)
}
