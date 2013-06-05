package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

var sharedEditAFloatWaitress = newEditAFloatWaitress()

type EditAFloatWaitress struct {
	*core.SimpleWaitress
}

func newEditAFloatWaitress() *EditAFloatWaitress {
	self := &EditAFloatWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *EditAFloatWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get title, content, and id
	r.ParseForm()
	aFloatId := r.FormValue("id")
	aFloatTitle := r.FormValue("title")
	aFloatContent := r.FormValue("content")
	// get aFloat
	aFloat, err := ds.GetAFloatById(c, aFloatId)
	oops(err, w)
	aFloat.Title = aFloatTitle
	aFloat.Content = aFloatContent
	// edit
	err = ds.EditAFloat(aFloat, c)
	oops(err, w)
	// print success message
	printSuccessPage(self.GetApp(), w, r, "修改成功！")
}

// get function
func (self *EditAFloatWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get id
	r.ParseForm()
	aFloatId := r.Form.Get("id")
	// get article
	aFloat, err := ds.GetAFloatById(c, aFloatId)
	oops(err, w)
	// setup cube
	buildCube := core.GetBuildCube(self.GetApp(), r)
	buildCube["PageCube"]["AFloat"] = aFloat
	// build page
	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "editAFTpl")
}
