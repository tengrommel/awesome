package services

import (
	"awesome/web/bookstore_users_api/domain/users"
	"awesome/web/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
	//return nil, &errors.RestErr{
	//	Status: http.StatusInternalServerError
	//}
}
