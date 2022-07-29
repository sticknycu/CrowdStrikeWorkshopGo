package domain

type MyStorage interface {
	GetMyContent(id string) (string, error)
	WriteMyContent(id string, content string) error
}
