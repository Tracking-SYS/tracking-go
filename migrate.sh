#!/bin/bash

if [[ -z "${MYSQL_ADDR}" ]]; then
  MYSQL_ADDR="root:123@tcp(localhost:3306)/tracking"
fi

if [ $1 = "up" ]; then
	echo "Upgrade version $2"
	./migrate -source file://./db/migrations/ -database "mysql://${MYSQL_ADDR}" up $2
elif [ $1 = "down" ]; then
	echo "Downgrade version $2"
	./migrate -source file://./db/migrations/ -database "mysql://${MYSQL_ADDR}" down $2
elif [ $1 = "reset" ]; then
	echo "Reset All"
	./migrate -source file://./db/migrations/ -database "mysql://${MYSQL_ADDR}" drop -f
elif [ $1 = "create" ]; then
	echo "Create table $2"
	./migrate create -ext sql -seq -dir ./db/migrations $2
fi