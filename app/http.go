package app

import (
	"app/core"
	"text/template"
	"time"
)

// address-waitress map
var wtsMap = map[string]core.Waitress{
	"/":              sharedHomeWaitress,
	"/article":       sharedArticleWaitress,
	"/manageblog":    sharedManageBlogWaitress,
	"/addadmin":      sharedAddAdminWaitress,
	"/addarticle":    sharedAddArticleWaitress,
	"/managearticle": sharedManageArticleWaitress,
	"/editarticle":   sharedEditArticleWaitress,
	"/manageaf":      sharedManageAFloatsWaitress,
	"/addaf":         sharedAddAFloatWaitress,
	"/editaf":        sharedEditAFloatWaitress,
}

// following pages shoud add the admin filter
var adminPages = []core.Waitress{
	// blog cube
	sharedManageBlogWaitress,
	// article
	sharedManageArticleWaitress,
	sharedAddArticleWaitress,
	sharedEditArticleWaitress,
	// afloat
	sharedManageAFloatsWaitress,
	sharedAddAFloatWaitress,
	sharedManageAFloatsWaitress,
}

/*
type ZBlogApp struct {
	*core.SimpleApplication
}

func newZBlogApp() *ZBlogApp {
	self := &ZBlogApp{}
	self.SimpleApplication = core.NewSimpleApplication()

	return self
}
*/

func init() {

	// create a new application
	zApp := core.NewSimpleApplication()

	// * set up default filter
	appCubeFilter := newAppCubeFilter(zApp)
	initNavFilter := newInitNavFilter(zApp)
	zApp.AddDefaultFilters(appCubeFilter, initNavFilter)

	// * set up admin-filte page
	adminFilter := newAdminFilter(zApp)
	core.AddFilterToWaitresses(adminFilter, adminPages)

	zApp.StartApp(wtsMap)
}

/*
func addFilterToWaitresses(filter core.Filter, wtses []core.Waitress) {
	for _, wts := range wtses {
		filters := wts.GetFilters()
		filters = append(filters, filter)
		wts.SetFilters(filters)
	}
}
*/

var sharedTempler = core.NewTemplerWithJsonFile(funcMap, "templates/tplmap.json")

var funcMap = template.FuncMap{
	"strEqual":    strEqual,
	"niceTimeStr": niceTimeStr,
}

func strEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return false
}

func niceTimeStr(targetTime time.Time) string {
	return targetTime.String()[0:16]
}
