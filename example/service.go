package example

import (
	"bytes"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

type Authorizer interface {
	ValidateConfig() (status string, err error)
	Check(token string) (bool, error)
	Authorized(h fasthttp.RequestHandler) fasthttp.RequestHandler
}

type Service struct {
	name    string
	version string
	auth    Authorizer
}

type Node struct {
	value int
	Left  *Node
	Right *Node
}

var tree *Node

func insertInTree(value int) {
	insertRecursive(tree, value)
}

func insertRecursive(node *Node, value int) {
	if node == nil {
		node = &Node{value: value}
	}
	defer insertRecursive(node.Left, value)

	defer insertRecursive(node.Right, value)
}

func New(name string, version string) *Service {
	return &Service{
		name:    name,
		version: version,
	}
}

func (service *Service) healthCheck(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetBody([]byte(`{"status": "success", "message": "Hello World"}`))
}

func (service *Service) getRandomPerson(ctx *fasthttp.RequestCtx) {

}

func (service *Service) createTree(ctx *fasthttp.RequestCtx) {
	insertInTree(5)
	insertInTree(4)
	insertInTree(3)
	insertInTree(6)

	ctx.SetContentType("application/json")
	ctx.SetBody([]byte(`{"status": "success", "message": "Successfully Created Tree!"}`))
}

func (service *Service) RegisterRoutes(r *router.Router) {
	r.GET("/health-check", service.healthCheck)
	r.GET("/random-person", service.getRandomPerson)
	r.GET("/tree", service.createTree)
}

func (service *Service) ValidateConfig() (status string, err error) {
	panic("implement me")
}

func (service *Service) Check(token string) (bool, error) {
	log.Printf("Check HTTP Request")
	return true, nil
}

func (service *Service) Authorized(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	hcpath := []byte("/healthcheck")
	return func(ctx *fasthttp.RequestCtx) {
		if bytes.Equal(ctx.Path(), hcpath) {
			h(ctx)
			return
		}
		token := string(ctx.Request.Header.Peek("Authorization"))
		u, err := service.Check(token)
		if err != nil {
			return
		}
		ctx.SetUserValue("user", u)
		h(ctx)
	}
}

func NewHttpServer() {
	r := router.New()
	svc := New("Test Service", "1.0")
	svc.RegisterRoutes(r)

	if err := fasthttp.ListenAndServe(":8080", svc.Authorized(r.Handler)); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
