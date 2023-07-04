package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type UserRepo interface {
	GetUserById(int) (*model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
}

type userRepoImpl struct {
	db *sql.DB
}

func (usrRepo *userRepoImpl) InsertUser(usr *model.UserModel) error {
	qry := "INSERT INTO user_credential(id,user_name, password, is_active) VALUES($1, $2, $3, $4)"

	_, err := usrRepo.db.Exec(qry, usr.Id, usr.Name, usr.Password, usr.Active)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.InsertUser() : %w", err)
	}
	return nil
}

func (usrRepo *userRepoImpl) GetUserByName(name string) (*model.UserModel, error) {
	qry := "SELECT id, user_name, is_active FROM user_credential WHERE user_name = $1"

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.Name, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetuserByName() : %w", err)
	}
	return usr, nil
}

func (usrRepo *userRepoImpl) GetUserById(id int) (*model.UserModel, error) {
	qry := "SELECT id, user_name, password, is_active FROM user_credential WHERE id = $1"

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, id).Scan(&usr.Id, &usr.Name, &usr.Password, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.getuserById() : %w", err)
	}
	return usr, nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
