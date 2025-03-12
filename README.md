# backend-boilerplate-template
This is our official nexlura backend boilerplate template

- run : go mod download
- run : go install github.com/air-verse/air@latest
- run > air
- run > docker compose up
- run this command before and db goose command > source .env

- ### make migrations
  > goose create add_some_column sql
- ### apply all migrations
  > make up
- ### apply all seed
  > make seed
- ### remove seed data
  > make weed
- ### apply all reset
  > make reset
- ### check db migration status
  > make db-status
>
> minio server start

