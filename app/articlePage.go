package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

var sharedArticleWaitress = newArticleWaitress()

// single article page
type ArticleWaitress struct {
	*core.SimpleWaitress
}

func newArticleWaitress() *ArticleWaitress {
	self := &ArticleWaitress{core.NewSimpleWaitress()}

	return self
}

// 
func (self *ArticleWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	err := r.ParseForm()
	oops(err, w)

	id := r.Form.Get("id")
	article, err := ds.GetArticleById(id, c)
	oops(err, w)
	// setup build cube
	buildCube := core.GetBuildCube(self.GetApp(), r)
	buildCube["PageCube"]["Article"] = article

	// build page
	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "artTpl")
}
