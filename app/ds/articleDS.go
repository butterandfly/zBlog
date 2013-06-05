package ds

import (
	M "app/model"
	"appengine"
	"appengine/datastore"
)

func GetAllArticles(c appengine.Context) ([]M.Article, error) {
	articleEntitys := make([]M.ArticleEntity, 0)
	// sort by create time
	q := datastore.NewQuery("ArticleEntity").Order("-CreateTime")
	keys, err := q.GetAll(c, &articleEntitys)
	if err != nil {
		return nil, err
	}
	articles := make([]M.Article, len(keys))
	for i, key := range keys {
		article := dsHelper.Entity2Model(key, &articleEntitys[i]).(M.Article)
		articles[i] = article
	}

	return articles, nil
}

func GetArticleById(id string, c appengine.Context) (*M.Article, error) {
	// get key
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, err
	}

	// get entity
	articleEntity := &M.ArticleEntity{}
	err = datastore.Get(c, key, articleEntity)
	if err != nil {
		return nil, err
	}

	// get model
	article := dsHelper.Entity2Model(key, articleEntity).(M.Article)

	return &article, nil
}

func DeleteArticleById(id string, c appengine.Context) error {
	// decode string
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	// delete
	err = datastore.Delete(c, key)
	return err
}

func AddArticle(article *M.Article, c appengine.Context) (err error) {
	// get entity
	entityName := "ArticleEntity"
	_, entity := dsHelper.Model2Entity(article)
	articleEntity := entity.(M.ArticleEntity)

	// save
	_, err = datastore.Put(c, datastore.NewIncompleteKey(c, entityName, nil), &articleEntity)

	return err
}

func EditArticle(article *M.Article, c appengine.Context) error {
	// get key and entity
	key, entity := dsHelper.Model2Entity(article)
	articleEntity := entity.(M.ArticleEntity)

	// save
	_, err := datastore.Put(c, key, &articleEntity)

	return err
}
