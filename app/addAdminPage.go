package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"net/http"
)

var sharedAddAdminWaitress = newAddAdminWaitress()

type AddAdminWaitress struct {
	*core.SimpleWaitress
}

func newAddAdminWaitress() *AddAdminWaitress {
	self := &AddAdminWaitress{core.NewSimpleWaitress()}

	return self
}

// over write
func (self *AddAdminWaitress) Post(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// add a admin
	r.ParseForm()
	if acount := r.Form.Get("acount"); acount != "" {
		err := ds.AddAdmin(c, acount)
		if err != nil {
			oops2(self.GetApp(), w, r, err)
		} else {
			printSuccessPage(self.GetApp(), w, r, "添加管理员成功！")
		}
	} else {
		printWarningPage(self.GetApp(), w, r, "添加管理员失败！")
	}
	return
}

// over write
func (self *AddAdminWaitress) Get(w http.ResponseWriter, r *http.Request, c appengine.Context) {
	// * check if has a admin
	// get the admin info
	keysCount, err := ds.AdminCount(c)
	oops(err, w)
	// check the count of keys
	if keysCount > 0 {
		// warning that admin has been exist
		printWarningPage(self.GetApp(), w, r, "管理员已经存在！")
		return
	}

	// * set the admin
	// build a form
	buildCube := core.GetBuildCube(self.GetApp(), r)
	err = sharedTempler.BuildPage(c, w, buildCube, "baseTpl", "addAdminTpl")
	oops(err, w)
	return
}
