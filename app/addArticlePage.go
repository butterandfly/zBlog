package app

import (
	"app/core"
	"app/ds"
	M "app/model"
	"appengine"
	"net/http"
	"time"
)

var sharedAddArticleWaitress = newAddArticleWaitress()

type AddArticleWaitress struct {
	*core.SimpleWaitress
}

func newAddArticleWaitress() *AddArticleWaitress {
	self := &AddArticleWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *AddArticleWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {

	// get info from page, and build a modle
	r.ParseForm()
	article := &M.Article{
		Title:      r.Form.Get("title"),
		Content:    r.Form.Get("content"),
		CreateTime: time.Now(),
		ID:         "",
	}
	// c.Warningf("article is: ", article)
	// save to datastore
	err := ds.AddArticle(article, c)
	oops(err, w)
	printSuccessPage(self.GetApp(), w, r, "添加成功！")

	return
}

// get function
func (self *AddArticleWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// build page
	buildCube := core.GetBuildCube(self.GetApp(), r)
	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "addArtTpl")
}
