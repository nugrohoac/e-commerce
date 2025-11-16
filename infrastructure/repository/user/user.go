package user

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/nugrohoac/e-commerce/entity"
)

//go:generate mockery --name Repository --output ../../../mocks/infrastructure/repository/user
type Repository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByPhone(ctx context.Context, phone string) (*entity.User, error)
}

type repository struct {
	db *sql.DB
}

func (r repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query, args, err := sq.Select("id",
		"name",
		"email",
		"phone",
		"password_hash",
	).From("user").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	var user entity.User
	if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.PasswordHash); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r repository) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	query, args, err := sq.Select("id",
		"name",
		"email",
		"phone",
		"password_hash",
	).From("user").
		Where(sq.Eq{"phone": phone}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	var user entity.User
	if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.PasswordHash); err != nil {
		return nil, err
	}

	return &user, nil
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

type Example interface {
	Run(ctx context.Context) error
}
