package app

import (
	"app/core"
	"app/ds"
	M "app/model"
	"appengine"
	"net/http"
	"time"
)

var sharedAddAFloatWaitress = newAddAFloatWaitress()

type AddAFloatWaitress struct {
	*core.SimpleWaitress
}

func newAddAFloatWaitress() *AddAFloatWaitress {
	self := &AddAFloatWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *AddAFloatWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {

	// get info from page, and build a modle
	r.ParseForm()
	aFloat := M.AFloat{
		Title:      r.Form.Get("title"),
		Content:    r.Form.Get("content"),
		CreateTime: time.Now(),
		ID:         "",
	}
	// c.Warningf("article is: ", article)
	// save to datastore
	err := ds.AddAFloat(c, &aFloat)
	oops(err, w)
	// self.InfoCube["AFloats"] = append(self.InfoCube["AFloats"].([]M.AFloat), aFloat)
	printSuccessPage(self.GetApp(), w, r, "添加成功！")

	return
}

// get function
func (self *AddAFloatWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// build page
	buildCube := core.GetBuildCube(self.GetApp(), r)
	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "addAFTpl")
}
