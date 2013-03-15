package helloword

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func init() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/add", addArticle)
	http.HandleFunc("/article", articlePage)
	http.HandleFunc("/aboutme", aboutmePage)
	http.HandleFunc("/hello", handler)
}

// this struct vill be store appengine
type Article struct {
	Title      string
	Content    []byte
	CreateTime time.Time
}

// simple hello world function
func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
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
	if u.String() != "zero.hero.lin" {
		fmt.Fprintf(w, "你不是管理员，请登陆管理员的google账户。")
		return
	}
	fmt.Fprintf(w, "Hello, %v!", u)
}

// home page handle function, only deal with the "GET" situaction
func homePage(w http.ResponseWriter, r *http.Request) {
	// get all articles
	articles := make([]Article, 0)
	c := appengine.NewContext(r)
	// sort by create time
	q := datastore.NewQuery("article").Order("-CreateTime")
	keys, err := q.GetAll(c, &articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// build page
	baseFile := "templates/base.tpl"
	homeFile := "templates/home.tpl"
	if t, err := template.ParseFiles(baseFile, homeFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		page := template.Must(t, err)
		// use articleMap to deal with the template
		articleMaps := modelsToMaps(keys, articles)
		if err := page.Execute(w, articleMaps); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// add article page handle function
func addArticle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
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
	if u.String() != "zero.hero.lin" {
		fmt.Fprintf(w, "你不是管理员，请登陆管理员的google账户。")
		return
	}

	if r.Method == "POST" {
		entityName := "article"

		// get info from page, and build a modle
		r.ParseForm()
		article := Article{
			Title:      r.Form.Get("title"),
			Content:    []byte(r.Form.Get("content")),
			CreateTime: time.Now(),
		}
		// save to datastore
		c := appengine.NewContext(r)
		_, err := datastore.Put(c, datastore.NewIncompleteKey(c, entityName, nil), &article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprint(w, "success!")
		}

		return

	} else {
		// build page
		baseFile := "templates/base.tpl"
		addFile := "templates/add_article.tpl"
		if t, err := template.ParseFiles(baseFile, addFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			page := template.Must(t, err)

			if err := page.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		}
	}
}

// single article page
func articlePage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	oops(err, w)

	id := r.Form.Get("id")
	key, err := datastore.DecodeKey(id)
	oops(err, w)
	article := &Article{}
	c := appengine.NewContext(r)
	err = datastore.Get(c, key, article)
	oops(err, w)
	articleMap := modelToMap(key, article)

	// build page
	baseFile := "templates/base.tpl"
	articleFile := "templates/article.tpl"
	if t, err := template.ParseFiles(baseFile, articleFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		page := template.Must(t, err)
		// use articleMap to deal with the template
		if err := page.Execute(w, articleMap); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// my resume page
func aboutmePage(w http.ResponseWriter, r *http.Request) {
	buildPage(w, "templates/aboutme.tpl", nil)
}

// this function helps to convert the article struct to article map
func modelsToMaps(keys []*datastore.Key, articles []Article) (maps []map[string]string) {
	maps = make([]map[string]string, 0)
	for i, article := range articles {
		key := keys[i]
		articleMap := modelToMap(key, &article)
		maps = append(maps, articleMap)
	}

	return maps
}

//
func modelToMap(key *datastore.Key, article *Article) (articleMap map[string]string) {
	articleMap = make(map[string]string)
	articleMap["Id"] = key.Encode()
	articleMap["Title"] = article.Title
	articleMap["Content"] = string(article.Content)
	createTime := article.CreateTime
	articleMap["CreateTime"] = createTime.String()[0:16]
	return articleMap
}

func oops(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func buildPage(w http.ResponseWriter, targetFile string, context interface{}) {
	// build page
	baseFile := "templates/base.tpl"
	if t, err := template.ParseFiles(baseFile, targetFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		page := template.Must(t, err)
		// use articleMap to deal with the template
		if err := page.Execute(w, context); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
