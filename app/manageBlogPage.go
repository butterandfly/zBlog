package app

import (
	"app/core"
	"app/ds"
	M "app/model"
	"appengine"
	"net/http"
)

// the page to manage blog
var sharedManageBlogWaitress = newManageBlogWaitress()

type ManageBlogWaitress struct {
	*core.SimpleWaitress
}

func newManageBlogWaitress() *ManageBlogWaitress {
	self := &ManageBlogWaitress{core.NewSimpleWaitress()}

	return self
}

// post function
func (self *ManageBlogWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// get blog name, noRight
	r.ParseForm()
	blogName := r.FormValue("blogname")
	noRight := false
	if r.FormValue("noRight") == "1" {
		noRight = true
	}
	// set blog name
	blogCube, _ := ds.GetBlogCube(c)
	if blogCube == nil {
		// add a new one
		blogCube = &M.BlogCube{
			BlogName: blogName,
			NoRight:  noRight,
			ID:       "",
		}
		err := ds.AddBlogCube(blogCube, c)
		oops(err, w)
	} else {
		//! edit
		blogCube.BlogName = blogName
		blogCube.NoRight = noRight
		err := ds.EditBlogCube(blogCube, c)
		oops(err, w)
	}
	//
	self.GetApp().GetAppCube()["BlogCube"] = blogCube
	// print success
	printSuccessPage(self.GetApp(), w, r, "更改成功！")
}

// get function
func (self *ManageBlogWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// build page
	buildCube := core.GetBuildCube(self.GetApp(), r)
	sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "manageBlogTpl")
}
