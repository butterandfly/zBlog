package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

var sharedEditArticleWaitress = newEditArticleWaitress()

type EditArticleWaitress struct {
	*core.SimpleWaitress
}

func newEditArticleWaitress() *EditArticleWaitress {
	self := &EditArticleWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *EditArticleWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get title, content, and id
	r.ParseForm()
	articleId := r.FormValue("id")
	articleTitle := r.FormValue("title")
	articleContent := r.FormValue("content")
	// get article
	article, err := ds.GetArticleById(articleId, c)
	oops(err, w)
	article.Title = articleTitle
	article.Content = articleContent
	// edit
	err = ds.EditArticle(article, c)
	oops(err, w)
	// print success message
	printSuccessPage(self.GetApp(), w, r, "修改成功！")
}

// get function
func (self *EditArticleWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get id
	r.ParseForm()
	articleId := r.Form.Get("id")
	// get article
	article, err := ds.GetArticleById(articleId, c)
	oops(err, w)
	// setup build cube
	buildCube := core.GetBuildCube(self.GetApp(), r)
	buildCube["PageCube"]["Article"] = article
	// build page
	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "editArtTpl")
}
