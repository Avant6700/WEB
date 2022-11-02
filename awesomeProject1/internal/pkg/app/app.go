package app

import (
	"awesomeProject1/internal/app/repository"
	"context"
)

type Application struct {
	ctx  context.Context
	repo *repository.Repository
}

func New(ctx context.Context) (*Application, error) {
	app := &Application{
		ctx: ctx,
	}
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}
	app.repo = repo
	return app, nil
}

func (a *Application) Run(ctx context.Context) error {
	//ComicssCfg := config.FromContext(ctx).
	a.StartServer()
	return nil
}
