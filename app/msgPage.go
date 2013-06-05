package app

import (
	"app/core"
	"appengine"
	"net/http"
)

var sharedMsgWaitress = newMsgWaitress()

// single article page
type MsgWaitress struct {
	*core.SimpleWaitress
}

func newMsgWaitress() *MsgWaitress {
	self := &MsgWaitress{core.NewSimpleWaitress()}

	return self
}

func (self *MsgWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// setup build cube
	buildCube := core.GetBuildCube(self.GetApp(), r)
	// buildCube["PageCube"]["Articles"] = articles
	pageCube := buildCube["PageCube"]

	if pageCube["MsgType"] == "warning" {
		sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "warningTpl")
	} else if pageCube["MsgType"] == "success" {
		sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "successTpl")
	}

}

func (self *MsgWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	self.Get(w, r, c)
}

// func (self *MsgWaitress) PrintSuccessMsg(msg string, w http.ResponseWriter, r *http.Request, c appengine.Context) {
// 	self.GetPageCube()["Msg"] = msg
// 	sharedTempler.BuildPage(c, w, self, "baseTpl", "warningTpl")
// }

// func (self *MsgWaitress) PrintWarningMsg(msg string, w http.ResponseWriter, r *http.Request, c appengine.Context) {
// 	self.GetPageCube()["Msg"] = msg
// 	sharedTempler.BuildPage(c, w, self, "baseTpl", "warningTpl")
// }
