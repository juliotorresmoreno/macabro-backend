package db

import (
	"fmt"
	"reflect"

	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/macabro/models"
)

type Session struct {
	permisionQueryRead  string
	permisionQueryWrite string
	user                string
	group               string
	*xorm.Session
}

func (e *Session) getPermisionQueryRead() string {
	return fmt.Sprintf(e.permisionQueryRead, e.user, e.group)
}

func (e *Session) getPermisionQueryWrite() string {
	return fmt.Sprintf(e.permisionQueryWrite, e.user, e.group)
}

func (e *Session) Get(bean interface{}) (bool, error) {
	if e.permisionQueryRead != "" {
		return e.Session.Where(e.getPermisionQueryRead()).Get(bean)
	}
	return e.Session.Get(bean)
}

func (e *Session) Update(bean interface{}, condiBean ...interface{}) (int64, error) {
	if e.permisionQueryWrite != "" {
		return e.Session.Where(e.getPermisionQueryWrite()).Update(bean, condiBean...)
	}
	return e.Session.Update(bean, condiBean...)
}

var aclType = reflect.ValueOf(models.ACL{}).Type()

func (e *Session) Insert(bean ...interface{}) (int64, error) {
	for _, b := range bean {
		field := reflect.ValueOf(b)
		if field.Kind() != reflect.Ptr {
			field = reflect.ValueOf(&b)
		}
		field = field.Elem().FieldByName("ACL")
		if field.CanSet() && field.Type() == aclType {
			acl := models.NewACL(e.user)
			field.Set(reflect.ValueOf(acl))
		}
	}

	return e.Session.Insert(bean...)
}

func (e *Session) InsertOne(bean interface{}) (int64, error) {
	field := reflect.ValueOf(bean)
	if field.Kind() != reflect.Ptr {
		field = reflect.ValueOf(&bean)
	}
	field = field.Elem().FieldByName("ACL")
	if field.CanSet() && field.Type() == aclType && field.IsZero() {
		acl := models.NewACL(e.user)
		field.Set(reflect.ValueOf(acl))
	}
	return e.Session.InsertOne(bean)
}

func (e *Session) Delete(bean interface{}) (int64, error) {
	if e.permisionQueryWrite != "" {
		return e.Session.Where(e.getPermisionQueryWrite()).Delete(bean)
	}
	return e.Session.Update(bean)
}

func (e *Session) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error {
	if e.permisionQueryRead != "" {
		return e.Session.Where(e.getPermisionQueryRead()).Find(rowsSlicePtr, condiBean...)
	}
	return e.Session.Find(rowsSlicePtr, condiBean...)
}

func (e Session) Select(str string) *Session {
	e.Session = e.Session.Select(str)
	return &e
}

func (e Session) Where(query interface{}, args ...interface{}) *Session {
	e.Session = e.Session.Where(query, args...)
	return &e
}

func (e Session) Or(query interface{}, args ...interface{}) *Session {
	e.Session = e.Session.Or(query, args...)
	return &e
}

func (e Session) And(query interface{}, args ...interface{}) *Session {
	e.Session = e.Session.And(query, args...)
	return &e
}
