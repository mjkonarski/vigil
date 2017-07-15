#!bin/bash

./wait-for-it.sh -h $POSTGRES_HOST -p $POSTGRES_PORT
RESULT=$?
if [[ $RESULT -ne 0 ]]; then
    exit
fi

goose postgres "host=$POSTGRES_HOST user=$POSTGRES_USER dbname=$VIGIL_DB password=$POSTGRES_PASSWORD sslmode=disable" up
RESULT=$?
if [[ $RESULT -ne 0 ]]; then
    exit
fi

go-wrapper run
