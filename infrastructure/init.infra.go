package infrastructure

import (
	"github.com/backend-boilerplate-template/utilities/startup"
)

type InfraError struct {
	Error        error
	ErrorMessage string
	ErrorCode    int
}

var (
	DB = startup.InitializeDatabaseConnection()
)

//func(){
//	defer DB.Conn.Close()
//}()
