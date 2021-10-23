package domain

import "fmt"

type HelloWorldService interface {
	Gretting() string
	SaveGretting(gretting string) string
}

type basicHelloWorld struct {
	HelloWorldRepository
}

func NewBasicHelloWorld() HelloWorldService {
	return &basicHelloWorld{}
}

func (b *basicHelloWorld) Gretting() string {
	return "Hello World"
}

func (b *basicHelloWorld) SaveGretting(gretting string) string {
	if err := b.HelloWorldRepository.Save(gretting); err != nil {
		fmt.Printf("Error when trying to save gretting: %v", err.Error())
		return "Error"
	}
	return gretting
}
