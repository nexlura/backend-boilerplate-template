seed:
	goose -dir database/seeds -no-versioning up

weed:
	goose -dir database/seeds -no-versioning down

up:
	goose -dir database/migrations up

reset:
	goose -dir database/migrations reset

db-status:
	goose status
