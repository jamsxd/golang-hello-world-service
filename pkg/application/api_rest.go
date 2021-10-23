package application

import (
	"fmt"
	"net/http"

	"github.com/jamsxd/golang-hello-world-service/pkg/domain"
)

type HelloWorldHandler interface {
	GetGretting(w http.ResponseWriter, r *http.Request)
}

type basicHelloWorldHandler struct {
	domain.HelloWorldService
}

func NewBasicHelloWorldHandler(helloWorld domain.HelloWorldService) HelloWorldHandler {
	return &basicHelloWorldHandler{
		HelloWorldService: helloWorld,
	}
}

func (b *basicHelloWorldHandler) GetGretting(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, b.HelloWorldService.Gretting())
}
