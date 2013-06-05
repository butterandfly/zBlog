package app

import (
	"app/core"
	"appengine"
	"net/http"
)

type StaticWaitress struct {
	*core.SimpleWaitress
	TemplNames []string
}

func NewStaticWaitressByTempl(templNames []string) *StaticWaitress {
	self := &StaticWaitress{}
	self.SimpleWaitress = core.NewSimpleWaitress()
	self.TemplNames = templNames

	return self
}

// 
func (self *StaticWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// build page
	sharedTempler.BuildPage(c, w, self, self.TemplNames...)
}
