package domain

type HelloWorldRepository interface {
	Save(string) error
}
