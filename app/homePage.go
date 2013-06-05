package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

// "/home" page
var sharedHomeWaitress = newHomeWaitress()

type HomeWaitress struct {
	*core.SimpleWaitress
}

func newHomeWaitress() *HomeWaitress {
	self := &HomeWaitress{core.NewSimpleWaitress()}

	return self
}

// over write "Get" method
func (self *HomeWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get all articles
	articles, err := ds.GetAllArticles(c)
	oops(err, w)
	aFloats, err := ds.GetAllAFloat(c)
	oops(err, w)
	// setup
	buildCube := core.GetBuildCube(self.GetApp(), r)
	buildCube["PageCube"]["Articles"] = articles
	buildCube["PageCube"]["AFloats"] = aFloats
	// build page
	err = sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "homeTpl")
	oops(err, w)
}
