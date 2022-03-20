package service

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/helloworld/v1"
	"{{cookiecutter.module_name}}/internal/biz"
)

// {{cookiecutter.service_name}}Service is a {{cookiecutter.repo_name}} service.
type {{cookiecutter.service_name}}Service struct {
	v1.Unimplemented{{cookiecutter.service_name}}Server

	uc *biz.{{cookiecutter.service_name}}Usecase
}

// New{{cookiecutter.service_name}}Service new a {{cookiecutter.repo_name}} service.
func New{{cookiecutter.service_name}}Service(uc *biz.{{cookiecutter.service_name}}Usecase) *{{cookiecutter.service_name}}Service {
	return &{{cookiecutter.service_name}}Service{uc: uc}
}

// SayHello implements helloworld.{{cookiecutter.service_name}}Server.
func (s *{{cookiecutter.service_name}}Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.Create{{cookiecutter.service_name}}(ctx, &biz.{{cookiecutter.service_name}}{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
