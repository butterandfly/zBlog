package app

import (
	"app/core"
	"app/ds"
	M "app/model"
	"appengine"
	"net/http"
)

type AppCubeFilter struct {
	*core.SimpleFilter
}

func newAppCubeFilter(app core.Application) *AppCubeFilter {
	self := &AppCubeFilter{core.NewSimpleFilter(app)}

	return self
}

func (self *AppCubeFilter) Filte(w http.ResponseWriter, r *http.Request, c appengine.Context) bool {
	c.Warningf("app cube filte")
	// (self.App.GetAppCube())["BlogName"] = "Core"
	if _, ok := (self.App.GetAppCube())["BlogCube"]; !ok {
		cube, _ := ds.GetBlogCube(c)
		if cube == nil {
			cube = &M.BlogCube{}
			cube.BlogName = "Blog"
			cube.NoRight = false
			ds.AddBlogCube(cube, c)
		}
		self.App.GetAppCube()["BlogCube"] = cube
	}

	return true
}
