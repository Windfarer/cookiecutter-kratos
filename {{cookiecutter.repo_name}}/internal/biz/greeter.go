package biz

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// {{cookiecutter.service_name}} is a {{cookiecutter.service_name}} model.
type {{cookiecutter.service_name}} struct {
	Hello string
}

// {{cookiecutter.service_name}}Repo is a Greater repo.
type {{cookiecutter.service_name}}Repo interface {
	Save(context.Context, *{{cookiecutter.service_name}}) (*{{cookiecutter.service_name}}, error)
	Update(context.Context, *{{cookiecutter.service_name}}) (*{{cookiecutter.service_name}}, error)
	FindByID(context.Context, int64) (*{{cookiecutter.service_name}}, error)
	ListByHello(context.Context, string) ([]*{{cookiecutter.service_name}}, error)
	ListAll(context.Context) ([]*{{cookiecutter.service_name}}, error)
}

// {{cookiecutter.service_name}}Usecase is a {{cookiecutter.service_name}} usecase.
type {{cookiecutter.service_name}}Usecase struct {
	repo {{cookiecutter.service_name}}Repo
	log  *log.Helper
}

// New{{cookiecutter.service_name}}Usecase new a {{cookiecutter.service_name}} usecase.
func New{{cookiecutter.service_name}}Usecase(repo {{cookiecutter.service_name}}Repo, logger log.Logger) *{{cookiecutter.service_name}}Usecase {
	return &{{cookiecutter.service_name}}Usecase{repo: repo, log: log.NewHelper(logger)}
}

// Create{{cookiecutter.service_name}} creates a {{cookiecutter.service_name}}, and returns the new {{cookiecutter.service_name}}.
func (uc *{{cookiecutter.service_name}}Usecase) Create{{cookiecutter.service_name}}(ctx context.Context, g *{{cookiecutter.service_name}}) (*{{cookiecutter.service_name}}, error) {
	uc.log.WithContext(ctx).Infof("Create{{cookiecutter.service_name}}: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
