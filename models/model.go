package models

// Model s
type Model interface {
	TableName() string
	Check() error
}
