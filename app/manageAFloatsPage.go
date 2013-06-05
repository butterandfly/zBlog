package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

// the page to manage AFloats
var sharedManageAFloatsWaitress = newManageAFloatsWaitress()

type ManageAFloatsWaitress struct {
	*core.SimpleWaitress
}

func newManageAFloatsWaitress() *ManageAFloatsWaitress {
	self := &ManageAFloatsWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *ManageAFloatsWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// * dealing deleting aFloat
	// get paras
	err := r.ParseForm()
	oops(err, w)
	aFloatId := r.Form.Get("delete")
	// delete
	err = ds.DeleteAFloatById(c, aFloatId)
	oops(err, w)
	// print success message
	printSuccessPage(self.GetApp(), w, r, "删除成功！")
}

// get function
func (self *ManageAFloatsWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get all afloat
	aFloats, err := ds.GetAllAFloat(c)
	oops(err, w)
	// setup build cube
	buildCube := core.GetBuildCube(self.GetApp(), r)
	buildCube["PageCube"]["AFloats"] = aFloats

	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "manageAFsTpl")
	return
}
