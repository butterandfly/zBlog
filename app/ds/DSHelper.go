package ds

import (
	M "app/model"
	"appengine"
	"appengine/datastore"
	"reflect"
	"strings"
)

var dsHelper = NewDSHelpper()

type DSHelpper struct {
	structMap map[string]reflect.Type
}

func NewDSHelpper() (dshelpper *DSHelpper) {
	// init helpper
	self := &DSHelpper{}
	self.structMap = make(map[string]reflect.Type)

	// init struct map
	articleType := reflect.TypeOf((*M.Article)(nil)).Elem()
	articleEntityType := reflect.TypeOf((*M.ArticleEntity)(nil)).Elem()
	adminType := reflect.TypeOf((*M.Admin)(nil)).Elem()
	adminEntityType := reflect.TypeOf((*M.AdminEntity)(nil)).Elem()
	blogCubeType := reflect.TypeOf((*M.BlogCube)(nil)).Elem()
	blogCubeEntityType := reflect.TypeOf((*M.BlogCubeEntity)(nil)).Elem()
	aFloatType := reflect.TypeOf((*M.AFloat)(nil)).Elem()
	aFloatEntityType := reflect.TypeOf((*M.AFloatEntity)(nil)).Elem()
	// base type
	var nullString string
	stringType := reflect.TypeOf(nullString)
	byteArrayType := reflect.TypeOf(([]byte)(nil))

	self.structMap["Article"] = articleType
	self.structMap["ArticleEntity"] = articleEntityType
	self.structMap["Admin"] = adminType
	self.structMap["AdminEntity"] = adminEntityType
	self.structMap["BlogCube"] = blogCubeType
	self.structMap["BlogCubeEntity"] = blogCubeEntityType
	self.structMap["AFloat"] = aFloatType
	self.structMap["AFloatEntity"] = aFloatEntityType
	// base type
	self.structMap["String"] = stringType
	self.structMap["ByteArray"] = byteArrayType

	return self
}

func (self *DSHelpper) Model2Entity(model interface{}) (key *datastore.Key, entity interface{}) {
	// model := &BlogCube{"zero", "1"}

	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()
	// get model name
	modelName := modelType.String()
	arr := strings.Split(modelName, ".")
	modelName = arr[len(arr)-1]

	// create entity
	entityName := modelName + "Entity"
	entityType := self.structMap[entityName]
	entityValue := reflect.New(entityType).Elem()

	// copy field
	count := modelValue.NumField()
	for i := 0; i < count; i++ {
		field := modelType.Field(i)
		fieldType := field.Type
		fieldValue := modelValue.Field(i)
		// ID situation
		if field.Name == "ID" {
			key, _ = datastore.DecodeKey(fieldValue.Interface().(string))
			continue
		}
		entityField := entityValue.FieldByName(field.Name)
		entityFieldType := entityField.Type()
		// string 2 []byte situation
		if entityFieldType == self.structMap["ByteArray"] && fieldType == self.structMap["String"] {
			str := fieldValue.Interface().(string)
			strByte := []byte(str)
			fieldValue = reflect.ValueOf(strByte)
		}
		entityField.Set(fieldValue)
	}
	entity = entityValue.Interface()

	return key, entity
}

func (self *DSHelpper) Entity2Model(key *datastore.Key, entity interface{}) (model interface{}) {
	// get entity type, name, value
	entityValue := reflect.ValueOf(entity).Elem()
	entityType := entityValue.Type()
	entityName := entityType.String()
	arr := strings.Split(entityName, ".")
	entityName = arr[len(arr)-1]

	// set model name, type, value
	modelName := entityName[0 : len(entityName)-6]
	// trimright has bug
	// modelName := strings.TrimRight(entityName, "Entity")
	modelType := self.structMap[modelName]
	modelValue := reflect.New(modelType).Elem()

	// copy field
	count := entityValue.NumField()
	for i := 0; i < count; i++ {
		field := entityType.Field(i)
		fieldType := field.Type
		fieldValue := entityValue.Field(i)
		modelField := modelValue.FieldByName(field.Name)
		modelFieldType := modelField.Type()

		// []byte 2 string situation
		if fieldType == self.structMap["ByteArray"] && modelFieldType == self.structMap["String"] {
			bytes := fieldValue.Interface().([]byte)
			str := string(bytes)
			fieldValue = reflect.ValueOf(str)
		}
		modelField.Set(fieldValue)
	}
	modelIDField := modelValue.FieldByName("ID")
	idValue := reflect.ValueOf(key.Encode())
	modelIDField.Set(idValue)

	model = modelValue.Interface()

	return model
}

func (self *DSHelpper) addEntity(c appengine.Context, entityName string, entityP interface{}) (key *datastore.Key, err error) {

	key, err = datastore.Put(c, datastore.NewIncompleteKey(c, entityName, nil), entityP)

	return key, err
}

/* can not work
func (self *DSHelpper) addEntityV2(c appengine.Context, entityName string, model interface{}) (key *datastore.Key, err error) {
	_, entity := self.Model2Entity(model)
	entityValue := reflect.ValueOf(entity)
	if !entityValue.CanAddr() {
		c.Warningf("can not be address...")
	}
	ptr := entityValue.Addr().Interface()
	key, err = datastore.Put(c, datastore.NewIncompleteKey(c, entityName, nil), ptr)
	return key, err
}
*/

func (self *DSHelpper) getAllSortedByCreateTime(c appengine.Context, entityName string, entitySliceP interface{}) (keys []*datastore.Key, err error) {
	q := datastore.NewQuery(entityName).Order("-CreateTime")
	keys, err = q.GetAll(c, entitySliceP)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (self *DSHelpper) Entitys2Models(c appengine.Context, keys []*datastore.Key, entitys interface{}, models interface{}) {
	entitysValue := reflect.ValueOf(entitys)

	modelSliceType := reflect.TypeOf(models).Elem()
	modelS := reflect.MakeSlice(modelSliceType, 0, 0)

	for i, key := range keys {
		entityValue := entitysValue.Index(i)
		if entityValue.CanAddr() {
			// c.Warningf("It can be addr")
		} else {
			// c.Warningf("It can !!not be addr")
		}
		aModel := dsHelper.Entity2Model(key, entityValue.Addr().Interface())
		aModelValue := reflect.ValueOf(aModel)
		// c.Warningf("aModel is: %v", aModel)
		modelS = reflect.Append(modelS, aModelValue)
	}

	modelsValue := reflect.ValueOf(models).Elem()
	modelsValue.Set(modelS)
}

func (self *DSHelpper) DeleteById(c appengine.Context, id string) (err error) {
	// decode string
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	// delete
	err = datastore.Delete(c, key)
	return err
}

func (self *DSHelpper) GetById(c appengine.Context, id string, entityPtr interface{}) (key *datastore.Key, err error) {
	// get key
	key, err = datastore.DecodeKey(id)
	if err != nil {
		return nil, err
	}

	// get entity
	err = datastore.Get(c, key, entityPtr)

	return key, err
}
