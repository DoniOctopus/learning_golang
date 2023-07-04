package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type UserRepo interface {
	// GetUserById(int) (*model.UserModel, error)
	// GetUserByName(string) (*model.UserModel, error)
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

// func (svcRepo *serviceRepoImpl) GetServiceByName(name string) (*model.ServiceModel, error) {
// 	qry := "SELECT id, name, uom, price FROM ms_service WHERE name = $1"

// 	svc := &model.ServiceModel{}
// 	err := svcRepo.db.QueryRow(qry, name).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceByName() : %w", err)
// 	}
// 	return svc, nil
// }

// func (svcRepo *serviceRepoImpl) GetServiceById(id int) (*model.ServiceModel, error) {
// 	qry := "SELECT id, name, uom, price FROM ms_service WHERE id = $1"

// 	svc := &model.ServiceModel{}
// 	err := svcRepo.db.QueryRow(qry, id).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, fmt.Errorf("error on serviceRepoImpl.getServiceById() : %w", err)
// 	}
// 	return svc, nil
// }

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
