package infrastructure

import (
	"github.com/backend-boilerplate-template/utilities/startup"
)

var (
	DB = startup.InitializeDatabaseConnection()
)

//func(){
//	defer DB.Conn.Close()
//}()
