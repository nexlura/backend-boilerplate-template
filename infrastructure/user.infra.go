package infrastructure

import (
	"context"
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/utilities"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

func SaveUser(newUser models.Profile) (models.Profile, responses.ResponseError) {
	// Read the sql file
	newUserSQL, sqlErr := utilities.ReadQuery("./queries/users/new.sql")
	if sqlErr != nil {
		//log.Printf("Error reading user SQL file: %v", sqlErr)
		return models.Profile{}, responses.ResponseError{
			Error: sqlErr, ErrorMessage: "cannot find sql query file in directory",
			ErrorCode: http.StatusInternalServerError}
	}

	// Generate the user id
	newUser.ID = utilities.GenerateUUID()

	// Execute the insertion query
	queriedRow, queriedError := DB.Conn.Query(
		context.Background(), newUserSQL, newUser.ID, newUser.FirstName,
		newUser.LastName, newUser.Email, newUser.Password, newUser.Phone,
		newUser.RoleId, newUser.Status, newUser.Avatar)

	// Throw error if insertion failed
	if queriedError != nil {
		return models.Profile{}, responses.ResponseError{Error: queriedError, ErrorMessage: queriedError.Error(),
			ErrorCode: responses.StatusInternalServerError}
	}

	// Close the db connection after querying
	defer queriedRow.Close()

	// Get the inserted row
	collectedRow, collectedRowError := pgx.CollectOneRow(queriedRow, pgx.RowToStructByName[models.Profile])

	// Throw error if getting the inserted row failed
	if collectedRowError != nil {
		return models.Profile{}, responses.ResponseError{Error: collectedRowError,
			ErrorMessage: collectedRowError.Error(),
			ErrorCode:    responses.StatusInternalServerError}
	}

	// Return the newly inserted item
	return collectedRow, responses.ResponseError{}
}

func FetchUsers(page int) ([]models.Profile, responses.ResponseError) {
	// Set the pagination limit and offset
	limit := 10
	offset := (page - 1) * limit

	// Read the sql file
	getAllSQL, sqlErr := utilities.ReadQuery("./queries/users/get_list.sql")

	// Throw error if any occurs when reading the sql file
	if sqlErr != nil {
		return []models.Profile{}, responses.ResponseError{Error: sqlErr,
			ErrorMessage: "cannot find sql query file in directory",
			ErrorCode:    responses.StatusInternalServerError}
	}

	// Execute the select all query
	queriedRows, queriedError := DB.Conn.Query(context.Background(), string(getAllSQL), limit, offset)

	// Throw error if select all query failed
	if queriedError != nil {
		log.Printf("Error retrieving accounts: %v", queriedError)
		return []models.Profile{}, responses.ResponseError{
			Error: queriedError, ErrorMessage: queriedError.Error(),
			ErrorCode: responses.StatusInternalServerError,
		}
	}

	// Close the db connection after querying
	defer queriedRows.Close()

	// Get the inserted row
	collectedItem, collectedItemErr := pgx.CollectRows(queriedRows, pgx.RowToStructByName[models.Profile])

	// Throw error if getting the selected row failed
	if collectedItemErr != nil {
		return []models.Profile{}, responses.ResponseError{Error: collectedItemErr,
			ErrorMessage: collectedItemErr.Error(),
			ErrorCode:    responses.StatusInternalServerError}
	}

	// Return the collected item
	return collectedItem, responses.ResponseError{}
}

func FetchUserByParam(param string) (models.Profile, responses.ResponseError) {
	// Read the SQL
	getSQL, sqlErr := utilities.ReadQuery("./queries/users/get_by_args.sql")
	if sqlErr != nil {
		return models.Profile{}, responses.ResponseError{Error: sqlErr,
			ErrorMessage: "cannot find sql query file in directory",
			ErrorCode:    responses.StatusInternalServerError}
	}

	//var user models.Profile
	queriedRow, queriedError := DB.Conn.Query(context.Background(), string(getSQL), param)

	// Throw error if select all query failed
	if queriedError != nil {
		log.Printf("%v", queriedError.Error())
		return models.Profile{}, responses.ResponseError{
			Error: queriedError, ErrorMessage: queriedError.Error(),
			ErrorCode: responses.StatusInternalServerError,
		}
	}

	// Close the db connection after querying
	defer queriedRow.Close()

	// Get the row
	collectedRow, collectedRowError := pgx.CollectOneRow(queriedRow, pgx.RowToStructByName[models.Profile])

	// Throw error if getting the row failed
	if collectedRowError != nil {
		return models.Profile{}, responses.ResponseError{
			Error:        collectedRowError,
			ErrorMessage: collectedRowError.Error(),
			ErrorCode:    responses.StatusNotFound,
		}
	}

	// Return the item
	return collectedRow, responses.ResponseError{}
}

func AlterUser(payload models.ProfileFrom) (models.Profile, responses.ResponseError) {
	// Read the sql query
	updateSQL, sqlErr := utilities.ReadQuery("./queries/users/update.sql")
	if sqlErr != nil {
		log.Printf("Error reading update query: %v", sqlErr)
		return models.Profile{}, responses.ResponseError{Error: sqlErr,
			ErrorMessage: "cannot find sql query file in directory",
			ErrorCode:    http.StatusInternalServerError}
	}

	// Execute the insertion query
	_, queriedError := DB.Conn.Query(
		context.Background(), string(updateSQL), payload.ID, payload.FirstName, payload.LastName, payload.Email,
		payload.Password, payload.Phone, payload.RoleId, payload.Status, payload.Avatar)

	// Throw error if updating failed
	if queriedError != nil {
		return models.Profile{}, responses.ResponseError{Error: queriedError, ErrorMessage: queriedError.Error(),
			ErrorCode: responses.StatusInternalServerError}
	}

	// Get the updated row
	collectedRow, _ := FetchUserByParam(payload.ID)

	// Return the newly inserted item
	return collectedRow, responses.ResponseError{}
}

func RemoveUser(id string) responses.ResponseError {
	// Read the sql query file
	deleteSQL, err := utilities.ReadQuery("./queries/users/delete.sql")
	if err != nil {
		return responses.ResponseError{Error: err,
			ErrorMessage: "cannot find sql query file in directory",
			ErrorCode:    responses.StatusInternalServerError,
		}
	}

	_, err = DB.Conn.Exec(context.Background(), string(deleteSQL), id)
	if err != nil {
		log.Printf("Error removing account: %v", err)
		return responses.ResponseError{
			Error:        err,
			ErrorMessage: err.Error(),
			ErrorCode:    responses.StatusInternalServerError,
		}
	}

	return responses.ResponseError{}
}
