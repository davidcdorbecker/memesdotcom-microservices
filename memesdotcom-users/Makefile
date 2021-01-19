mysql:
	docker run --name mysql -v mysqldata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=pass -d --rm -p 3306:3306 mysql
createdb:
	docker exec -i mysql mysql -uroot -ppass <<< "create database users_db;"
dropdb:
	docker exec -i mysql mysql -uroot -ppass <<< "drop database users_db;"
restore:
	cat ./db/migration/init_schema.up.sql | docker exec -i mysql /usr/bin/mysql -uroot -ppass users_db
backup:
	docker exec -i mysql /usr/bin/mysqldump -uroot -ppass users_db > backup.sql

.PHONY: mysql createdb dropdb restore backup