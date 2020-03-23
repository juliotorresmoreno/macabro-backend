package db

import "github.com/go-xorm/xorm"

type Engine struct {
	permisionQueryRead  string
	permisionQueryWrite string
	user                string
	group               string
	*xorm.Engine
}

func (e *Engine) Get(bean interface{}) (bool, error) {
	return e.NewSession().Get(bean)
}

func (e *Engine) Update(bean interface{}, condiBean ...interface{}) (int64, error) {
	return e.NewSession().Update(bean, condiBean...)
}

func (e *Engine) Delete(bean interface{}) (int64, error) {
	return e.NewSession().Update(bean)
}

func (e *Engine) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error {
	return e.NewSession().Find(rowsSlicePtr, condiBean...)
}

func (e *Engine) Select(str string) *Session {
	return e.NewSession().Select(str)
}

func (e *Engine) Insert(bean ...interface{}) (int64, error) {
	return e.NewSession().Insert(bean...)
}

func (e *Engine) InsertOne(bean interface{}) (int64, error) {
	return e.NewSession().InsertOne(bean)
}

func (e *Engine) NewSession() *Session {
	return &Session{
		user:                e.user,
		group:               e.group,
		permisionQueryRead:  e.permisionQueryRead,
		permisionQueryWrite: e.permisionQueryWrite,
		Session:             e.Engine.NewSession(),
	}
}

func (e *Engine) Where(query interface{}, args ...interface{}) *Session {
	return e.NewSession().Where(query, args...)
}

func (e *Engine) Table(tableNameOrBean interface{}) *Session {
	return e.NewSession().Table(tableNameOrBean)
}
