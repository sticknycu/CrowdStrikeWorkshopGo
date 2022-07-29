package domain

//go:generate mockgen -source=storage.go -package=mocks -destination=mocks/storage.go
//go:generate mockgen -source=my_storage.go -package=mocks -destination=mocks/my_storage.go

type Storage interface {
	GetContent(id string) (string, error)
	WriteContent(id string, content string) error
}
