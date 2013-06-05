package app

import (
	"app/core"
	"app/ds"
	M "app/model"
	"appengine"
	"appengine/datastore"
	"net/http"
	"time"
)

type Article struct {
	Content    []byte
	Title      string
	CreateTime time.Time
}

var sharedGetOldDataWaitress = newGetOldDataWaitress()

type GetOldDataWaitress struct {
	*core.SimpleWaitress
}

func newGetOldDataWaitress() *GetOldDataWaitress {
	self := &GetOldDataWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *GetOldDataWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {

}

// get function
func (self *GetOldDataWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// Put a test data
	// Delete this for the real gae runtime
	// testArticle := &Article{
	// 	Title:      "Old Data",
	// 	Content:    []byte("Now you see old dta"),
	// 	CreateTime: time.Now(),
	// }
	// datastore.Put(c, datastore.NewIncompleteKey(c, "article", nil), testArticle)

	// * Get all old article.
	articles := make([]Article, 0)
	q := datastore.NewQuery("article").Order("-CreateTime")
	q.GetAll(c, &articles)

	// * Add to current datastore.
	for _, article := range articles {
		articleModel := &M.Article{
			Title:      article.Title,
			Content:    string(article.Content),
			CreateTime: article.CreateTime,
			ID:         "",
		}
		err := ds.AddArticle(articleModel, c)
		oops(err, w)
	}
	printSuccessPage(self.GetApp(), w, r, "数据复制成功")
}
