package app

import (
	"app/core"
	"net/http"
)

// print error on the page
func oops(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// print error, version 2
func oops2(app core.Application, w http.ResponseWriter, r *http.Request, err error) {
	printWarningPage(app, w, r, err.Error())
}

// print warning message
func printWarningPage(app core.Application, w http.ResponseWriter, r *http.Request, msg string) {
	buildPage := core.GetBuildCube(app, r)
	buildPage["PageCube"]["Msg"] = msg
	buildPage["PageCube"]["MsgType"] = "warning"
	app.ForwardByWaitress(sharedMsgWaitress, w, r)
	return
}

// print success message
func printSuccessPage(app core.Application, w http.ResponseWriter, r *http.Request, msg string) {
	buildPage := core.GetBuildCube(app, r)
	buildPage["PageCube"]["Msg"] = msg
	buildPage["PageCube"]["MsgType"] = "success"
	app.ForwardByWaitress(sharedMsgWaitress, w, r)
	return
}
