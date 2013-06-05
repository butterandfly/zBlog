package app

import (
	"app/core"
	"appengine"
	"net/http"
)

type InitNavFilter struct {
	*core.SimpleFilter
}

func newInitNavFilter(app core.Application) *InitNavFilter {
	self := &InitNavFilter{core.NewSimpleFilter(app)}

	return self
}

func (self *InitNavFilter) Filte(w http.ResponseWriter, r *http.Request, c appengine.Context) bool {
	c.Debugf("initnavfilter filte function")

	pageCube := self.GetApp().GetPageCube(r)
	pageCube["ActiveNav"] = "home"

	return true
}
