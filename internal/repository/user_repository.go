package repository

import (
	"context"
	"database/sql"

	"github.com/VadimRight/User_Microserver/domain"
	"github.com/VadimRight/User_Microserver/domain/entity"
	"github.com/VadimRight/User_Microserver/internal/repository/repositories_query"
	"github.com/VadimRight/User_Microserver/pkg/datasource"
	"github.com/VadimRight/User_Microserver/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db   *sqlx.DB
	conn datasource.ConnTx
}

func NewRepository(c *sqlx.DB) domain.Repository {
	return &repository{conn: c, db: c}
}

// Atomic implements Repository Interface for transaction query
func (r *repository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx domain.Repository) error) error {
	txConn, err := r.db.BeginTxx(ctx, opt)
	if err != nil {
		return err
	}

	newRepository := &repository{conn: txConn, db: r.db}

	repo(newRepository)

	if err := new(datasource.DataSource).EndTx(txConn, err); err != nil {
		return err
	}

	return nil
}

func (repo *repository) InsertUser(ctx context.Context, user entity.User) (userID int64, err error) {
	args := utils.Array{
		user.Email,
		user.Username,
		user.Password,
		user.IsActive,
		user.IsVerified,
	}

	err = new(datasource.DataSource).ExecSQL(repo.conn.ExecContext(ctx, repositories_query.InsertUser, args...)).Scan(nil, &userID)
	if err != nil {
		return userID, err
	}

	return userID, nil
}

func (repo *repository) GetUserByID(ctx context.Context, userID int64, options ...entities.LockingOpt) (userData entity.User, err error) {
	args := utils.Array{
		userID,
	}

	row := func(idx int) utils.Array {
		return utils.Array{
			&userData.Id,
			&userData.Email,
			&userData.Username,
			&userData.IsActive,
			&userData.IsVerified,
		}
	}

	query := repositories_query.GetUserByID

	if len(options) >= 1 && options[0].PessimisticLocking {
		query += " FOR UPDATE"
	}

	if err = new(datasource.DataSource).QuerySQL(repo.conn.QueryxContext(ctx, query, args...)).Scan(row); err != nil {
		return userData, err
	}

	return userData, err
}

func (repo *repository) IsUserExist(ctx context.Context, email string) bool {
	args := utils.Array{email}

	var id int64
	row := func(idx int) utils.Array {
		return utils.Array{
			&id,
		}
	}

	err := new(datasource.DataSource).QuerySQL(repo.conn.QueryxContext(ctx, repositories_query.IsUserExist, args...)).Scan(row)
	if err != nil {
		return false
	}

	return id != 0
}
