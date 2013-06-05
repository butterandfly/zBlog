package ds

import (
	M "app/model"
	"appengine"
	"appengine/datastore"
)

func AddAFloat(c appengine.Context, aFloat *M.AFloat) error {
	entityName := "AFloatEntity"
	_, entity := dsHelper.Model2Entity(aFloat)
	aFloatEntity := entity.(M.AFloatEntity)
	_, err := dsHelper.addEntity(c, entityName, &aFloatEntity)
	return err
}

func GetAllAFloat(c appengine.Context) (aFloats []M.AFloat, err error) {
	entityName := "AFloatEntity"
	entitySlice := make([]M.AFloatEntity, 0)
	keys, err := dsHelper.getAllSortedByCreateTime(c, entityName, &entitySlice)
	if err != nil {
		c.Errorf("error in getallafloat function: ", err)
		return nil, err
	}

	aFloats = []M.AFloat{}
	dsHelper.Entitys2Models(c, keys, entitySlice, &aFloats)
	return aFloats, nil
}

func DeleteAFloatById(c appengine.Context, id string) (err error) {
	return dsHelper.DeleteById(c, id)
}

func GetAFloatById(c appengine.Context, id string) (aFloatP *M.AFloat, err error) {
	aFloatEntity := M.AFloatEntity{}
	key, err := dsHelper.GetById(c, id, &aFloatEntity)
	if err != nil {
		return nil, err
	}

	aFloat := dsHelper.Entity2Model(key, &aFloatEntity).(M.AFloat)
	return &aFloat, nil
}

func EditAFloat(aFloat *M.AFloat, c appengine.Context) error {
	// get key and entity
	key, entity := dsHelper.Model2Entity(aFloat)
	aFloatEntity := entity.(M.AFloatEntity)

	// save
	_, err := datastore.Put(c, key, &aFloatEntity)

	return err
}
