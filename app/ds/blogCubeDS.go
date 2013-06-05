package ds

import (
	M "app/model"
	"appengine"
	"appengine/datastore"
)

// type BlogCubeDS struct {
// 	BlogApp map[string]interface{}
// }

// func (self *BlogCubeDS) getBlogCube(c appengine.Context) (*model.BlogCube, error) {
func GetBlogCube(c appengine.Context) (*M.BlogCube, error) {
	// blogApp := BlogApp
	// blogApp := self.BlogApp
	// if blogApp.BCube != nil {
	// 	return blogApp.BCube, nil
	// }

	entitys := make([]M.BlogCubeEntity, 0)
	q := datastore.NewQuery("BlogCubeEntity").Limit(1)
	keys, err := q.GetAll(c, &entitys)
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, nil
	}
	key := keys[0]
	entity := &(entitys[0])
	cube := dsHelper.Entity2Model(key, entity).(M.BlogCube)

	// blogApp.BCube = &cube
	return &cube, nil
}

func AddBlogCube(model *M.BlogCube, c appengine.Context) (err error) {
	c.Warningf("addBlogCube function")
	entity := &(M.BlogCubeEntity{})
	entity.BlogName = model.BlogName
	entityName := "BlogCubeEntity"
	_, err = datastore.Put(c, datastore.NewIncompleteKey(c, entityName, nil), entity)
	return err
}

func EditBlogCube(model *M.BlogCube, c appengine.Context) (err error) {
	key, entity := dsHelper.Model2Entity(model)
	blogCubeEntity := entity.(M.BlogCubeEntity)
	c.Warningf("The entity is %v", blogCubeEntity)
	_, err = datastore.Put(c, key, &blogCubeEntity)
	return err
}
