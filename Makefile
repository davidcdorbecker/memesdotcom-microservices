up-project:
	docker-compose up --build
up-from-scratch:
	docker-compose up -d mysql
	sleep 5
	docker exec -i mysql mysql -uroot -ppass -e "create database users_db;"
	cat ./memesdotcom-users/db/migration/init_schema.up.sql | docker exec -i mysql /usr/bin/mysql -uroot -ppass users_db
	docker-compose up --no-recreate
down-project:
	docker-compose down
users:
	docker-compose up -d mysql
	docker-compose up -d users
auth:
	docker-compose up -d redis
	docker-compose up -d auth
createdb:
	docker exec -i mysql mysql -uroot -ppass -e "create database users_db;"
dropdb:
	docker exec -i mysql mysql -uroot -ppass -e "drop database users_db;"
restore:
	cat ./memesdotcom-users/db/migration/init_schema.up.sql | docker exec -i mysql /usr/bin/mysql -uroot -ppass users_db
backup:
	docker exec -i mysql /usr/bin/mysqldump -uroot -ppass users_db > backup.sql
