package app

import (
	"app/core"
	"app/ds"
	"appengine"
	"appengine/user"
	"fmt"
	"net/http"
)

type AdminFilter struct {
	*core.SimpleFilter
}

func newAdminFilter(app core.Application) *AdminFilter {
	self := &AdminFilter{core.NewSimpleFilter(app)}

	return self
}

func (self *AdminFilter) Filte(w http.ResponseWriter, r *http.Request, c appengine.Context) (isValid bool) {
	c.Warningf("admin filter .........")
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	valid, err := isAdmin(c, u.String())
	oops(err, w)
	if !valid {
		url, _ := user.LogoutURL(c, r.URL.String())
		printString := fmt.Sprintf("当前google账户(%s)不是管理员！(<a href=\"%s\">退出此账号</a>)", u.String(), url)
		printWarningPage(self.GetApp(), w, r, printString)
		return false
	}
	return true
}

func isAdmin(c appengine.Context, user string) (answer bool, err error) {
	admins, err := ds.GetAllAdmin(c)

	if err != nil {
		return false, err
	}
	for _, admin := range admins {
		if user == admin.GoogleAcount {
			return true, nil
		}
	}
	return false, nil
}
