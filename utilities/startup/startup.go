package startup

import (
	"context"
	"embed"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Database struct {
	Conn *pgxpool.Pool // *pgx.Conn
}

type EnvType struct {
	dbURL                string
	mixPanelProjectToken string
}

func LoadEnv() EnvType {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//use ../.env because main.go inside /cmd
	err = godotenv.Load(filepath.Join(pwd, "/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dbURL string
	if os.Getenv("ENVIRONMENT") == "LOCAL" {
		dbURL = os.Getenv("LOCAL_DATABASE_URL")
	} else if os.Getenv("ENVIRONMENT") == "PROD" {
		dbURL = os.Getenv("PROD_DATABASE_URL")
	} else {
		dbURL = os.Getenv("DOCKER_DATABASE_URL")
	}

	var mixPanelProjectToken string
	if os.Getenv("MIXPANEL_PROJECT_TOKEN") != "" {
		mixPanelProjectToken = os.Getenv("MIXPANEL_PROJECT_TOKEN")
	}

	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	log.Println("loaded env")

	return EnvType{
		dbURL, mixPanelProjectToken,
	}
}

var (
	//db              *sql.DB
	embedMigrations embed.FS
)

func InitializeDatabaseConnection() Database {
	env := LoadEnv()

	//var pool *pgxpool.Pool

	conn, err := pgxpool.New(context.Background(), env.dbURL) //pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	//goose.SetBaseFS(embedMigrations)
	//
	//if dialectErr := goose.SetDialect("postgres"); err != nil {
	//	panic(dialectErr)
	//}
	//
	//log.Println("running migrations")
	//db := stdlib.OpenDBFromPool(pool)
	//if migrationErr := goose.Up(db, "database/migrations"); err != nil {
	//	panic(migrationErr)
	//}

	log.Println(env.dbURL)

	log.Printf("Successfully connected to database")
	return Database{Conn: conn}
}
