package ds

import (
	"app/model"
	"appengine"
	"appengine/datastore"
)

func AddAdmin(c appengine.Context, acount string) error {
	adminEntity := model.AdminEntity{acount}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "AdminEntity", nil), &adminEntity)
	return err
}

func GetAllAdmin(c appengine.Context) (admins []model.Admin, err error) {
	adminEntitys := make([]model.AdminEntity, 0)
	q := datastore.NewQuery("AdminEntity")
	keys, err := q.GetAll(c, &adminEntitys)
	if err != nil {
		return nil, err
	}
	admins = make([]model.Admin, len(keys))
	for i, key := range keys {
		admin := dsHelper.Entity2Model(key, &adminEntitys[i]).(model.Admin)
		admins[i] = admin
	}
	return admins, nil
}

func AdminCount(c appengine.Context) (count int, err error) {
	q := datastore.NewQuery("AdminEntity").KeysOnly()
	keys, err := q.GetAll(c, nil)
	if err != nil {
		return -1, err
	}
	return len(keys), nil
}
