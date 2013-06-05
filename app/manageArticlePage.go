package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

// teh page to manage articles
var sharedManageArticleWaitress = newManageArticleWaitress()

type ManageArticleWaitress struct {
	*core.SimpleWaitress
}

func newManageArticleWaitress() *ManageArticleWaitress {
	self := &ManageArticleWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *ManageArticleWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// * dealing deleting article
	// get paras
	err := r.ParseForm()
	oops(err, w)
	articleId := r.Form.Get("delete")
	// delete
	err = ds.DeleteArticleById(articleId, c)
	oops(err, w)
	// print success message
	printSuccessPage(self.GetApp(), w, r, "删除成功！")
}

// get function
func (self *ManageArticleWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get all articles
	articles, err := ds.GetAllArticles(c)
	oops(err, w)
	// setup build cube
	buildCube := core.GetBuildCube(self.GetApp(), r)
	buildCube["PageCube"]["Articles"] = articles

	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "manageArtTpl")
	return
}
