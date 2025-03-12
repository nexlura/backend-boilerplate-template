package infrastructure

import (
	"context"
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/utilities"
	"log"
	"net/http"
	"time"
)

func SaveUser(newUser models.Profile) (models.Profile, InfraError) {
	/* Create new user */
	newUserSQL, err := utilities.ReadQuery("./queries/users/new.sql")
	if err != nil {
		log.Printf("Error reading user SQL file: %v", err)
		return models.Profile{}, InfraError{err, "cannot find sql query file in directory", http.StatusInternalServerError}
	}

	// Generate the user id
	newUser.ID = utilities.GenerateUUID()

	// Execute the query
	err = DB.Conn.QueryRow(context.Background(), newUserSQL, newUser.ID, newUser.FirstName, newUser.LastName,
		newUser.Email, newUser.Password, newUser.Avatar, newUser.Phone).Scan(&newUser.ID)
	if err != nil {
		return models.Profile{}, InfraError{err, err.Error(), http.StatusInternalServerError}
	}

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	return newUser, InfraError{}
}

//func FetchUsers(page int) ([]models.ProfileDto, InfraError) {
//	limit := 10
//	offset := (page - 1) * limit
//
//	getAllSQL, err := utilities.ReadQuery("./queries/users/get_list.sql")
//	if err != nil {
//		return []models.ProfileDto{}, InfraError{err, "cannot find sql query file in directory", http.StatusInternalServerError}
//	}
//
//	rows, err := DB.Conn.Query(context.Background(), string(getAllSQL), limit, offset)
//	if err != nil {
//		log.Printf("Error retrieving accounts: %v", err)
//		return []models.ProfileDto{}, InfraError{err, err.Error(), http.StatusInternalServerError}
//	}
//	defer rows.Close()
//
//	var users []*models.Profile
//
//	for rows.Next() {
//		var u models.Profile
//		if err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Phone, &u.Address, &u.Country,
//			&u.AccountType, &u.Status, &u.Avatar, &u.Provider, &u.AuthToken, &u.CreatedAt, &u.UpdatedAt); err != nil {
//
//			log.Printf("Error scanning accounts: %v", err)
//			continue
//		}
//		users = append(users, &u)
//	}
//
//	if err = rows.Err(); err != nil {
//		log.Printf("Error finalizing row processing: %v", err)
//		return []models.ProfileDto{}, InfraError{err, err.Error(), http.StatusInternalServerError}
//	}
//
//	return models.ProfileFromDomainList(users), InfraError{}
//}
//
//func UpdateUser(user models.Profile, id string) (models.Profile, InfraError) {
//	// Get the sql query
//	updateSQL, err := utilities.ReadQuery("./queries/users/update.sql")
//	if err != nil {
//		log.Printf("Error reading update query: %v", err)
//		return models.Profile{}, InfraError{err, "cannot find sql query file in directory",
//			http.StatusInternalServerError}
//	}
//
//	// Execute the query
//	_, err = DB.Conn.Exec(context.Background(), string(updateSQL), user.FirstName, user.LastName, user.Email,
//		user.Password, user.Phone, user.Address, user.Country, user.AccountType, user.Status, user.Avatar,
//		user.AuthToken, id)
//	if err != nil {
//		log.Printf("Error updating user profile: %v", err)
//		return models.Profile{}, InfraError{err, err.Error(), http.StatusInternalServerError}
//	}
//
//	// Get the updated user
//	updatedUser, _ := FetchUserById(id)
//
//	return updatedUser, InfraError{}
//}
//
//func RemoveUser(id string) InfraError {
//	// Get the sql query
//	deleteSQL, err := utilities.ReadQuery("./queries/users/delete.sql")
//	if err != nil {
//		return InfraError{err, "cannot find sql query file in directory", http.StatusInternalServerError}
//	}
//
//	_, err = DB.Conn.Exec(context.Background(), string(deleteSQL), id)
//	if err != nil {
//		log.Printf("Error deleting account: %v", err)
//		return InfraError{err, err.Error(), http.StatusInternalServerError}
//	}
//
//	return InfraError{}
//}
//
//func FetchUserByEmail(email string) (models.Profile, InfraError) {
//	// Get the sql query
//	userSQL, err := utilities.ReadQuery("./queries/users/get_by_email.sql")
//	if err != nil {
//		return models.Profile{}, InfraError{err, "cannot find sql query file in directory", http.StatusInternalServerError}
//	}
//
//	var user models.Profile
//	err = DB.Conn.QueryRow(context.Background(), string(userSQL), email).Scan(
//		&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Avatar,
//		&user.Provider, &user.Phone, &user.AuthToken, &user.CreatedAt, &user.UpdatedAt)
//
//	if err != nil {
//		//log.Printf("Error finding user: %v", err)
//		return models.Profile{}, InfraError{
//			err,
//			"user does not exist",
//			http.StatusNotFound,
//		}
//	}
//
//	return user, InfraError{}
//}
//
//func FetchUserById(id string) (models.Profile, InfraError) {
//	// Read the SQL
//	getSQL, err := utilities.ReadQuery("./queries/users/get_by_id.sql")
//	if err != nil {
//		return models.Profile{}, InfraError{err, "cannot find sql query file in directory", http.StatusInternalServerError}
//	}
//
//	var user models.Profile
//	err = DB.Conn.QueryRow(context.Background(), string(getSQL), id).Scan(
//		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Address, &user.Country,
//		&user.AccountType, &user.Status, &user.Avatar, &user.Provider, &user.AuthToken, &user.CreatedAt, &user.UpdatedAt,
//	)
//
//	if err != nil {
//		log.Printf("Error retrieving account: %v", err)
//		return models.Profile{}, InfraError{err, "account not found", http.StatusNotFound}
//	}
//
//	return user, InfraError{}
//}

//pgx.CollectRows(rows, pgx.RowToStructByName[models.Properties])
