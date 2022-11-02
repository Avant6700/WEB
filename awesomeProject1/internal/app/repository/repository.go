package repository

import (
	"awesomeProject1/internal/app/dsn"
	"awesomeProject1/internal/app/model"
	"errors"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New() (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetComics() ([]model.Comics, error) {
	var comics []model.Comics
	err := r.db.Find(&comics).Error
	return comics, err
}

func (r *Repository) GetComicsPrice(uuid uuid.UUID, comics *model.Comics) (uint64, error) {
	err := r.db.First(&comics, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	return 0, nil
}

func (r *Repository) AddComics(comics model.Comics) error {
	err := r.db.Create(&comics).Error
	return err
}

func (r *Repository) ChangePrice(uuid uuid.UUID, price uint64) (int, error) {
	var comics model.Comics
	err := r.db.First(&comics, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Model(&comics).Update("Price", price).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) DeleteComics(uuid uuid.UUID) (int, error) {
	var comics model.Comics
	err := r.db.First(&comics, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Delete(&comics, uuid).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}
